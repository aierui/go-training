package snowflake

import (
	"sync"
	"time"
)

const (
	workerIdBits     int64 = 4
	dataCenterIdBits int64 = 4
	sequenceBits     int64 = 10

	maxWorkerId     int64 = -1 ^ (-1 << uint64(workerIdBits))
	maxDataCenterId int64 = -1 ^ (-1 << uint64(dataCenterIdBits))
	maxSequence     int64 = -1 ^ (-1 << uint64(sequenceBits))

	timeLeft uint8 = 18
	dataLeft uint8 = 14
	workLeft uint8 = 10

	epoch int64 = 1577948645000 // 起始时间 2020-01-02 15:04:05
)

// 1位，不用。二进制中最高位为 1 的都是负数，但是我们生成的 id 一般都使用整数，所以这个最高位固定是 0
// 41位，用来记录时间戳（毫秒）。
// 41 位可以表示 2^41−1 个数字，
//
//如果只用来表示正整数（计算机中正数包含 0），可以表示的数值范围是：0 至 2^41−1，减 1 是因为可表示的数值范围是从 0 开始算的，而不是 1。\
//也就是说 41 位可以表示 2^41−1 个毫秒的值，转化成单位年则是 (2^41−1)/(1000∗60∗60∗24∗365)=69 年
//

// 10位，用来记录工作机器 id。

// 可以部署在 2^10=1024 个节点，包括 5 位 datacenterId 和 5 位 workerId\
// 5 位（bit）可以表示的最大正整数是 2^5−1=31，即可以用 0、1、2、3、....31 这 32 个数字，来表示不同的 datecenterId 或 workerId

// 12位，序列号，用来记录同毫秒内产生的不同 id。
// 12 位（bit）可以表示的最大正整数是 2^12−1=4095，
//即可以用 0、1、2、3、....4094 这 4095 个数字，来表示同一机器同一时间截（毫秒) 内产生的 4095 个 ID 序号

type worker struct {
	mu           sync.Mutex
	lastStamp    int64
	workerId     int64
	dataCenterId int64
	sequence     int64
}

func NewNode(workerId int64, dataCenterId int64) *worker {
	if workerId > maxWorkerId {
		workerId = 1
	}

	if dataCenterId > maxDataCenterId {
		dataCenterId = 1
	}

	return &worker{
		lastStamp:    -1,
		workerId:     workerId,
		dataCenterId: dataCenterId,
		sequence:     1,
	}
}

func (w *worker) getCurrentTime() int64 {
	return time.Now().UnixNano() / 1e6
}

//var i int = 1
func (w *worker) NextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	timestamp := w.getCurrentTime()
	if w.lastStamp == timestamp {
		// 这其实和 <==>
		// w.sequence++
		// if w.sequence++ > maxSequence  等价
		w.sequence = (w.sequence + 1) & maxSequence
		if w.sequence == 0 {
			// 之前使用 if, 只是没想到 GO 可以在一毫秒以内能生成到最大的 Sequence, 那样就会导致很多重复的
			// 这个地方使用 for 来等待下一毫秒
			for timestamp <= w.lastStamp {
				timestamp = w.getCurrentTime()
			}
		}
	} else {
		w.sequence = 0
	}
	w.lastStamp = timestamp

	return ((timestamp - epoch) << timeLeft) | (w.dataCenterId << dataLeft) | (w.workerId << workLeft) | w.sequence
}

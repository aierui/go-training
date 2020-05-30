package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	val string
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int)  {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i,j
}

func (pq *PriorityQueue) Push(x interface{})  {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop()  interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}


// 更新函数会修改队列中指定元素的优先级以及值。
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.val = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// 这个示例首先会创建一个优先队列，并在队列中包含一些元素
// 接着将一个新元素添加到队列里面，并对其进行操作
// 最后按优先级有序地移除队列中的各个元素。
func main() {
	// 一些元素以及它们的优先级。
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// 创建一个优先队列，并将上述元素放入到队列里面，
	// 然后对队列进行初始化以满足优先队列（堆）的不变性。
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			val:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// 插入一个新元素，然后修改它的优先级。
	item := &Item{
		val:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.val, 5)

	// 以降序形式取出并打印队列中的所有元素。
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s  index=%v\n", item.priority, item.val, item.index)
	}
}

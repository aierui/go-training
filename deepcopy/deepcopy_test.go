package deepcopy

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"

	"github.com/mohae/deepcopy"
)

type Basics struct {
	String      string
	Strings     []string
	StringArr   [4]string
	Bool        bool
	Bools       []bool
	Byte        byte
	Bytes       []byte
	Int         int
	Ints        []int
	Int8        int8
	Int8s       []int8
	Int16       int16
	Int16s      []int16
	Int32       int32
	Int32s      []int32
	Int64       int64
	Int64s      []int64
	Uint        uint
	Uints       []uint
	Uint8       uint8
	Uint8s      []uint8
	Uint16      uint16
	Uint16s     []uint16
	Uint32      uint32
	Uint32s     []uint32
	Uint64      uint64
	Uint64s     []uint64
	Float32     float32
	Float32s    []float32
	Float64     float64
	Float64s    []float64
	Complex64   complex64
	Complex64s  []complex64
	Complex128  complex128
	Complex128s []complex128
	Interface   interface{}
	Interfaces  []interface{}
	Params      map[string]interface{}
}

var src = Basics{
	String:      "kimchi",
	Strings:     []string{"uni", "ika"},
	StringArr:   [4]string{"malort", "barenjager", "fernet", "salmiakki"},
	Bool:        true,
	Bools:       []bool{true, false, true},
	Byte:        'z',
	Bytes:       []byte("abc"),
	Int:         42,
	Ints:        []int{0, 1, 3, 4},
	Int8:        8,
	Int8s:       []int8{8, 9, 10},
	Int16:       16,
	Int16s:      []int16{16, 17, 18, 19},
	Int32:       32,
	Int32s:      []int32{32, 33},
	Int64:       64,
	Int64s:      []int64{64},
	Uint:        420,
	Uints:       []uint{11, 12, 13},
	Uint8:       81,
	Uint8s:      []uint8{81, 82},
	Uint16:      160,
	Uint16s:     []uint16{160, 161, 162, 163, 164},
	Uint32:      320,
	Uint32s:     []uint32{320, 321},
	Uint64:      640,
	Uint64s:     []uint64{6400, 6401, 6402, 6403},
	Float32:     32.32,
	Float32s:    []float32{32.32, 33},
	Float64:     64.1,
	Float64s:    []float64{64, 65, 66},
	Complex64:   complex64(-64 + 12i),
	Complex64s:  []complex64{complex64(-65 + 11i), complex64(66 + 10i)},
	Complex128:  complex128(-128 + 12i),
	Complex128s: []complex128{complex128(-128 + 11i), complex128(129 + 10i)},
	Interfaces:  []interface{}{42, true, "pan-galactic"},
	Params: map[string]interface{}{
		"test":   1,
		"oid1":   "12345643das",
		"oid2":   2134567623213246,
		"data":   64.1,
		"slice1": "[\"test\", \"m1\", \"n1\"]",
	},
}

// go test -v -run=^$ -bench=. -benchtime=10s -cpuprofile=prof.cpu -memprofile=prof.mem -memprofilerate=2
// go tool pprof -http=:8081 prof.cpu
// go tool pprof -http=:8082 prof.mem

func TestGOB(t *testing.T) {
	var dst Basics
	err := GOBDeepCopy(&dst, &src)
	if err != nil {
		t.Error(err)
	}

	rtn := reflect.DeepEqual(dst, src)
	fmt.Println(rtn)
}

func TestDeepCopy(t *testing.T) {
	dst := deepcopy.Copy(src).(Basics)
	if !dst.Bool {
		t.Error("reflect deep copy failed")
	}
	rtn := reflect.DeepEqual(dst, src)
	fmt.Println(rtn)
}

func TestDeepCopyByPointer(t *testing.T) {
	dst := DeepCopyByPointer(&src)
	rtn := reflect.DeepEqual(&src, dst)
	if !rtn {
		t.Error("deep copy failed")
	}
}

func Benchmark_DeepCopyByPointer(b *testing.B) {
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		//var dst Basics
		dst := DeepCopyByPointer(&src)
		rtn := reflect.DeepEqual(&src, dst)
		if !rtn {
			b.Error("deep copy failed")
		}
	}
}

func Benchmark_GOBDeepCopy(b *testing.B) {
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		var dst Basics
		err := GOBDeepCopy(&dst, &src)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_ReflectDeepCopy(b *testing.B) {
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		dst := deepcopy.Copy(src).(Basics)
		if !dst.Bool {
			b.Error("reflect deep copy failed")
		}
	}
}

// GOBDeepCopy provides the method to creates a deep copy of whatever is passed to
// it and returns the copy in an interface. The returned value will need to be
// asserted to the correct type.
func GOBDeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func DeepCopyByPointer(src *Basics) *Basics {
	tmp := &Basics{}
	*tmp = *src
	return tmp
}

# Usage

- go get -u honnef.co/go/tools
- go install honnef.co/go/tools/cmd/structlayout
- go install honnef.co/go/tools/cmd/structlayout-pretty
- go install honnef.co/go/tools/cmd/structlayout-optimize
- go get github.com/ajstarks/svgo/structlayout-svg

```go
structlayout -json github.com/aierui/go-training/memory/align Ag  | structlayout-svg -t 'align-guarantee' > ./memory/align/ag.svg
```

```go
// 内存对其检测

golangci-lint run --disable-all --enable maligned align.go

```
 

```go
//go:nosplit (不支持栈增长)
func newproc(siz int32, fn *funcval) {
	argp := add(unsafe.Pointer(&fn), sys.PtrSize)
	gp := getg()
	pc := getcallerpc()
	systemstack(func() {
		newg := newproc1(fn, argp, siz, gp, pc)

		_p_ := getg().m.p.ptr()
		runqput(_p_, newg, true)

		if mainStarted {
			wakep()
		}
	})
}
```
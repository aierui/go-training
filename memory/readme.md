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
 

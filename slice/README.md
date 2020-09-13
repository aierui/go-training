# 问题



s := make([]v, hint)

[makeslice](https://github.com/golang/go/blob/master/src/runtime/slice.go#L83) implements Go map creation for make([]v, hint).


## 怎样正确估算切片的长度和容量？

- [growslice](https://github.com/golang/go/blob/master/src/runtime/slice.go#L125) handles slice growth during append.

## 切片的底层数组什么时候会被替换？

## 如果有多个切片指向了同一个底层数组，那么你认为应该注意些什么？



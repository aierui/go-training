# 问题

m := make(map[k]v, hint)

[makemap](https://github.com/golang/go/blob/master/src/runtime/map.go#L303) implements Go map creation for make(map[k]v, hint).


## 为什么字典的键类型会受到约束？

- hash table
- hash function
- bucket


## 字典的键类型不能是哪些类型？


## 在值为nil的字典上执行读操作会成功吗，那写操作呢？


## 字典类型的值是并发安全的吗？
# UUID

- [golang 生成良好的唯一ID/uuid库比较](https://golangnote.com/topic/177.html)

uuid v1-v5 都是有 rfc 文档的

另一方面来讲，对这类需求需要更多是发号器。

常见的方案有：
- 数据自增
- redis 计数器
- 雪花算法

开源方案有：
- uuid
- [TinyId](https://github.com/didi/tinyid)
- [Leaf](https://github.com/Meituan-Dianping/Leaf)
- [UidGenerator](https://github.com/baidu/uid-generator)
- ……



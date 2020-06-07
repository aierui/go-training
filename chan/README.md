## Channel


//关闭后的通道有以下特点：

/*

1.对一个关闭的通道再发送值就会导致panic。
2.对一个关闭的通道进行接收会一直获取值直到通道为空。
3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4.关闭一个已经关闭的通道会导致panic。


## 无缓冲通道

## 有缓冲的通道

## 单向通道


无缓冲的通道又称为阻塞的通道。我们来看一下下面的代码：



func main() {
    ch := make(chan int)
    ch <- 10
    fmt.Println("发送成功")
}

上面这段代码能够通过编译，但是执行的时候会出现以下错误：


 fatal error: all goroutines are asleep - deadlock!

    goroutine 1 [chan send]:
    main.main()
            .../

因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。

上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？


一种方法是启用一个goroutine去接收值，例如：


func recv(c chan int) {
    ret := <-c
    fmt.Println("接收成功", ret)
}
func main() {
    ch := make(chan int)
    go recv(ch) // 启用goroutine从通道接收值
    ch <- 10 // 发送
    fmt.Println("发送成功")
}

无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。


###############

先发送后接口 就会因为阻塞而产生 deadlock

func main() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
	go recv(ch) // 启用goroutine从通道接收值
}

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/Users/jinrong6/develop/sinawww/go-training/Golang_Puzzlers/src/puzzlers/channel/range-slice-example.go:13 +0x59

##############


*/

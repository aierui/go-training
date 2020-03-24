# Go In Action

几个初步的 Go 语言特点

- 语言级并发
- 强大的标准库
- 垃圾回收
- 类型安全
- 简洁的面向对象
- 反射
- 函数多值返回

一共 25 个关键字：

```
package         包声明
import          包引入

const           常量

var             变量

type            类型

goto            流程控制
if              流程控制
else            流程控制
break           流程控制
continue        流程控制
default         流程控制
switch          流程控制
case            流程控制
fallthrough     流程控制
for    	        流程控制

func            函数相关
return          函数相关
defer           函数相关

struct          面向对象
interface       面向对象

map             数据结构
range           数据结构

go              并发相关
chan            并发相关
select          并发相关
```


## 包声明

```
package main

``` 

默认地将源码目录作为包的 package name。

之外：某个名称（包内方法或变量或常量或结构体等）在包外是否可见，就取决于其首个字符是否为大写字母

**小写**字母开头的程序实体相当于 class 中的带 private 关键词的私有函数

**大写**字母开头的程序实体相当于 class 中的带 public 关键词的公有函数

Go 语言中的程序实体包括变量、常量、函数、结构体和接口。

## 包引用

```
// 方式一
import "fmt"

// 方式二
import (
    "fmt"
    "log"
)
```


- 支持相对路径和绝对路径
- 绝对路径：import "lib/math" math.Sin
- 别名导入：import m "lib/math" m.Sin
- 点导入：import . "lib/math" Sin
- 下划线导入：import _ "lib/math" 只调用包中 init 函数


## 常量

- 字面常量：字面常量（literal）是指程序中硬编码的常量。
- 常量定义：通过 const 关键字给字面常量指定一个友好的名字。
- 预定义常量：Go 语言预定义了：true、false、iota
- 枚举

```go
// 常量可定义为数值、布尔值或字符串等类型
const API_CODE_SUCCESS = 'success'

// 也可以指定具体的数据类型
const Pi float64 = 3.1415926

// 多个常量
const (
    size int64 = 1024
    eof = -1 // 无类型整形常量
)

// 常量右值可以是一个在**编译期运算**的常量表达式： 
const make = 1 << 3 

// 但是不可以是运行期才能够得出结果的表达式，
const params = os.GetEnv("para") // 这样会导致编译错误

// iota 比较特殊，每一个 const 关键字出现时被重置为0，然后在下一个const出现之前，每出现一次 iota，其所代表的数字会自动增1
const ( // iota 被重设为 0 
    c0 = iota // c0 = 0
    c1        // c1 = 1
    c2        // c2 = 2
)

const (
    a = 1 << iota // a = 1
    b             // b = 2
    c             // c = 4
)
```

## 变量

变量是几乎所有的编程语言中最基本的组成元素。

```go
// 用法一：
var a, b int = 2, 1024

// 用法二：只能用在函数内部,在函数外部使用则会无法编译通过，一般用var方式来定义全局变量
shortVar1, shortVar2 := 2, 1024

// 用法三, _ 的赋值会被丢弃
_, under2 := 34, 35

// 错误: 已声明但未使用的变量会在编译阶段报错
// var i int

// 多变量
var (
    v1 int
    v2 string
)
```


## 类型

Go 语言内置基础类型

- 布尔类型：bool
- 整形：int8、byte、int16、int、uint、uintptr等
- 浮点类型：float32、float64
- 复数类型：complex64、complex128
- 字符串：string
- 字符类型：rune
- 错误类型：error

Go 语言支持的复合类型

- 指针（pointer）
- 数组（array）
- 切片（slice）
- 字典（map）
- 通道（chan）
- 结构体（struct）
- 接口（interface）


## 类型转换

类型选择 通过圆括号中的关键字 type 使用类型断言语法。类型选择是类型转换的一种形式。

```go
value.(type)
```

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
	fmt.Printf("unexpected type %T", t)       // %T 输出 t 是什么类型
case bool:
	fmt.Printf("boolean %t\n", t)             // t 是 bool 类型
case int:
	fmt.Printf("integer %d\n", t)             // t 是 int 类型
case *bool:
	fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
case *int:
	fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
}
```


## 类型断言

若我们只关心一种类型呢？若我们知道该值拥有一个 string 而想要提取它呢？ 只需一种情况的类型选择就行，但它需要类型断言。类型断言接受一个接口值， 并从中提取指定的明确类型的值。其语法借鉴自类型选择开头的子句，但它需要一个明确的类型， 而非 type 关键字：


```go
value.(typeName)
```

```go
if str, ok := value.(string); ok {
	return str
} else if str, ok := value.(Stringer); ok {
	return str.String()
}
```

## array

数组是非常有用的，有时还能避免过多的内存分配， 但它们主要用作切片的构件。

```go
var arr [n]type // 在[n]type中，n表示数组的长度，type表示存储元素的类型, 即 C 中的静态数组。数组的长度是其类型的一部分，因此数组不能改变大小。
```

数组声明的几种方式：

```go
// 方式一：显示声明数据 arr
var arr [6]int
arr[0] = 0
arr[1] = 1
...
arr[5] = 5

// 方式二：literal 方式
arr := [6]int{0, 1, 2, 3, 4, 5}
```

多维数据

```go
// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
doubleArray := [2][4]int{
        [4]int{1, 2, 3, 4}, 
        [4]int{5, 6, 7, 8},
    }
// 上面的声明可以简化，直接忽略内部的类型
easyArray := [2][4]int{
        {1, 2, 3, 4}, 
        {5, 6, 7, 8},
    }
```


以下为数组在Go和C中的主要区别。在Go中，

- 数组是值。将一个数组赋予另一个数组会复制其所有元素。
- 特别地，若将某个数组传入某个函数，它将接收到该数组的一份副本而非指针。
- 数组的大小是其类型的一部分。类型 [10]int 和 [20]int 是不同的。

数组为值的属性很有用，但代价高昂；若你想要C那样的行为和效率，你可以传递一个指向该数组的指针。

```go
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // 注意显式的取址操作
```

但这并不是Go的习惯用法，切片才是。


## slice

切片通过对数组进行封装，为数据序列提供了更通用、强大而方便的接口。 除了矩阵变换这类需要明确维度的情况外，Go中的大部分数组编程都是通过切片来完成的。

**切片保存了对底层数组的引用，若你将某个切片赋予另一个切片，它们会引用同一个数组。 若某个函数将一个切片作为参数传入，则它对该切片元素的修改对调用者而言同样可见， 这可以理解为传递了底层数组的指针。**


```go
var slice []int // 和声明 array 一样，只是少了长度，即「动态数组」，为引用类型
```

切片的几种声明方式

```go
// 方式一：普通显示声明
var slice []int = arr[1:4] // 其中可见元素为arr的第2个元素地址到第4个元素地址，但容量为第2个元素地址到最后一个元素地址

// 方式二：literal 方式
slice := []int{0, 1, 2, 3, 4, 5}
```


切片扩容

只要切片不超出底层数组的限制，它的长度就是可变的，只需将它赋予其自身的切片即可。 切片的容量可通过内建函数 cap 获得，它将给出该切片可取得的最大长度。 


## map

映射是方便而强大的内建数据结构，它可以关联不同类型的值。其键可以是任何相等性操作符支持的类型， 如整数、浮点数、复数、字符串、指针、接口（只要其动态类型支持相等性判断）、结构以及数组。 切片不能用作映射键，因为它们的相等性还未定义。与切片一样，映射也是引用类型。 若将映射传入函数中，并更改了该映射的内容，则此修改对调用者同样可见。

```go
var m make(map[keyType][valueType]) // 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。

// 修改或插入一个元素
m[key] = elem // 如果值是一个 nil 调用则会抛出异常

// 获得一个元素
elem = m[key]

// 删除一个元素
delete(m, key) // 如果 key 不存在，调用也不会发生什么，也不会有什么副作用

// 通过双赋值检测某个健是否存在
elem, ok = m[key]
// 如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
// 同样的，当从 map 中读取某个不存在的键时，结果是 map 的 **元素类型的零值**。
```


映射需要注意的几点：

- 字典无序
- 长度不定
- 支持 len()
- 非线程安全，需要使用 mutex lock


## new,make 操作

Go提供了两种分配原语，即内建函数 new 和 make。 它们所做的事情不同，所应用的类型也不同。它们可能会引起混淆，但规则却很简单。 让我们先来看看 new。这是个用来分配内存的内建函数， 但与其它语言中的同名函数不同，它不会初始化内存，只会将内存置零。 也就是说，new(T) 会为类型为 T 的新项分配已置零的内存空间， 并返回它的地址，也就是一个类型为 *T 的值。用Go的术语来说，它返回一个指针， 该指针指向新分配的，类型为 T 的零值。


```go
// 等价
new(File) <==> &File{}  //返回了一个指针，指向新分配的类型 File 的零值
```

内建函数 make(T, args) 的目的不同于 new(T)。它只用于创建切片、映射和信道，并返回类型为 T（而非 *T）的一个已初始化 （而非置零）的值。 出现这种用差异的原因在于，这三种类型本质上为引用数据类型，它们在使用前必须初始化。

```go
// 用户创建 slice、map、chan
make(T, args)
```

```go
// 会分配一个具有100个 int 的数组空间，接着创建一个长度为10， 容量为100并指向该数组中前10个元素的切片结构。
make([]int, 10, 100)
```

```go
var p *[]int = new([]int)       // 分配切片结构；*p == nil；基本没用
var v  []int = make([]int, 100) // 切片 v 现在引用了一个具有 100 个 int 元素的新数组

// 没必要的复杂：
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// 习惯用法：
v := make([]int, 100)
```

请记住，make 只适用于映射、切片和信道且不返回指针。若要获得明确的指针， 请使用 new 分配内存。


## 零值

```go
//零值是一种“变量未填充前”的默认值
int     0
int8    0
int32   0
int64   0
uint    0x0
rune    0 //rune的实际类型是 int32
byte    0x0 // byte的实际类型是 uint8
float32 0 //长度为 4 byte
float64 0 //长度为 8 byte
bool    false
string  ""
```


## 流程控制


**if/else**

```go
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```

**switch**

There are two forms: expression switches and type switches.

```go
// expression switches
// 每个 case 默认 break
// 左花括号 { 必须与 switch 处于同一行；
// 条件表达式不限制为常量或者整数；
// 单个case中，可以出现多个结果选项；
// 与C语言等规则相反，Go语言不需要用break来明确退出一个case；
// 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case；
// 可以不设定switch之后的条件表达式，在此种情况下，整个switch结构与多个if...else...的逻辑作用等同。
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // missing switch expression means "true"
case x < 0: return -x
default: return x
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}
```


```go
// type switches
switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
```

**for**

```go
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
	// testdata.a is never evaluated; len(testdata.a) is constant
	// i ranges from 0 to 6
	f(i)
}

var a [10]string
for i, s := range a {
	// type of i is int
	// type of s is string
	// s == a[i]
	g(i, s)
}

var key string
var val interface{}  // element type of m is assignable to val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// empty a channel
for range ch {}
```

```go
Range expression                          1st value          2nd value

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, <-chan E       element  e  E
```


**goto**

goto 语句被多数语言学者所反对，告诫不要使用。

**break/continue/fallthrough**

在实际的使用中，需要根据具体的逻辑目标、程序执行的时间和空间限制、代码的可读性、编译器的代码优化设定等多种因素，灵活组合。

## 函数

在 Go 语言中，函数的基本特点有：多值返回、可命名的结果形参、Defer语句等。函数可是一等的（first-class）公民，函数类型也是一等的数据类型。这是什么意思呢？

简单来说，这意味着函数不但可以用于封装代码、分割功能、解耦逻辑，还可以化身为普通的值，在其他函数间传递、赋予变量、做类型判断和转换等等，就像切片和字典的值那样。
而更深层次的含义就是：函数值可以由此成为能够被随意传播的独立逻辑组件（或者说功能模块）。

“函数是一等的公民”是函数式编程（functional programming）的重要特征。Go 语言在语言层面支持了函数式编程。

1. 接受其他的函数作为参数传入；
2. 把其他的函数作为结果返回。


```go
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {

	// 多返回值
	return value1, value2
}

// 比较下列用法
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func SumAndProduct(A, B int) (int, int) {
	return A+B, A*B
}
func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A+B
	Multiplied = A*B
	return
}

// 变长参数
// 变量 arg 是一个int的slice
func myfunc(arg ...int) {
    for _, n := range arg {
        fmt.Printf("And the number is: %d\n", n)
    }
}

// 传值，传指针
// 比较下列用法, 与 C 类似
func add1(a int) int {
	a = a+1 // 我们改变了a的值
	return a //返回一个新值
}
func add1(a *int) int { // 请注意，
	*a = *a+1 		    // 修改了 a 的值
	return *a 			// 返回该值
}
func main() {
	x := 3
	fmt.Println("x = ", x)    // 输出 "x = 3"
	x1 := add1(x)  			  // 调用 add1(x)
	fmt.Println("x+1 = ", x1) // 输出"x+1 = 4"
	fmt.Println("x = ", x)    // 输出"x = 3"
    
   	x := 3
	fmt.Println("x = ", x)    // 输出 "x = 3"
	x1 := add1(&x)  		  // 调用 add1(&x) 传x的地址
	fmt.Println("x+1 = ", x1) // 输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 输出 "x = 4"
}
```

```go
lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns

// prints 3 2 1 0 before surrounding function returns
for i := 0; i <= 3; i++ {
	defer fmt.Print(i)
}

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
```

## 函数类型与传值
 

 ```go

 // 声明了一个函数类型
 type backTrack(path []int) ([]int)
 ```

 
## panic 与 recover

panic 与 recover 机制替代了异常捕获 try/catch

```go
var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}
```

```go
// 调用 recover 将停止回溯过程，并返回传入 panic 的实参。 
// 由于在回溯时只有被推迟函数中的代码在运行，因此 recover 只能在被推迟的函数中才有效。

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() { // 
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work) // 若 do(work) 触发了Panic，其结果就会被记录， 而该Go程会被干净利落地结束，不会干扰到其它Go程。我们无需在推迟的闭包中做任何事情， recover 会处理好这一切。
}
```


## main 与 init 函数

- main 只能应用于package main
- init 函数 能够应用于所有的 package
- 定义时不能有任何的参数和返回值
- 程序会自动调用init()和main()
- 每个package中的init函数都是可选的，但package main就必须包含一个main函数。

![Go语言程序的初始化和执行顺序](../images/go-main-init-func.png)


## struct



## 面向对象



## 反射



## 并发


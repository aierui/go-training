package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	s := "abcdefg"

	for i := 0; i < len(s); i++ {
		fmt.Printf("s index ==> %T -> %v %p <(%+v)>  <(%#v)>  <=> %s \n", s[i], s[i], s[i], &s, s[i], s[i])
	}

	fmt.Println()

	str1 := "0123456789/;"

	for i := range str1 {
		fmt.Printf("index => %v strV => %v strS => %s\n", i, str1[i], str1[i])
	}
	fmt.Println()
	for i := 0; i <= 255; i++ {

		//fmt.Printf("index => %v convert string => %v %s %#U === %U \n", i, strconv.Itoa(i), byte(i), byte(i), byte(i))
		fmt.Printf("index => %v convert string => %v %s %#U === %U \n", i, strconv.Itoa(i), uint8(i), uint8(i), uint8(i))
	}

	//strings.Index(s[j:i], string(s[i]))

	//
	/**
	General
	%v 以默认的方式打印变量的值

	%T 打印变量的类型

	Integer
	%+d 带符号的整型，fmt.Printf("%+d", 255)输出+255
	%q 打印单引号
	%o 不带零的八进制
	%#o 带零的八进制
	%x 小写的十六进制
	%X 大写的十六进制
	%#x 带0x的十六进制
	%U 打印Unicode字符
	%#U 打印带字符的Unicode
	%b 打印整型的二进制


	%5d 表示该整型最大长度是5，下面这段代码

	  fmt.Printf("|%5d|", 1)
	  fmt.Printf("|%5d|", 1234567)
	输出结果如下：


	|    1|
	|1234567|


	%-5d则相反，打印结果会自动左对齐
	%05d会在数字前面补零。


	Float
	%f (=%.6f) 6位小数点
	%e (=%.6e) 6位小数点（科学计数法）
	%g 用最少的数字来表示
	%.3g 最多3位数字来表示
	%.3f 最多3位小数来表示

	String
	%s 正常输出字符串
	%q 字符串带双引号，字符串中的引号带转义符
	%#q 字符串带反引号，如果字符串内有反引号，就用双引号代替
	%x 将字符串转换为小写的16进制格式
	%X 将字符串转换为大写的16进制格式
	% x 带空格的16进制格式


	Struct
	%v 正常打印。比如：{sam {12345 67890}}
	%+v 带字段名称。比如：{name:sam phone:{mobile:12345 office:67890}
	%#v 用Go的语法打印。
	比如main.People{name:”sam”, phone:main.Phone{mobile:”12345”, office:”67890”}}


	Boolean
	%t 打印true或false


	Pointer
	%p 带0x的指针
	%#p 不带0x的指针

	作者：gowk
	链接：https://www.jianshu.com/p/8be8d36e779c
	来源：简书
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


	// https://golang-examples.tumblr.com/post/86795367134/fmtprintf-format-reference-cheat-sheet


	最后，对字符串类型做一个总结，希望能够帮助大家能够了解golang的字符串，并且能够在开发过程中正确处理它。

	golang的string是以UTF-8编码的,而UTF-8是一种1-4字节的可变长字符集，每个字符可用1-4字节 来表示

	使用下标方式s[i]访问字符串s，s[i]是UTF-8编码后的一个字节(uint8)，即按字节遍历

	使用for i,v := range s 方式访问s，i是字符串下标编号，v是对应的字符值(int32=rune)，即按字符遍历

	使用fmt.Printf打印时，%c占位符打印的是字符,%+v占位符打印的是这个类型自身，如fmt.Printf(“%+v”,s[i])打印的就是字节一个十进制的无符号整数s[i]

	如果希望以随机方式访问字符串s的每个字符，可以先转为[]rune数组，再以下标访问

	*/

}

func mergeSort(a, b []int) []int {
	c := append(a, b...)
	sort.Ints(c)
	str := "hello"

	fmt.Println(strconv.Itoa(int(str[1])))

	return c
}

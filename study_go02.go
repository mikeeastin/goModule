package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const pi = 3.1415
const (
	k1 = 100
	k2
	k3
)
const (
	a1 = iota
	a2
	a3
	a4
)
const (
	b1 = iota
	b2 = 100
	b3
	b4 = iota
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	Trace = log.New(ioutil.Discard, "TRACE:", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	fmt.Println("hello")

	// const  iota
	/*	fmt.Println(k2)
		fmt.Println(pi)
		fmt.Println(a1)
		fmt.Println(a2)
		fmt.Println(a3)
		fmt.Println(b1)
		fmt.Println(b2)
		fmt.Println(b3)
		fmt.Println(b4)*/

	// 日志
	/*s1 := "abcde.aaee"
	split := strings.Split(s1, ".")
	fmt.Println(split[0])
	fmt.Println(split[1])

	Trace.Println("i have something standard to say")
	Info.Println("sepecial Information")
	Warning.Println("there  is something you  need to know about")
	Error.Println("somehing has failed ")*/

	//数组
	/*	aaa := [6]int{1, 2, 3,4,5,6}
		bbb := []int{1, 2, 3}
		fmt.Printf("%T\n  %T\n", aaa, bbb)
		var total int
		for k,v :=range aaa{
			fmt.Println(k,v)
			total +=v
		}
		fmt.Println("总和",total)*/

	// 切片 引用类型
	/*s1 := make([]int, 2, 10)
	//fmt.Printf("%T", s1)
	s1[0] = 12
	s1 = append(s1, 23)
	s3 := make([]int, 3, 10)
	fmt.Println(s1)
	fmt.Println(s3)
	copy(s3, s1)
	fmt.Println(s3)
	s5:=append(s3,s3[:1]...)
	fmt.Println(s5)*/

	// 指针声明错误方式
	/*	var a2  *int
	*a2 = 12
	fmt.Println(a2)*/

	//指针声明正确方式
	/*var a1 = new(int) //返回的是对应类型的指针
	*a1 = 45
	fmt.Println(a1)
	a3 := make([]int, 2) // make 给slice map chan 申请内存的，make返回的是对应类型本身
	fmt.Println(a3)*/

	// map
	/*var map1 map[int]string
	map1 = make(map[int]string, 5)
	map1[1] = "aaa"
	map1[2] = "bbb"
	fmt.Println(map1)
	for k, v := range map1 {
		fmt.Println(k, v)
	}*/

	//元素类型为map的切片
	/*	var s1 = make([]map[int]string, 10, 10)
		s1[0] = make(map[int]string, 1)
		s1[0][0] = "bbbb"
		fmt.Println(s1)*/

	//值为切片类型的map
	/*var s2 = make(map[string][]int, 10)
	s2["aaa"] = []int{1, 2, 4}
	fmt.Println(s2)
	var s3 []int //没有分配内存
	//s3[0] = 23   // 此行会报错，内有分配内存不能直接赋值
	s3 = []int{1, 3, 4}
	fmt.Println(s3)
	s4 := make([]int, 5, 5) //已分配内存
	s4 = []int{5, 6, 7}
	fmt.Println(s4)

	var m2 map[string]int
	m2 = make(map[string]int, 10)
	m2["aaa"] = 2
	fmt.Println(m2)*/

	// 回文判断
	/*	s := "abc34cba"
		fmt.Println(s)
		huiwen := false
		var lenth int
		lenth = len(s)
		for k, _ := range s {
			fmt.Println(k, s[k])
			if s[k] == s[lenth-k-1] {
				huiwen = true
			} else {
				huiwen = false
				break
			}
			if k > lenth/2 {
				break
			}
		}
		if huiwen {
			fmt.Println("是", s)
		} else {
			fmt.Println("不是")
		}

	*/

	/*	f2 := func() {
			fmt.Println("annou func")
		}
		f2()*/

	// refer
	/*myname := myfunc()
	fmt.Printf("main 函数里的name: %s\n", name)
	fmt.Println("main 函数里的myname: ", myname)*/

// panic
/*	set_data(20)
	// 如果能执行到这句，说明panic被捕获了
	// 后续的程序能继续运行
	fmt.Println("everything is ok")*/
	// 实例化
	myself := Profile{name: "小明", age: 24, gender: "male"}
	// 调用函数
	myself.FmtProfile()
	fmt.Println("new name"+myself.name)
}

type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

// 定义一个与 Profile 的绑定的方法
//传递指针可修改属性，传递变量则不能修改
func (person *Profile) FmtProfile() {
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
	person.name="asdf"
}


var name string = "go"

func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

func set_data(x int) {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 故意制造数组越界，触发 panic
	var arr [10]int
	arr[x] = 88
}
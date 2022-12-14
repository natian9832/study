package main

import "fmt"

// defer
func func01() {
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	fmt.Println(f4())
}
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// a,b,c为闭包演示
func func02() {
	param := c(b, 1, 2)
	a(param)
}
func a(f func()) {
	fmt.Println("This is a")
	f()
}
func b(x, y int) {
	fmt.Println("This is b")
	fmt.Printf("x+y=%d\n", x+y)
}
func c(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func func03() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// panic,recover
func func04() {
	funcA()
	funcB()
	funcC()
}
func funcA() {
	println("funcA()")
}
func funcB() {
	//刚刚打开数据库连接
	defer func() {
		err := recover()
		fmt.Println(err)
		println("释放数据库连接")
	}()
	panic("出现了严重的错误!!!") //程序崩溃退出
	println("funcB()")
}
func funcC() {
	println("funcC()")
}

// fmt包的使用
func func05() {
	//m1 := make(map[string]int, 1)
	//m1["那天"] = 100
	//fmt.Printf("%T\n", m1)
	//fmt.Printf("%v\n", m1)
	//fmt.Printf("%#v\n", m1)

	//n := 10
	//fmt.Printf("%d%%\n", n)
	//fmt.Printf("%b\n", n)
	//fmt.Printf("%o\n", n)
	//fmt.Printf("%x\n", n)
	//fmt.Printf("%X\n", n)

	//flag := true
	//fmt.Printf("%t\n", flag)

	var (
		name string
		age  int
	)
	fmt.Print("请输入姓名和年龄:")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("name:%s,age:%d\n", name, age)
	fmt.Println(true)
}

func main() {
	//func01()
	//func02()
	//func03()
	//func04()
	func05()
}

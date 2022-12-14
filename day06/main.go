package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func init() {
	fmt.Println("导入main包时自动调用")
}

//type cat struct{}
//type dog struct{}
//type person struct{}
//// 接口类型
//type speaker interface {
//	speak()
//	//run()
//}
//
//func (c cat) speak() {
//	fmt.Println("喵喵喵～")
//}
//
//func (c cat) run() {
//	fmt.Println("猫在跑")
//}
//
//func (d dog) speak() {
//	fmt.Println("汪汪汪～")
//}
//
//func (p person) speak() {
//	fmt.Println("啊啊啊～")
//}
//
//func beat(x speaker) {
//	x.speak()
//	//x.run()
//}
//func testInterface() {
//	var c1 cat
//	//var d1 dog
//	//var p1 person
//	var s speaker
//	s = c1
//
//	//beat(c1)
//	//beat(d1)
//	//beat(p1)
//	beat(s)
//}

//type animal interface {
//	move()
//	eat(string)
//}
//
//type cat struct {
//	name string
//	feet int8
//}
//
//type chicken struct {
//	feet int8
//}
//
//func (c chicken) move() {
//	fmt.Println("鸡动!")
//}
//
//func (c chicken) eat(something string) {
//	fmt.Printf("鸡吃%s!\n", something)
//}
//
//func (c cat) move() {
//	fmt.Println("走猫步!")
//}
//
//func (c cat) eat(something string) {
//	fmt.Printf("猫吃%s!\n", something)
//}
//func testInterface() {
//	var a1 animal
//	//c1 := cat{"淘气", 4}
//	//chicken1 := chicken{2}
//	//a1 = c1
//	//a1.eat("鱼")
//	//a1.eat("那天")
//	//a1 = chicken1
//	//a1.move()
//	//a1.eat("稻谷")
//	fmt.Printf("%T\n", a1)
//	//fmt.Printf("%T\n", c1)
//	//fmt.Printf("%T\n", chicken1)
//}

//type animal interface {
//	move()
//	eat(something string)
//}
//type cat struct {
//	name string
//	feet int8
//}
//
////func (c cat) move() {
////	fmt.Println("走猫步...")
////}
////
////func (c cat) eat(food string) {
////	fmt.Printf("猫吃%s\n", food)
////}
//
//func (c *cat) move() {
//	fmt.Println("走猫步...")
//}
//
//func (c *cat) eat(food string) {
//	fmt.Printf("猫吃%s\n", food)
//}
//
//func testInterface() {
//	var a1 animal
//	c1 := cat{"Tom", 4}
//	c2 := &cat{"汤姆", 4}
//	a1 = c1
//	fmt.Println(a1)
//	a1 = c2
//	fmt.Println(a1)
//}

//type animal interface {
//	mover
//	eater
//}
//type mover interface {
//	move()
//}
//type eater interface {
//	eat(string2 string)
//}
//type cat struct {
//	name string
//	feet int8
//}
//
//func (c *cat) move() {
//	fmt.Println("走猫步...")
//}
//func (c *cat) eat(food string) {
//	fmt.Printf("猫吃%s\n", food)
//}
//func testInterface() {
//	var a animal
//	c := &cat{"Tom", 4}
//	a = c
//	a.move()
//}

//func show(a interface{}) {
//	fmt.Printf("type:%T\tvalue:%v\n", a, a)
//}
//func testInterface() {
//	var m1 map[string]interface{}
//	m1 = make(map[string]interface{}, 16)
//	m1["name"] = "那天"
//	m1["age"] = 26
//	m1["married"] = false
//	m1["hobby"] = [...]string{"Java", "Go", "Python"}
//	//fmt.Println(m1)
//
//	show(false)
//	show(1)
//	show("那天")
//	show(nil)
//	show(m1)
//}

//func assign(a interface{}) {
//	fmt.Printf("%T\n", a)
//	t := fmt.Sprintf("%T", a)
//	fmt.Printf("%v\n", t)
//	str, ok := a.(string)
//	fmt.Printf("ok:%v\n", ok)
//	if ok {
//		fmt.Printf("str:%s\n", str)
//	} else {
//		fmt.Println("a不是string类型")
//	}
//}
//func testInterface() {
//	assign(1)
//}
//
//func testPackage() {
//	fmt.Println(calc.Add(1, 2))
//	fmt.Println(calc.Sub(1, 2))
//}

// 文件操作
func testFile1() {
	file, err := os.Open("./day06/main.go")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err != nil {
			fmt.Println("读文件失败")
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			fmt.Println("文件读取完毕")
			break
		}
	}

}
func testFile2() {
	file, err := os.Open("./day06/main.go")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()
	for {
		//创建一个读对象
		reader := bufio.NewReader(file)
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("文件读取失败")
			return
		}
		fmt.Print(line)
	}

}

func main() {
	//testInterface()
	//testPackage()
	//testFile1()
	testFile2()
}

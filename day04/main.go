package main

import "fmt"

// 结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func (p *person) eat() {
	fmt.Println(p.name + "正在吃饭")
}

func func01() {
	var p1 person
	p1.name = "那天"
	p1.age = 25
	p1.gender = "男"
	p1.hobby = []string{"Java", "Python", "Go", "C++"}
	fmt.Println(p1)
	fmt.Println(p1.name)
}

// 匿名结构体
func func02() {
	var s struct {
		name string
		age  int
	}
	s.name = "那天"
	s.age = 25
	fmt.Printf("type:%T,vlaue:%v\n", s, s)
}

// 指针类型结构体
func func03() {
	p := person{"那天", 25, "男", []string{"JAVA", "PYTHON", "GO"}}
	fmt.Println(p)
	//funcA(p)
	funcB(&p)
	fmt.Println(p)
}

// 传值
func funcA(p person) {
	p.age = 100
}

// 传指针
func funcB(p *person) {
	//语法糖和(*p).age一样
	p.age = 101
}

// 成员函数和结构体初始化
func func04() {
	//p := person{"那天", 25, "男", []string{"JAVA", "PYTHON", "GO", "C++"}}

	p := person{name: "那天", age: 25}

	//var p person
	//p.name = "那天"
	p.eat()
	fmt.Println(p)
}

// 继承
func func05() {

}

func main() {
	//func01()
	//func02()
	//func03()
	func04()
}

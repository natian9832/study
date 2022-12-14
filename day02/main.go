package main

import "fmt"

// 01.数组
func func01() {
	//var a [3]bool
	//for index, value := range a {
	//	println(index, value)
	//}

	//a = [3]bool{true, false, true}
	//fmt.Println(a)

	//根据后面推断长度
	//a := [...]int{1, 2, 3, 4, 5}
	//fmt.Println(a)

	//后续补0
	//a := [5]int{1, 2}
	//fmt.Println(a)

	//根据索引初始化
	//a := [5]int{0: 1, 4: 5}
	//fmt.Println(a)

	//数组遍历
	//cities := [...]string{"北京", "南京", "和县"}
	//for i := 0; i < len(cities); i++ {
	//	fmt.Println(cities[i])
	//}
	//for index, city := range cities {
	//	fmt.Println(index, city)
	//}

	//多维数组及遍历
	//var a [3][3]int
	//a = [3][3]int{
	//	[3]int{1, 2, 3},
	//	[3]int{4, 5, 6},
	//	[3]int{7, 8, 9},
	//}
	//for i := 0; i < len(a); i++ {
	//	for j := 0; j < len(a[i]); j++ {
	//		fmt.Printf("%d\t", a[i][j])
	//	}
	//	fmt.Println()
	//}

	//数组是值类型
	a := [3]int{1, 2, 3}
	b := a
	a[0] = 100
	fmt.Println(a, b) //[100 2 3] [1 2 3]
}

// 02.切片
func func02() {
	//切片的定义
	var a []int
	a = []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := a
	a[0] = 100
	fmt.Println(b)
}

// 指针
func func03() {
	//n := 10
	//fmt.Println(&n)
	//p := &n
	//fmt.Printf("%T\n", p)
	//m := *p
	//fmt.Println(m)
	//fmt.Printf("%T\n", m)
	//x := &p
	//fmt.Printf("%T\n", x)
	//fmt.Println(**x)

	//var a *int
	////*a = 100
	//fmt.Println(a)

	////new关键字
	//var a = new(int)
	//*a = 100
	//fmt.Println(a)

}

// map
func func04() {
	var m1 map[string]string
	//fmt.Println(m1 == nil)
	//要先初始化
	m1 = make(map[string]string, 10)
	m1["name"] = "那天"
	m1["age"] = "25"
	//fmt.Println(m1)
	//value, ok := m1["addr"]
	//if ok {
	//	fmt.Println(value)
	//} else {
	//	fmt.Println("这个值不存在")
	//}
	//for key, value := range m1 {
	//	fmt.Println(key, value)
	//}
	//delete(m1, "age")
	//fmt.Println(m1)
}

// 求数组的和
func test01(arr [5]int) (result int) {
	for _, v := range arr {
		result += v
	}
	return
}

func main() {
	//func01()
	//func02()
	//func03()
	func04()
}

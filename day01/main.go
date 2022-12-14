package main

import "fmt"

// Func01 string
func Func01() {
	//s1 := "\"a\\b\\c\""
	//fmt.Println(s1)
	//
	//s2 := `
	//	欲把心事付瑶琴
	//	知音少
	//	弦断有谁听
	//`
	//fmt.Println(s2)
	//
	////字符串相关操作
	//s3 := "hello"
	//fmt.Printf("s3的长度为:%d\n", len(s3))
	//
	//s4 := "那"
	//s5 := "天"
	//fmt.Println(s4 + s5)
	//s6 := s4 + s5
	//fmt.Println(s6)
	//s7 := fmt.Sprintf("%s%s", s4, s5)
	//fmt.Println(s7)
	//
	//s8 := "a.b.c.d.e.f.g"
	//split := strings.Split(s8, ".")
	//fmt.Println(split)
	//
	//contains := strings.Contains(s8, "c")
	//fmt.Println(contains)
	//
	//s9 := "012345"
	//println(strings.HasPrefix(s9, "12"))
	//println(strings.HasSuffix(s9, "4"))
	//
	//println(strings.Index(s9, "34"))
	//println(strings.Index(s9, "14"))
	//
	//println(strings.Join(split, "+"))

	//s1 := "那天123"
	//for _, c := range s1 {
	//	fmt.Printf("%c\t", c)
	//	fmt.Printf("type:%T\n", c)
	//}
	//
	//s2 := "那天"
	//s3 := []rune(s2) //[那 天]
	//s3[0] = '酒'
	//println(string(s3))
	//
	//c1 := "天"
	//c2 := '天' //rune int32
	//fmt.Printf("%T\n", c1)
	//fmt.Printf("%T\n", c2)
	//c3 := "a"
	//c4 := byte('a') //byte int8
	//fmt.Printf("%T\n", c3)
	//fmt.Printf("%T\n", c4)
}

// Func02 类型转化
func Func02() {
	n := 10
	var f float64
	f = float64(n)
	println(f)

}

// Func03 if/for语句
func Func03() {
	//for i := 0; ; i += 2 {
	//	println(i)
	//	if i == 10 {
	//		break
	//	}
	//}

	//i := 0
	//for i < 10 {
	//	println(i)
	//	i += 3
	//}

	s := "hello那天"
	println(len(s))
	for index, value := range s {
		//println(index, value)
		fmt.Printf("%d,%c\n", index, value)
	}
}

// 统计一句话中汉字的个数
func test01(s string) (count int) {
	//count := 0
	for _, c := range s {
		if c < '0' || c > 'z' {
			count++
			//fmt.Printf("%c\n", c)
		}
	}
	return
}

// 打印九九乘法表
func test02() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		println()
	}
}

func main() {
	//Func01()
	//Func02()
	//Func03()
	test02()
}

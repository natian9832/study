package main

func testTime() {
	//now := time.Now()
	//fmt.Printf("current time : %v\n", now)
	//
	//year := now.Year()
	//month := now.Month()
	//day := now.Day()
	//hour := now.Hour()
	//minute := now.Minute()
	//second := now.Second()
	//fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//时间戳
	//fmt.Println(now.Unix())
	//fmt.Println(now.UnixNano())

	//时间戳转时间
	//ret := time.Unix(1627378000, 0)
	//fmt.Println(ret)
	//
	//fmt.Println(time.Second)

	//now + 1小时
	//fmt.Println(now.Add(time.Hour * 24))

	//定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t)
	//}

	//时间格式化(把语言中的时间对象转化为字符串)
	//fmt.Println(now.Format("2006-01-02"))
	//fmt.Println(now.Format("2006/01/02 15:04:05"))
	//fmt.Println(now.Format("2006/01/02 03:04:05"))
	//fmt.Println(now.Format("2006/01/02 03:04:05.000"))

	//把字符串时间转化为时间戳
	//timeObj, err := time.Parse("2006/01/02", "1998/03/02")
	//if err != nil {
	//	fmt.Printf("你传入的时间格式不正确:%v\n", err)
	//} else {
	//	fmt.Println(timeObj)
	//	fmt.Println(timeObj.UnixMilli())
	//}

}

func testLog() {

}
func main() {
	testTime()
	//testLog()
}

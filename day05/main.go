package main

import "fmt"

// 函数版学生管理系统
// 能够查看、新增、删除学生信息
type student struct {
	id   int64
	name string
}

// student构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

var (
	allStudent map[int64]*student
)

func showAllStudent() {
	if len(allStudent) == 0 {
		fmt.Println("学生列表为空,请去添加学生信息~")
		return
	}
	for _, value := range allStudent {
		fmt.Printf("学号:%d,姓名:%s\n", value.id, value.name)
	}
}
func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	if allStudent[id] != nil {
		fmt.Println("你输入学号的学生已经在系统中~")
		return
	}
	fmt.Print("请输入学生的姓名:")
	fmt.Scanln(&name)
	stu := newStudent(id, name)
	allStudent[id] = stu
	fmt.Println("添加成功~")

}
func deleteStudent() {
	var (
		id   int64
		flag string
	)
	fmt.Print("请输入你要删除的学生id:")
	fmt.Scanln(&id)
	if allStudent[id] == nil {
		fmt.Println("你要删除的用户不存在~")
		return
	}
	fmt.Printf("你确定要删除学号为[%d]的学生y/n?", id)
	fmt.Scanln(&flag)
	if flag == "y" || flag == "Y" || flag == "yes" || flag == "Yes" {
		delete(allStudent, id)
		fmt.Println("删除成功~")
	}
}
func showMenu() {
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.显示菜单
		5.退出系统
	`)
}

func main() {
	//初始化开辟内存空间
	allStudent = make(map[int64]*student, 48)
	count := 0
	//1.打印菜单
	fmt.Println("欢迎光临学生管理系统")
	showMenu()
	for true {
	label:
		fmt.Printf("[%d]请输入你的操作:", count)
		count++
		var choice int
		//2.等待用户选择
		fmt.Scanln(&choice)
		//3.执行对应的函数
		if choice == 1 {
			showAllStudent()
		} else if choice == 2 {
			addStudent()
		} else if choice == 3 {
			deleteStudent()
		} else if choice == 4 {
			showMenu()
		} else if choice == 5 {
			var flag string
			fmt.Print("你确定要退出系统y/n?")
			fmt.Scanln(&flag)
			if flag == "y" || flag == "Y" || flag == "yes" || flag == "Yes" {
				break
			}
		} else {
			fmt.Println("你输入的序号不正确,请重新输入!!!")
			goto label
		}
	}
	fmt.Println("bye~")
}

package main

import "fmt"

type User struct {
	Username string
	Password string
}

var registers []User

func register() {
	var u User
	fmt.Println("请输入帐号")
	fmt.Scan(&u.Username)
	fmt.Println("请输入密码")
	fmt.Scan(&u.Password)
	ok := RegisterCheck(u.Username, u.Password)
	if ok {
		fmt.Println("注册成功")
		registers = append(registers, u)
	}
}

func RegisterCheck(username, password string) bool {
	for i := 0; i < len(registers); i++ {
		if registers[i].Username == username {
			fmt.Println("用户名已被注册")
			return false
		}
	}
	if len(password) < 6 {
		fmt.Println("密码不能小于六位")
		return false
	}
	return true
}
func login() {
	var u User
	fmt.Println("请输入帐号")
	fmt.Scan(&u.Username)
	fmt.Println("请输入密码")
	fmt.Scan(&u.Password)
	ok := LoginCheck(u.Username, u.Password)
	if ok {
		fmt.Println("登录成功")
	}
}
func LoginCheck(username, password string) bool {
	flag := false
	RealPassword := ""
	for i := 0; i < len(registers); i++ {
		if registers[i].Username == username {
			flag = true
			RealPassword = registers[i].Password

		}
	}
	if !flag {
		fmt.Println("不存在此用户")
		return false
	}
	if RealPassword != password {
		fmt.Println("密码错误")
		return false
	}
	return true
}
func main() {
	var option int
	fmt.Println("欢迎使用本系统")
	fmt.Println("1.注册功能 2.登录功能 3.退出功能")

	//我们简单的使用1 2 3来模拟注册 登录 退出操作
	for {
		fmt.Scan(&option)
		if option == 3 {
			fmt.Println("退出成功")
			break
		} else if option == 2 {
			login()
		} else if option == 1 {
			register()
		}
	}

}

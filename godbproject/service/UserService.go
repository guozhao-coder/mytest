package service

import (
	"fmt"
	"godbproject/dao"
	"godbproject/entity"
)

func Insert(){
	//实例化结构体
	var u entity.Users
	fmt.Println("输入员工信息 姓名，密码")
	//获取键盘输入
	var n string
	var p int
	fmt.Scanln(&n)
	fmt.Scanln(&p)
	u.Username = n
	u.Password = p
	userinfo := new(dao.Userinfo)
	userinfo.InsertUser(u)
}
func Delete()  {
	fmt.Println("请输入要删除的员工id")
	var i int
	fmt.Scanln(&i)
	userinfo := new(dao.Userinfo)
	userinfo.DeleteUser(i)
}
func Update(){
	//实例化结构体
	var u entity.Users
	fmt.Println("请输入要修改的员工id")
	var i int
	fmt.Scanln(&i)
	fmt.Println("请输入新的姓名，密码")
	var n string
	var p int
	fmt.Scanln(&n)
	fmt.Scanln(&p)
	u.Username=n
	u.Password=p

	userinfo := new(dao.Userinfo)
	userinfo.UpdateUser(i,u)

}
func Query()  {
	userinfo := new(dao.Userinfo)
	userinfo.QueryUser()
}


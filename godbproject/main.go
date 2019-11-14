package main

import (
	"fmt"
	"godbproject/service"
)

func main() {
	/*
	//实例化结构体
	//var u entity.Users

	//
	//ins :=new(dao.Userinfo)

	//测试增加
	//ins.InsertUser(u)

	//测试删除
	//ins.DeleteUser(3)

	//测试修改
	//ins.UpdateUser(4,u)

	//测试查询
	//ins.QueryUser()

	 */
	fmt.Println("欢迎进入管理系统")
	for {
		fmt.Println("请选择功能：1增加员工；2删除员工 ；3根据id修改员工信息 ；4查询所有员工 ；5退出")
		var i int
		fmt.Scanln(&i)
		switch i {
		case 1:
			service.Insert()
		case 2:
			service.Delete()
		case 3:
			service.Update()
		case 4:
			service.Query()
		case 5:
			goto Zaijian
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}

	Zaijian:
	   fmt.Println("再见！")
}

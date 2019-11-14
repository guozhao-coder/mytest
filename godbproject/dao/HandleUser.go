package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"godbproject/entity"
)

type HandleUser interface {
	InsertUser(users entity.Users)
	DeleteUser(id int)
	UpdateUser(id int,users entity.Users)
	QueryUser() []entity.Users
}

type Userinfo struct {
}


//增加
func (i *Userinfo) InsertUser(users entity.Users) {
	db, e := sql.Open("mysql", "root:guozhao@/mygo?charset=utf8")
	CheckErr(e)
	stmt,e := db.Prepare("insert into users values (default ,?,?)")
	CheckErr(e)
	_, e = stmt.Exec(users.Username, users.Password)
	CheckErr(e)
	fmt.Println("插入成功")
}
//删除
func (d *Userinfo) DeleteUser(id int){
	db, e := sql.Open("mysql", "root:guozhao@/mygo?charset=utf8")
	CheckErr(e)
	stmt, e := db.Prepare("delete from users where id =?")
	CheckErr(e)
	_, e = stmt.Exec(id)
	CheckErr(e)
	fmt.Println("删除成功")
}
//修改
func (u *Userinfo) UpdateUser(id int,users entity.Users) {
	db, e := sql.Open("mysql", "root:guozhao@/mygo?charset=utf8")
	CheckErr(e)
	stmt, e := db.Prepare("update users set username = ?, password = ? where id = ?")
	CheckErr(e)
	_, e = stmt.Exec(users.Username, users.Password, id)
	CheckErr(e)
	fmt.Println("修改成功")
}
func (q *Userinfo) QueryUser() []entity.Users {
	//定义一个users类型的切片
	var us []entity.Users
	db, e := sql.Open("mysql", "root:guozhao@/mygo?charset=utf8")
	CheckErr(e)
	stmt, e := db.Query("select * from users")
	CheckErr(e)
	//遍历结果集
	for stmt.Next(){
		var u entity.Users
		e := stmt.Scan(&u.Id,&u.Username,&u.Password)
		CheckErr(e)
		fmt.Printf("id:%d, name:%s ,pwd:%d\n",u.Id,u.Username,u.Password)
		//追加到切片
		us = append(us ,u)
	}
	return us
}

func CheckErr(e error)  {
	if e != nil{
		panic(e)
	}
}

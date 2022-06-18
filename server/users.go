package server

import (
	"database/sql"
	"errors"
	"fmt"
	"mssgserver/utils"
)

type User struct {
	Id   int `db:"id"`
	Age int `db:"age"`
	RealName string `db:"real_name"`
}

// 插入数据
func (u *User) InsertRowDemo() error {
	sqlStr := "insert into user(real_name, age) values (?,?)"
	ret, err := utils.DB.Exec(sqlStr, "常见1", 193)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return err
	}
	fmt.Printf("insert success, the id is %d.\n", theID)

	return nil
}

// 单个查询
func (u *User) SelectRowDemo(id int) error {
	sqlStr1 := `select id, real_name,age from user where id=?`
	err := utils.DB.Get(u,sqlStr1,id)
	if errors.Is(err,sql.ErrNoRows) {
		return err
	}
	return err
}

// 返回列表
func (u *User) SelectListDemo() (user []User) {
	sqlStr1 := `select id, real_name,age from user`
	utils.DB.Select(&user,sqlStr1)
	return
}
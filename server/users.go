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
	err := utils.DB.Select(&user,sqlStr1)
	if errors.Is(err,sql.ErrNoRows) {
		return
	}
	return
}

//更新
func (u *User) UpdateRowDemo(id int, age int)  (i int64) {
	sqlStr := "update user set age=? where id > ?"
	ret, err := utils.DB.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return //默认返回 返回值的零值
	}
	i, err1 := ret.RowsAffected() // 操作影响的行数
	if err1 != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err1)
		return
	}

	return
}

// 删除数据
func  (u *User) DeleteRowDemo(id int) (b bool) {
	sqlStr1 := `select id from user where id=?`
	err := utils.DB.Get(u,sqlStr1,id)
	if errors.Is(err,sql.ErrNoRows) {
		return
	}

	sqlStr := "delete from user where id = ?"
	ret, err := utils.DB.Exec(sqlStr, id)
	if  err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	_, err = ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	return true
}
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123qwe@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("open failed,err:%v\n", err)
		return
	}
	fmt.Println("连接中。。")

	//尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Printf("connet failed,err:%v\n", err)
		return
	}
	fmt.Println("连接成功")
}

package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // 或者其他驱动
)

var db3 *sqlx.DB

func init() {
	db3, _ = sqlx.Open("sqlite3", "test3.db")
	db3.Exec("DROP TABLE IF EXISTS employees")
	// 建表
	db3.Exec("CREATE TABLE IF NOT EXISTS employees (id INTEGER PRIMARY KEY, name TEXT, department TEXT, salary INTEGER)")
	// 插入测试数据
	db3.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "张三", "技术部", 10000)
	db3.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "李四", "技术部", 20000)
	db3.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "王五", "产品部", 30000)
}

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

/*
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中
*/
func main() {
	// 查询部门为 "技术部" 的员工信息
	var employees []Employee
	err := db3.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		panic(err)
	}
	for _, emp := range employees {
		fmt.Printf("%+v\n", emp)
	}
	// 查询工资最高的员工信息
	var emp Employee
	err = db3.Get(&emp, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", emp)
}

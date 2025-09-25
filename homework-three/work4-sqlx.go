package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // 或者其他驱动
)

var db4 *sqlx.DB

func init() {
	var err error
	db4, err = sqlx.Connect("sqlite3", "test4.db")
	if err != nil {
		panic(err)
	}
	// 删除表
	db4.MustExec("DROP TABLE IF EXISTS books")
	// 建表
	db4.MustExec("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT, price REAL)")
	// 插入数据
	db4.MustExec("INSERT INTO books (title, author, price) VALUES (?,?,?)", "Go 语言", "老王", 88.88)
	db4.MustExec("INSERT INTO books (title, author, price) VALUES (?,?,?)", "Python 编程", "老张", 66.66)
	db4.MustExec("INSERT INTO books (title, author, price) VALUES (?,?,?)", "Java 编程", "老刘", 77.77)

}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

/*
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全
*/
func main() {
	var books []Book
	err := db4.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", 50)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		fmt.Printf("%+v\n", book)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

type DB struct {
	db *sql.DB
}

func newDB() *DB {
	db, err := sql.Open("sqlite3", "test1.db")
	if err != nil {
		log.Fatalf("连接数据库失败：%v", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS students (id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"name TEXT, age INTEGER, grade TEXT)")
	if err != nil {
		log.Fatalf("“创建 students 表失败”：%v", err)
	}
	return &DB{db: db}
}

func (d *DB) Close() {
	d.db.Close()
}

func (d *DB) trucateStudents() {
	d.db.Exec("DROP TABLE students")
}

func (d *DB) InsertStudent(s *student) error {
	stmt, err := d.db.Prepare("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("“准备插入语句失败”：%v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.Name, s.Age, s.Grade)
	if err != nil {
		return fmt.Errorf("“插入新记录失败”：%v", err)
	}
	return nil
}

func (d *DB) UpdateStudent(s *student) error {
	stmt, err := d.db.Prepare("UPDATE students SET name = ?, age = ?, grade = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("“准备更新语句失败”：%v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.Name, s.Age, s.Grade, s.ID)
	if err != nil {
		return fmt.Errorf("“更新记录失败”：%v", err)
	}
	return nil
}

func (d *DB) DeleteStudentBySql(sql string, args ...interface{}) error {
	stmt, err := d.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("“准备删除语句失败”：%v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		return fmt.Errorf("“删除记录失败”：%v", err)
	}
	return nil
}

func (d *DB) QueryStudents(sql string, args ...interface{}) ([]*student, error) {
	rows, err := d.db.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("“查询 students 表失败”：%v", err)
	}
	defer rows.Close()
	students := make([]*student, 0)
	for rows.Next() {
		var id int
		var name string
		var age int
		var grade string
		err = rows.Scan(&id, &name, &age, &grade)
		if err != nil {
			return nil, fmt.Errorf("“扫描查询结果失败”：%v", err)
		}
		students = append(students, &student{ID: id, Name: name, Age: age, Grade: grade})
	}
	return students, nil
}

/*
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
func main() {
	db := newDB()
	// 插入新记录
	err := db.InsertStudent(&student{Name: "张三", Age: 20, Grade: "三年级"})
	if err != nil {
		log.Fatal(err)
	}
	// 查询年龄大于 18 岁的学生信息
	students, _ := db.QueryStudents("SELECT * FROM students WHERE age > ?", 18)
	for _, s := range students {
		fmt.Printf("%+v\n", s)
	}
	// 更新年级
	err = db.UpdateStudent(&student{ID: 1, Name: "张三", Age: 20, Grade: "四年级"})
	if err != nil {
		log.Fatal(err)
	}
	students, _ = db.QueryStudents("SELECT * FROM students WHERE age > ?", 18)
	for _, s := range students {
		fmt.Printf("%+v\n", s)
	}
	// 删除年龄小于 15 岁的学生记录
	err = db.DeleteStudentBySql("DELETE FROM students WHERE age < ?", 15)
	if err != nil {
		log.Fatal(err)
	}
	// 清空表(测试用)
	db.trucateStudents()
}

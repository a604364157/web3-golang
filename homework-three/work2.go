package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB2 struct {
	db *sql.DB
}

var db DB2

func init() {
	db.db, _ = sql.Open("sqlite3", "test2.db")
	db.db.Exec("DROP TABLE IF EXISTS accounts")
	db.db.Exec("DROP TABLE IF EXISTS transactions")
	db.db.Exec("CREATE TABLE IF NOT EXISTS accounts (id INTEGER PRIMARY KEY, balance INTEGER)")
	db.db.Exec("CREATE TABLE IF NOT EXISTS transactions (id INTEGER PRIMARY KEY, from_account_id INTEGER, to_account_id INTEGER, amount INTEGER)")
	// 初始化账户 A
	db.db.Exec("INSERT INTO accounts (id, balance) VALUES (?,?)", 1, 1000)
	// 初始化账户 B
	db.db.Exec("INSERT INTO accounts (id, balance) VALUES (?,?)", 2, 2000)
}

type Account struct {
	ID      int
	Balance int
}

type Transaction struct {
	ID            int
	FromAccountID int
	ToAccountID   int
	Amount        int
}

func transferTest01() {
	// 开启事务
	tx, _ := db.db.Begin()
	// 未知错误，回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	// 查询账户 A 的余额
	var a Account
	err := tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", 1).Scan(&a.Balance)
	if err != nil {
		panic(err)
	}
	if a.Balance < 100 {
		panic("账户 A 余额不足，无法转账")
	}
	// 转账先扣除账户 A 的余额，再增加账户 B 的余额
	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", 100, 1)
	if err != nil {
		panic("扣除账户 A 余额失败:" + err.Error())
	}
	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", 100, 2)
	if err != nil {
		panic("增加账户 B 余额失败:" + err.Error())
	}
	// 记录转账信息
	_, err = tx.Exec("INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (?, ?, ?)", 1, 2, 100)
	if err != nil {
		panic("记录转账信息失败:" + err.Error())
	}
	if err = tx.Commit(); err != nil {
		panic("提交事务失败:" + err.Error())
	}

	// 数据验证
	rows, _ := db.db.Query("SELECT * FROM accounts WHERE id = ?", 1)
	defer rows.Close()
	for rows.Next() {
		var a Account
		rows.Scan(&a.ID, &a.Balance)
		println("账户 A 余额:", a.Balance)
	}
	rows, _ = db.db.Query("SELECT * FROM accounts WHERE id = ?", 2)
	defer rows.Close()
	for rows.Next() {
		var a Account
		rows.Scan(&a.ID, &a.Balance)
		println("账户 B 余额:", a.Balance)
	}
	rows, err = db.db.Query("SELECT * FROM transactions")
	defer rows.Close()
	for rows.Next() {
		var t Transaction
		rows.Scan(&t.ID, &t.FromAccountID, &t.ToAccountID, &t.Amount)
		println("转账记录: ID:", t.ID, " 转出账户ID:", t.FromAccountID, " 转入账户ID:", t.ToAccountID, " 转账金额:", t.Amount)
	}

}

/*
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）
和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务
*/
func main() {
	transferTest01()
}

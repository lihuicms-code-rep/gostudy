package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Go操作MySQL

var db *sql.DB   //全局对象,DB是是一个数据库（操作）句柄，代表一个具有零到多个底层连接的连接池
var xdb *sqlx.DB //sqlx的DB句柄

type User struct {
	ID   int64  `db:"id"`   //userID
	Name string `db:"name"` //name
	Age  int32  `db:"age"`  //age
}

func InitDB() (err error) {
	//Data Source Name
	dsn := "root:Lh123#@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn) //Open只验证参数,并不创建连接
	if err != nil {
		fmt.Printf("open mysql error:%v\n", err)
		return err
	}

	//尝试与数据库建立连接
	err = db.Ping()
	if err != nil {
		fmt.Printf("ping to mysql error:%v\n", err)
		return err
	}

	return nil
}

//查询操作示例

//单行查询, func (db *DB) QueryRow(query string, args ...interface{}) *Row
func QueryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var user User

	//确保QueryRow之后调用Scan,才会释放连接
	err := db.QueryRow(sqlStr, 1).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("query row scan error %v\n", err)
		return
	}

	fmt.Printf("query row result user:%v\n", user)
}

//多行查询, func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
func QueryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query rows error %v", err)
		return
	}

	//注意,关闭row释放连接
	defer rows.Close()

	//遍历结果集
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("rows scan error:%v", err)
			continue
		}

		fmt.Printf("user:%v\n", u)
	}
}

//插入数据,func (db *DB) Exec(query string, args ...interface{}) (Result, error)
func InsertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "詹姆斯", 33)
	if err != nil {
		fmt.Printf("insert into user error:%v\n", err)
		return
	}

	insId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert Id error:%v\n", err)
		return
	}

	fmt.Printf("insert into user success, lastId:%d\n", insId)
}

//更新数据,也是使用Exec
func UpdateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

//删除数据,也是使用Exec
func DeleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

//MySQL预处理

//查询预处理
func PrepareQueryDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare error:%v\n", err)
		return
	}
	defer stmt.Close()

	//根据返回的状态用于之后的查询
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("state query error:%v\n", err)
		return
	}

	defer rows.Close()

	//循环读取数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("rows scan error:%v", err)
			continue
		}

		fmt.Printf("user:%v\n", u)
	}

}

//插入预处理
func PrepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("沙河娜扎", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

//事务示例
func TransactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	_, err = tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	sqlStr2 := "Update user set age=40 where id=?"
	_, err = tx.Exec(sqlStr2, 4)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}

//sqlx包的使用
func InitDBWithSQLX() error {
	dsn := "root:Lh123#@tcp(127.0.0.1:3306)/sql_test"
	xdb, _ = sqlx.Connect("mysql", dsn)
	db.SetMaxOpenConns(20) //设置与数据库建立连接的最大数目
	db.SetMaxIdleConns(10) //设置连接池中的最大闲置连接数
	return nil
}

//查询
// 查询单条数据示例
func queryRowXDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u User
	err := xdb.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowXDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []User
	err := xdb.Select(&users, sqlStr, 0) //多条数据的查询就凸显出便捷了,直接将结构返回给具体的结构
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

//插入,修改,删除基本与原生sql中一致,都是使用db.Exec
// 插入数据
func insertRowXDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := xdb.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowXDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := xdb.Exec(sqlStr, 39, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowXDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := xdb.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}


//sqlx中的事务操作,使用db.Beginx()和tx.MustExec()来简化之前sql.Exec后的错误处理Rollback()
func transactionXDemo() {
	tx, err := xdb.Beginx() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=40 where id=?"
	tx.MustExec(sqlStr1, 2)
	sqlStr2 := "Update user set age=50 where id=?"
	tx.MustExec(sqlStr2, 4)
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}


//SQL注入

func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	var users []User
	err := xdb.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	err := InitDB()
	if err != nil {
		fmt.Printf("InitDB error:%v\n", err)
		return
	}

	fmt.Println("InitDB success")

	err = InitDBWithSQLX()
	if err != nil {
		fmt.Printf("InitDBWithSQLX error:%v\n", err)
		return
	}
	fmt.Println("InitDBWithSQLX success")

	//InsertRowDemo()
	//QueryMultiRowDemo()
	//UpdateRowDemo()
	//DeleteRowDemo()
	//PrepareInsertDemo()
	//TransactionDemo()
	//PrepareQueryDemo()
	//queryRowXDemo()
	//queryMultiRowXDemo()

	//sqlInjectDemo("xxx' or 1=1#")
	//sqlInjectDemo("xxx' union select * from user #")
	//sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
}

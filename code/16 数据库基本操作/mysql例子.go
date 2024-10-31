package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局的 db 变量
var db *sql.DB

// user 结构体用于映射数据库字段
type user struct {
	id   int
	age  int
	name string
}

// 初始化 MySQL 连接
func initMySQL() (err error) {
	// DSN格式: 用户名:密码@tcp(IP:端口)/数据库名
	dsn := "root:123456@tcp(127.0.0.1:3306)/test" // 修改为你的实际数据库名
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试与数据库建立连接
	err = db.Ping()
	if err != nil {
		return err
	}

	// 设置连接池参数
	db.SetMaxOpenConns(100) // 设置最大连接数
	db.SetMaxIdleConns(10)  // 设置最大空闲连接数

	// 确保表存在
	err = ensureTableExists()
	if err != nil {
		return err
	}

	return nil
}

// 确保表存在
func ensureTableExists() error {
	// 检查表是否存在的查询
	checkTableSQL := `
        SELECT COUNT(*)
        FROM information_schema.TABLES 
        WHERE TABLE_SCHEMA = 'test'   
        AND TABLE_NAME = 'user_test'
    `
	var count int
	err := db.QueryRow(checkTableSQL).Scan(&count)
	if err != nil {
		return fmt.Errorf("检查表是否存在失败: %v", err)
	}

	// 如果表不存在，创建表
	if count == 0 {
		createTableSQL := `
            CREATE TABLE user_test (
                id BIGINT(20) NOT NULL AUTO_INCREMENT,
                name VARCHAR(20) DEFAULT '',
                age INT(11) DEFAULT '0',
                PRIMARY KEY(id)
            ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4
        `
		_, err = db.Exec(createTableSQL)
		if err != nil {
			return fmt.Errorf("创建表失败: %v", err)
		}
		fmt.Println("表 user_test 创建成功")
	}
	return nil
}

// 查询单条数据
func queryRowDemo() (*user, error) {
	sqlStr := "SELECT id, name, age FROM user_test WHERE id = ?"
	var u user
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// 查询多条数据
func queryMultiRowDemo() {
	sqlStr := "SELECT id, name, age FROM user_test WHERE id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("查询数据失败，err:%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("扫描数据失败, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入数据
func insertRowDemo() {
	sqlStr := "INSERT INTO user_test(name, age) VALUES(?, ?)"
	result, err := db.Exec(sqlStr, "小明", 22)
	if err != nil {
		fmt.Printf("插入数据失败, err:%v\n", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("获取插入ID失败, err:%v\n", err)
		return
	}
	fmt.Printf("插入数据成功, id:%d\n", id)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "UPDATE user_test SET age = ? WHERE id = ?"
	result, err := db.Exec(sqlStr, 23, 1)
	if err != nil {
		fmt.Printf("更新数据失败, err:%v\n", err)
		return
	}

	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("获取受影响行数失败, err:%v\n", err)
		return
	}
	fmt.Printf("更新数据成功, 受影响的行数:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "DELETE FROM user_test WHERE id = ?"
	result, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("删除数据失败, err:%v\n", err)
		return
	}

	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("获取受影响行数失败, err:%v\n", err)
		return
	}
	fmt.Printf("删除数据成功, 受影响的行数:%d\n", n)
}

func main() {
	// 初始化 MySQL 连接
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	// 程序退出前关闭数据库连接
	defer db.Close()

	// 测试各种数据库操作
	fmt.Println("=== 插入数据 ===")
	insertRowDemo()

	fmt.Println("\n=== 查询单条数据 ===")
	u1, err := queryRowDemo()
	if err != nil {
		fmt.Printf("查询失败：%v\n", err)
	} else {
		fmt.Printf("查询结果: id:%d, name:%s, age:%d\n", u1.id, u1.name, u1.age)
	}

	fmt.Println("\n=== 查询多条数据 ===")
	queryMultiRowDemo()

	fmt.Println("\n=== 更新数据 ===")
	updateRowDemo()

	fmt.Println("\n=== 删除数据 ===")
	deleteRowDemo()
}

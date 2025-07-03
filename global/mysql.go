package global

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

var (
	Mdb *sql.DB // 全局数据库连接池变量，初始值为 nil
)

func InitMysql() error {
	// 读取 MYSQL 环境变量
	dsn := os.Getenv("MYSQL")
	if dsn == "" {
		return fmt.Errorf("MYSQL: 环境变量未设置")
	}
	// 检查是否包含 /
	if !strings.Contains(dsn, "/") {
		return fmt.Errorf("MYSQL: 环境变量格式错误")
	}

	var err error
	// 连接数据库
	Mdb, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("MySQL 连接失败: %v", err)
	}

	// 设置连接池
	Mdb.SetMaxOpenConns(25)                 // 连接池最大连接数
	Mdb.SetMaxIdleConns(25)                 // 连接池最大空闲连接数
	Mdb.SetConnMaxLifetime(5 * time.Minute) // 连接池最大连接时间

	// 测试连接
	if err := Mdb.Ping(); err != nil {
		// mysqlErr 是具体错误, ok是mysqlErr是否属于MySQLError类型, 即断言是否成功的布尔值
		// 即 ok: 断言成功,是MySQLError错误. mysqlErr: 具体的错误 mysqlErr.Number: 错误码
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1049 {
			fmt.Println("MySQL数据库不存在 , 尝试创建数据库")

			Mdb.Close()

			// 重连MySQL
			Mdb, err = sql.Open("mysql", strings.Split(dsn, "/")[0]+"/")
			if err != nil {
				fmt.Println("MySQL 连接失败")
				return fmt.Errorf("MySQL 连接失败: %v", err)
			}

			// 创建
			_, err = Mdb.Exec(fmt.Sprintf("CREATE DATABASE %s", strings.Split(dsn, "/")[1]))
			if err != nil {
				return fmt.Errorf("创建数据库失败: %v", err)
			}

			fmt.Printf("数据库 %s 创建成功", strings.Split(dsn, "/")[1])

			Mdb.Close()

			// 再次重连
			Mdb, err = sql.Open("mysql", dsn)
			if err != nil {
				fmt.Println("MySQL 连接失败")
				return fmt.Errorf("MySQL 连接失败: %v", err)
			}

			// 设置连接池
			Mdb.SetMaxOpenConns(25)
			Mdb.SetMaxIdleConns(25)
			Mdb.SetConnMaxLifetime(5 * time.Minute)

			fmt.Println("MySQL 连接成功")

			return nil

		}
		return fmt.Errorf("ping MySQL 失败: %v", err)
	}
	fmt.Println("MySQL 连接成功")

	return nil
}

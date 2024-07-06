package database

import (
	"log"
	"os"

	"github.com/JeaMianh/MyBlog/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 创建一个 gorm.DB 类型的指针，保存连接信息
// 一个高级别的连接对象
var DBconn *gorm.DB

func init() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// 被调用时建立连接
func ConnectDB() {

	// 获取环境变量
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	charset := os.Getenv("DB_CHARSET")
	parseTime := os.Getenv("DB_PARSE_TIME")
	loc := os.Getenv("DB_LOC")

	// 构建 DSN
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=" + charset + "&parseTime=" + parseTime + "&loc=" + loc

	// 下面两行是从 gorm 的文档里复制过来的，dsn 包含了连接到数据库所需的信息，一个很长的字符串。
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn 数据源名称，包含数据库连接信息，如用户名、密码、主机地址、端口和数据库名称。

	// gorm.Open() 返回一个 *gorm.DB 类型的变量，赋值给 db
	// mysql.Open() 是 mysql 驱动里的一个函数，用于解析 dsn
	// &gorm.Config{}，GORM 的配置结构体，用于配置 GORM 的行为
	// 这里只将日志级别配置为 error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connection failed.")
	}

	log.Print("Connection successful.")

	// 根据传入的模型结构体，自动更新或创建数据库表结构
	// 在开发和首次部署使用，在生产环境频繁使用会影响性能
	// 不能执行删除操作
	// AutoMigrate() 传入一个结构体指针
	// new() 传入一个类型，返回一个指向改类型的指针
	db.AutoMigrate(new(model.Blog))

	DBconn = db
}

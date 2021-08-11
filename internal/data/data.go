package data

import (
	"context"
	"fmt"
	"ginapi/internal/conf"
	"ginapi/internal/data/model"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserData,
	NewUserRepo,
	NewGormClient,
	NewRedisClient,
)

// Data .
type Data struct {
	// TODO warpped database client
	db    *gorm.DB
	redis *redis.Client
	log   *zap.Logger
}

type UserData struct {
	db    *gorm.DB
	redis *redis.Client
	log   *zap.Logger
}

func NewGormClient(config *conf.Conf) (*gorm.DB, func(), error) {
	dsn := config.Conf.GetString("data.database.source")
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	dbOpen, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		panic(err)
	}
	db, err := dbOpen.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(config.Conf.GetInt("data.database.max_idle_conn"))
	db.SetMaxOpenConns(config.Conf.GetInt("data.database.max_open_conn"))
	// 迁移 schema
	dbOpen.AutoMigrate(model.User{})
	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Println("close mysql err:", err)
		}
		log.Println("close mysql success")
	}
	log.Println("open mysql success")
	return dbOpen, cleanup, nil
}

// NewData .
func NewData(
	db *gorm.DB,
	redis *redis.Client,
	log *zap.Logger,
) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{db: db, redis: redis}, cleanup, nil
}

func NewUserData(
	db *gorm.DB,
	redis *redis.Client,
	log *zap.Logger,
) (*UserData, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &UserData{db: db, redis: redis, log: log}, cleanup, nil
}

func NewRedisClient(config *conf.Conf) (*redis.Client, func(), error) {
	var ctx = context.Background()
	fmt.Println(fmt.Sprintf("%s:%s", config.Conf.GetString("data.redis.addr"), config.Conf.GetString("data.redis.port")))
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Conf.GetString("data.redis.addr"), config.Conf.GetString("data.redis.port")),
		Username:     config.Conf.GetString("username"),
		Password:     config.Conf.GetString("password"),
		DB:           config.Conf.GetInt("default_db"), // use default DB
		DialTimeout:  config.Conf.GetDuration("dial_timeout"),
		ReadTimeout:  config.Conf.GetDuration("read_timeout"),
		WriteTimeout: config.Conf.GetDuration("write_timeout"),
	})
	_, err := rdb.Ping(ctx).Result()
	cleanup := func() {
		if err := rdb.Close(); err != nil {
			log.Println("close redis err:", err)
		}
		log.Println("close redis success")
	}
	if err != nil {
		log.Println("redis open error:", err)
		return nil, cleanup, err
	}
	log.Println("redis open success")
	return rdb, cleanup, nil
}

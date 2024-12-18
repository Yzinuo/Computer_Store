package comstore

import (
	g "computer_store/internal/global"
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitLogger(conf *g.Config) {
	var level slog.Level
	switch conf.Log.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelDebug
	}

	option := &slog.HandlerOptions{
		AddSource: false, //是否显示源文件的信息
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}

	var handler slog.Handler
	switch conf.Log.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, option)
	case "text":
		handler = slog.NewTextHandler(os.Stdout, option)
	default:
		handler = slog.NewTextHandler(os.Stdout, option)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger) // 设定默认的slog
}

func InitDatabase(conf *g.Config) *gorm.DB {
	dbType := conf.DbType()
	dsn := conf.DbDSN()
	
	var level logger.LogLevel
	switch conf.Log.Level {
	case "silent":
		level = logger.Silent
	case "error":
		level = logger.Error
	case "warn":
		level = logger.Warn
	case "info":
		level = logger.Info
	default:
		level = logger.Error
	}
	
	config := &gorm.Config{
		Logger: logger.Default.LogMode(level),
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	var db *gorm.DB
	var err error
	switch dbType {
	case "mysql":
		db,err = gorm.Open(mysql.Open(dsn),config)
	default:
		log.Fatal("不支持的数据库类型")
	}

	if err != nil {
		log.Fatal("数据库连接失败",err)
	}
	slog.Info("数据库连接成功")

	return db
}

func InitRedis(conf *g.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB: conf.Redis.DB,
	})

	_,err := rdb.Ping(context.Background()).Result()
	if err!= nil {
		log.Fatal("redis连接失败",err)
	}

	slog.Info("redis连接成功")
	return rdb
}
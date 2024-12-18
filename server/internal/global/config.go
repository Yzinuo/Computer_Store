package global

import (
	"fmt"
	"log"
)

type Config struct {
	Server struct {
		Mode          string // debug | release
		Port          string
		DbType        string // mysql | sqlite
		DbAutoMigrate bool   // 是否自动迁移数据库表结构
		DbLogMode     string // silent | error | warn | info
	}
	Log struct {
		Level     string // debug | info | warn | error
		Prefix    string
		Format    string // text | json
		Directory string
	}
	JWT struct {
		Secret string
		Expire int64 // hour
		Issuer string
	}
	Mysql struct {
		Host     string // 服务器地址
		Port     string // 端口
		Config   string // 高级配置
		Dbname   string // 数据库名
		Username string // 数据库用户名
		Password string // 数据库密码
	}
	SQLite struct {
		Dsn string // Data Source Name
	}
	Redis struct {
		DB       int    // 指定 Redis 数据库
		Addr     string // 服务器地址:端口
		Password string // 密码
	}
	Session struct {
		Name   string
		Salt   string
		MaxAge int
	}
	Email struct {
		From     string // 发件人 要发邮件的邮箱
		Host     string // 服务器地址, 例如 smtp.qq.com 前往要发邮件的邮箱查看其 smtp 协议
		Port     int    // 前往要发邮件的邮箱查看其 smtp 协议端口, 大多为 465
		SmtpPass string // 邮箱密钥 不是密码是开启smtp后给你的密钥
		SmtpUser string // 邮箱账号
	}
	Captcha struct {
		SendEmail  bool // 是否通过邮箱发送验证码
		ExpireTime int  // 过期时间
	}
	Upload struct {
		// Size      int    // 文件上传的最大值
		OssType   string // local | qiniu|aliyun
		Path      string // 本地文件访问路径
		StorePath string // 本地文件存储路径
	}
	Qiniu struct {
		ImgPath       string // 外链链接
		Zone          string // 存储区域
		Bucket        string // 空间名称
		AccessKey     string // 秘钥AK
		SecretKey     string // 秘钥SK
		UseHTTPS      bool   // 是否使用https
		UseCdnDomains bool   // 上传是否使用 CDN 上传加速
	}
	Aliyun struct {
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
		Bucket          string
		ImgPath         string
	}
}

var Conf *Config

func GetConfig() *Config {
	if Conf == nil {
		log.Panic("配置还没有初始化")
		return nil
	}
	return Conf
}

//func ReadConfig(path string) *Config {
//	v := viper.New()
//	v.SetConfigFile(path)
//	v.AutomaticEnv()                                   // 如果配置文件中有环境变量
//	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // 环境变量中分割符号都是下划线
//
//	if err := v.ReadInConfig(); err != nil {
//		panic("读取配置文件失败" + err.Error())
//	}
//	if err := v.Unmarshal(&Conf); err != nil {
//		panic("unmarsh 解析配置文件失败" + err.Error())
//	}
//	log.Println("配置初始化成功:" + path)
//
//	return Conf
//}

func ReadConfig(path string) {
	fmt.Printf("-----------")
}

func (*Config) DbType() string {
	if Conf.Server.DbType == "" {
		Conf.Server.DbType = "sqlite"
	}
	return Conf.Server.DbType
}

func (*Config) DbDSN() string {
	switch Conf.Server.DbType {
	case "mysql":
		conf := Conf.Mysql
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?%s",
			conf.Username, conf.Password, conf.Host, conf.Port, conf.Dbname, conf.Config,
		)
	case "sqlite":
		return Conf.SQLite.Dsn
	// 默认使用 sqlite, 并且使用内存数据库
	default:
		Conf.Server.DbType = "sqlite"
		if Conf.SQLite.Dsn == "" {
			Conf.SQLite.Dsn = "file::memory:"
		}
		return Conf.SQLite.Dsn
	}
}
package configs

import (
	"log"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// 全局配置的对象
var AllConfig ServerConfig

// 全局配置的结构体
type ServerConfig struct {
	Basic  BasicConfig  `mapstructure:"basic"`  // 基础配置
	Auth   AuthConfig   `mapstructure:"auth"`   // 认证配置
	Logger LoggerConfig `mapstructure:"logger"` // 日志配置
	Pager  PagerConfig  `mapstructure:"pager"`  // 分页配置
	RDS    RDSConfig    `mapstructure:"rds"`    // 关系型数据库配置
	Cache  CacheConfig  `mapstructure:"cache"`  // 缓存配置
}

// 基础配置
type BasicConfig struct {
	URL        string `mapstructure:"url"`         // URL
	Host       string `mapstructure:"host"`        // 绑定地址
	Port       int    `mapstructure:"port"`        // 绑定端口
	Mode       string `mapstructure:"mode"`        // 运行模式：[debug、test、release]
	OnlineTime string `mapstructure:"online_time"` // 上线时间
}

// 认证配置
type AuthConfig struct {
	JWTExpire int64  `mapstructure:"jwt_expire"` // JWT Token 过期时间，单位为秒
	JWTSecret string `mapstructure:"jwt_secret"` // JWT 密钥
}

// 日志配置
type LoggerConfig struct {
	Level         string `mapstructure:"level"`          // 日志级别
	Directory     string `mapstructure:"directory"`      // 日志目录
	RotationTime  int    `mapstructure:"rotation_time"`  // 日志轮换时间间隔，单位为小时
	RotationCount uint   `mapstructure:"rotation_count"` // 日志轮换文件保留个数
}

// 分页配置
type PagerConfig struct {
	Limit  int64 `mapstructure:"limit"`  // 每页条目数，默认值
	Offset int64 `mapstructure:"offset"` // 偏移量，默认值
}

// 关系型数据库配置
type RDSConfig struct {
	DriverName      string           `mapstructure:"driver_name"`       // 驱动名称
	Debug           bool             `mapstructure:"debug"`             // 调试模式
	MySQL           MySQLConfig      `mapstructure:"mysql"`             // MySQL 相关配置
	PostgreSQL      PostgreSQLConfig `mapstructure:"postgresql"`        // PostgreSQL 相关配置
	MaxIdleConns    int              `mapstructure:"max_idle_conns"`    // 空闲连接池中连接的最大数量
	MaxOpenConns    int              `mapstructure:"max_open_conns"`    // 打开数据库连接的最大数量
	ConnMaxLifetime int              `mapstructure:"conn_max_lifetime"` // 连接可复用的最大时间，单位为小时
}

// MySQL 配置
type MySQLConfig struct {
	Host     string `mapstructure:"host"`     // 访问地址
	Port     int    `mapstructure:"port"`     // 访问端口
	Username string `mapstructure:"username"` // 用户名
	Password string `mapstructure:"password"` // 密码
	Database string `mapstructure:"database"` // 数据库
}

// PostgreSQL 配置
type PostgreSQLConfig struct {
	Host     string `mapstructure:"host"`     // 访问地址
	Port     int    `mapstructure:"port"`     // 访问端口
	Username string `mapstructure:"username"` // 用户名
	Password string `mapstructure:"password"` // 密码
	Database string `mapstructure:"database"` // 数据库
}

// Cache 配置
type CacheConfig struct {
	DriverName string      `mapstructure:"driver_name"` // 驱动名称
	Redis      RedisConfig `mapstructure:"redis"`       // Redis 相关配置
}

// Redis 配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`      // 访问地址
	Port     int    `mapstructure:"port"`      // 访问端口
	Password string `mapstructure:"password"`  // 密码
	DB       int    `mapstructure:"db"`        // 数据库
	PoolSize int    `mapstructure:"pool_size"` // 连接池大小
}

// 解析配置文件
func ParseConfigFile(configFile string) error {
	// 指定配置文件路径
	viper.SetConfigFile(configFile)
	// 指定配置文件格式
	viper.SetConfigType("yaml")

	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 反序列化
	if err := viper.Unmarshal(&AllConfig); err != nil {
		return err
	}

	// 动态加载配置文件：监视配置文件
	viper.WatchConfig()
	// 配置文件发生更改
	viper.OnConfigChange(func(event fsnotify.Event) {
		if event.Op == fsnotify.Write {
			// 反序列化
			if err := viper.Unmarshal(&AllConfig); err != nil {
				log.Printf("unable to unmarshal config, err: %e\n", err)
			}
			log.Println("the configuration file has changed")
		}
	})
	return nil
}

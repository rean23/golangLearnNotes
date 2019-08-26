package configs

type Config struct {
	Debug bool //调试模式
	Database Connections	//数据库连接
}

var env = Config{
	Debug:    true,
	Database: Connection,
}

func GetConfig() *Config {
	return &env
}

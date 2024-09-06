package config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}
type ServerConfig struct {
	SecretKey         string
	ExpirationMinutes int
	ExpirationHours   int
}
type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
}

var Cfg Config

func (config Config) Init(loadedConfig map[string]string) {
	Cfg.Server = ServerConfig{
		SecretKey:         loadedConfig["SECRET_KEY"],
		ExpirationMinutes: 50,
		ExpirationHours:   24,
	}
	Cfg.Database = DatabaseConfig{
		Host:     loadedConfig["DB_HOST"],
		Username: loadedConfig["DB_USERNAME"],
		Password: loadedConfig["DB_PASSWORD"],
		DBName:   loadedConfig["DB_NAME"],
		Port:     loadedConfig["DB_PORT"],
	}
}

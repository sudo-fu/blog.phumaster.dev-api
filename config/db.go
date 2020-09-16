package config

// DBConfig nil
type DBConfig struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
}

// DB nil
func DB() *DBConfig {
	return &DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		DB:       "blog",
		User:     "root",
		Password: "default",
	}
}

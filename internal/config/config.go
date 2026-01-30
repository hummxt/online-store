package config

type Config struct {
	Server struct {
		Port string `mapstructure:"PORT"`
		Mode string `mapstructure:"MODE"`
	} `mapstructure:",squash"`

	DB struct {
		Host     string `mapstructure:"DB_HOST"`
		Port     string `mapstructure:"DB_PORT"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		Name     string `mapstructure:"DB_NAME"`
		SSLMode  string `mapstructure:"DB_SSLMODE"`
	} `mapstructure:",squash"`

	Redis struct {
		Addr     string `mapstructure:"REDIS_ADDR"`
		Password string `mapstructure:"REDIS_PASSWORD"`
		DB       int    `mapstructure:"REDIS_DB"`
	} `mapstructure:",squash"`

	JWT struct {
		Secret    string `mapstructure:"JWT_SECRET"`
		ExpiresIn string `mapstructure:"JWT_EXPIRES_IN"`
	} `mapstructure:",squash"`
}

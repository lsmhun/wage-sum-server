package configuration

type Config struct {
	DbType     string `env:"WAGESUM_DB_TYPE" envDefault:"postgres"`
	DbHost     string `env:"WAGESUM_DB_HOST" envDefault:"127.0.0.1"`
	DbPort     string `env:"WAGESUM_DB_PORT" envDefault:"5432"`
	DbName     string `env:"WAGESUM_DB_NAME" envDefault:"wagesum"`
	DbUsername string `env:"WAGESUM_DB_USERNAME" envDefault:"postgres"`
	DbPassword string `env:"WAGESUM_DB_PASSWORD,unset" envDefault:"mysecretpassword"`

	HttpServerPort string `env:"WAGESUM_HTTP_SERVER_PORT" envDefault:"3000"`
}

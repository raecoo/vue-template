package api

type Configuration struct {
	DbHost      string `env:"DB_HOST" env-description:"Database host address" env-default:"localhost"`
	DbPort      string `env:"DB_PORT" env-description:"Database host port" env-default:"5432"`
	DbUserName  string `env:"DB_USERNAME" env-description:"Database username" env-default:"postgres"`
	DbPassword  string `env:"DB_PASSWORD" env-description:"Database password"`
	LogFileName string `env:"LOG_FILE_NAME" env-description:"Service log file name" env-default:"api-access.log"`
	AuthSecret  string `env:"API_AUTH_SECRET" env-description:"API authentication secret" env-required:"true"`
}

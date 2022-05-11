package entity

type AppConfig struct {
	AppEnv   string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Url string
}

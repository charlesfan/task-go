package config

type Redis struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

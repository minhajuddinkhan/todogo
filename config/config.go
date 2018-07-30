package config

type Db struct {
	Host             string
	Port             string
	SSL              bool
	Name             string
	Username         string
	Password         string
	Dialect          string
	ConnectionString string
}
type Configuration struct {
	JWTSecret string
	Port      string
	Db        Db
}

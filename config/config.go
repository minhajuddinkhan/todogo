package config

//Db Db
type Db struct {
	Host             string
	Port             string
	SSL              bool
	Name             string
	Username         string
	Password         string
	Dialect          string
	ConnectionString string
	VolumePath       string //For local memory dbs
}

//Configuration Configuration
type Configuration struct {
	JWTSecret string
	Port      string
	Db        Db
}

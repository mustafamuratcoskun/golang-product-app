package postgresql

type Config struct {
	Host                  string
	Port                  string
	UserName              string
	Password              string
	DbName                string
	MaxConnections        string
	MaxConnectionIdleTime string
}

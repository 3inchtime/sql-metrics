package sql_metrics

type Config struct {
	Username    string
	Password    string
	Server      string
	Port        string
	MaxOpenConn int
	MaxIdleConn int
	MaxLifetime int
}

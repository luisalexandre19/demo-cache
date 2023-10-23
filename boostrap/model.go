package boostrap

type AppConfig struct {
	Port                  string
	Identification        string
	BussinesContainerAddr string
	Provider              string
}

type RedisConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Context      string
	Db           int
	Ttl          int //In minutes
	PoolSize     int
	MaxConnAge   int //In seconds
	IdleTimeout  int //In seconds
	ReadTimeout  int
	WriteTimeout int
	DialTimeout  int
	FailFast     bool
}

package config

var configInstance *config

type config struct {
	hostPostgres     string
	portPostgres     int
	userPostgres     string
	passwordPostgres string
	dbNamePostgres   string
	jwtSecret        string
	developmentMode  bool
	passwordCost     int

	//-----------------
	redisHost string
	redisPort int
}

func newConfig() *config {
	return &config{
		hostPostgres:     "localhost",
		portPostgres:     5432,
		userPostgres:     "postgres",
		passwordPostgres: "postgres",
		dbNamePostgres:   "mydatabase",
		jwtSecret:        "43453gdgerge#%$$%¨FGYwwfwFAFAOUQI1314141",
		passwordCost:     12,
		redisHost:        "localhost",
		redisPort:        6379,
	}
}

func GetConfig() *config {
	if configInstance == nil {
		configInstance = newConfig()
	}
	return configInstance
}

func (c *config) GetHost() string {
	return c.hostPostgres
}

func (c *config) GetPort() int {
	return c.portPostgres
}

func (c *config) GetUser() string {
	return c.userPostgres
}

func (c *config) GetPassword() string {
	return c.passwordPostgres
}

func (c *config) GetDBName() string {
	return c.dbNamePostgres
}

func (c *config) GetJWTSecret() string {
	return c.jwtSecret
}

func (c *config) IsDevelopmentMode() bool {
	return c.developmentMode
}

func (c *config) GetPasswordCost() int {
	return c.passwordCost
}

func (c *config) GetRedisHost() string {
	return c.redisHost
}

func (c *config) GetRedisPort() int {
	return c.redisPort
}

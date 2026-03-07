package config

var configInstance *config

type config struct {
	host            string
	port            int
	user            string
	password        string
	dbName          string
	jwtSecret       string
	developmentMode bool
	passwordCost    int
}

func newConfig() *config {
	return &config{
		host:         "localhost",
		port:         5432,
		user:         "postgres",
		password:     "postgres",
		dbName:       "mydatabase",
		jwtSecret:    "43453gdgerge#%$$%¨FGYwwfwFAFAOUQI1314141",
		passwordCost: 12,
	}
}

func GetConfig() *config {
	if configInstance == nil {
		configInstance = newConfig()
	}
	return configInstance
}

func (c *config) GetHost() string {
	return c.host
}

func (c *config) GetPort() int {
	return c.port
}

func (c *config) GetUser() string {
	return c.user
}

func (c *config) GetPassword() string {
	return c.password
}

func (c *config) GetDBName() string {
	return c.dbName
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

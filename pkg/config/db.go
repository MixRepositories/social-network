package config

func GetDbConfig() string {
	conf := New()
	return conf.DbUsername + ":" + conf.DbPassword + "@(" + conf.DbHost + ")/" + conf.DbName
}

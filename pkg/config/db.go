package config

func GetDbConfig() string {
	conf := New()
	return conf.DbUsername + ":" + conf.DbPassword + "@(localhost:3306)/" + conf.DbName
}

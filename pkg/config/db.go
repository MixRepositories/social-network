package config

func GetDbConfig() string {
	conf := New()
	return conf.DbUsername + ":" + conf.DbPassword + "@(mysql:3306)/" + conf.DbName
}

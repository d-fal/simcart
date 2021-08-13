package config

// database struct
type Database struct {
	Username    string `yaml:"postgres.username"`
	Password    string `yaml:"postgres.password"`
	Host        string `yaml:"postgres.host"`
	Schema      string `yaml:"postgres.schema"`
	Automigrate bool   `yaml:"postgres.automigrate"`
	Logger      bool   `yaml:"postgres.logger"`
	Namespace   string
}

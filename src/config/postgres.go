package config

type PostgresConfig struct {
	Host string
	Port int
	User string
	Password string
	Dbname string
	Sslmode string
	SearchPath string
}

func GetPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Host: "host.docker.internal",	
		Port: 5432,
		User: "postgres",
		Password: "password",
		Dbname: "postgres",
		Sslmode: "disable",
		SearchPath: "go_practice",
	}
}
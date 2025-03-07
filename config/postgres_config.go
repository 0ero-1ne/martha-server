package config

import "fmt"

type PostgresConfig struct {
	Hostname string `ini:"hostname"`
	Port     string `ini:"port"`
	Dbname   string `ini:"dbname"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func (postgres *PostgresConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		postgres.Username,
		postgres.Password,
		postgres.Hostname,
		postgres.Port,
		postgres.Dbname,
	)
}

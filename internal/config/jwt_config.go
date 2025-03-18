package config

type JWTConfig struct {
	Secret string `ini:"secret"`
}

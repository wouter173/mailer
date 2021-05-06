package structs

type Config struct {
	Keys []Key `toml:"keys"`
}

type Key struct {
	Service string `toml:"service"`
	Token   string `toml:"token"`
}

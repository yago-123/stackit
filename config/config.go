package config

type Config struct {
	// todo(): put all config options here

}

type StoreConfig struct {
	// todo(): storage related config options
}

type APIConfig struct {
	// todo(): API related config options
}

func New() *Config {
	return &Config{}
}

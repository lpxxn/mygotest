package config

type CommonConfig struct {
	Server ConfigServer
}

type ConfigServer struct {
	Addr string
}

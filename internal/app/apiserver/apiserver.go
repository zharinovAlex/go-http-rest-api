package apiserver

type APIServer struct {
	config *Config
}

func New(conf *Config) *APIServer {
	return &APIServer{
		config: conf,
	}
}

func (server *APIServer) Start() error {
	return nil
}

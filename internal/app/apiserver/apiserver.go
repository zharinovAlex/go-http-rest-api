package apiserver

import "github.com/sirupsen/logrus"

type APIServer struct {
	config *Config
	logger *logrus.Logger
}

func New(conf *Config) *APIServer {
	return &APIServer{
		config: conf,
		logger: logrus.New(),
	}
}

func (server *APIServer) Start() error {
	if err := server.configLogger(); err != nil {
		return err
	}

	server.logger.Info("Starting API server")

	return nil
}

func (server *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}

	server.logger.SetLevel(level)
	return nil
}

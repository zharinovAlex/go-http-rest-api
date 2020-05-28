package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(conf *Config) *APIServer {
	return &APIServer{
		config: conf,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	if err := server.configLogger(); err != nil {
		return err
	}

	server.configRouter()

	server.logger.Info("Starting API server")

	return http.ListenAndServe(server.config.BindAddress, server.router)
}

func (server *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}

	server.logger.SetLevel(level)
	return nil
}

func (server *APIServer) configRouter() {
	server.router.HandleFunc("/status", server.handleStatusCheck())
}

func (server *APIServer) handleStatusCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	}
}

package v1

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yago-123/stackit/config"
	"github.com/yago-123/stackit/pkg/storage"

	"github.com/gin-gonic/gin"
)

// todo(): these should be provided via config.Config package or env vars
const (
	ServerReadTimeout  = 5 * time.Second
	ServerWriteTimeout = 5 * time.Second
	ServerIdleTimeout  = 10 * time.Second
	MaxHeaderBytes     = 1 << 20
)

// todo(): rename to something more descriptive
type ServerAPI struct {
	cfg    *config.Config
	store  storage.Store
	server *http.Server
}

func New(cfg *config.Config, store storage.Store) *ServerAPI {
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", "localhost", 8080), // todo(): replace with cfg
		Handler:        setupRouter(store),
		ReadTimeout:    ServerReadTimeout,
		WriteTimeout:   ServerWriteTimeout,
		IdleTimeout:    ServerIdleTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
	}

	return &ServerAPI{
		cfg:    cfg,
		store:  store,
		server: server,
	}
}

func (s *ServerAPI) Start() error {
	err := s.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		// todo(): use s.cfg.Logger instead
		log.Printf("HTTP API memory stopped successfully")
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *ServerAPI) Stop() error {
	return nil
}

func setupRouter(store storage.Store) *gin.Engine {
	router := gin.Default() // todo(): replace with gin.New()
	handlr := newHandler(store)

	router.POST("/api/v1/user", handlr.PostUser)
	router.GET("/api/v1/users", handlr.GetUsers)

	return router
}

package server

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IServer interface {
	gin.IRouter
	io.Closer
}

type Server struct {
	*gin.Engine

	httpServer *http.Server
	options    *Options
	listener   net.Listener
}

func New(options *Options) (svr *Server, err error) {
	engine := gin.New()

	listener, err := net.Listen("tcp", options.config.ListenAddress)
	if err != nil {
		return nil, err
	}

	httpServer := &http.Server{
		Handler: engine,
	}

	server := &Server{
		Engine:     engine,
		options:    options,
		listener:   listener,
		httpServer: httpServer,
	}

	if options.logger != nil {
		options.logger.Info("running", zap.String("module", "server"), zap.String("addr", options.config.ListenAddress))
	}

	if options.health != nil {
		server.options.health.SetSystemState("server", true)
	}

	go httpServer.Serve(listener)

	return server, nil
}

func (svr *Server) Close() error {
	if svr.options.logger != nil {
		svr.options.logger.Info("shutdown notice")
	}

	if svr.options.health != nil {
		svr.options.health.SetSystemState("server", false)
	}

	if svr.options.config.WarningDuration > time.Duration(0) {
		time.Sleep(svr.options.config.WarningDuration)
	}

	if svr.options.logger != nil {
		svr.options.logger.Info("shutdown start")
	}

	ctx := context.Background()
	if svr.options.config.ShutdownDuration > time.Duration(0) {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, svr.options.config.ShutdownDuration)
		defer cancel()
	}

	if svr.options.logger != nil {
		svr.options.logger.Info("shutdown")
	}

	return svr.httpServer.Shutdown(ctx)
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

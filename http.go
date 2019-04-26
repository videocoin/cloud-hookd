package hookd

import (
	"context"

	manager_v1 "github.com/VideoCoin/cloud-api/manager/v1"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// HTTPServerConfig addresses for http server
type HTTPServerConfig struct {
	Addr               string
	UserProfileRPCADDR string
	CamerasRPCADDR     string
	ManagerRPCADDR     string
}

// HTTPServer http server reciver
// holds echo, config, and log objects
type HTTPServer struct {
	cfg  *HTTPServerConfig
	e    *echo.Echo
	log  *logrus.Entry
	hook *Hook
}

// NewHTTPServer returns reference to new HTTPServer object
func NewHTTPServer(cfg *HTTPServerConfig, log *logrus.Entry) (*HTTPServer, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	managerConn, err := grpc.Dial(cfg.ManagerRPCADDR, opts...)
	if err != nil {
		return nil, err
	}

	manager := manager_v1.NewManagerServiceClient(managerConn)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	status, err := manager.Health(context.Background(), &empty.Empty{})
	if status.GetStatus() != "healthy" || err != nil {
		panic(err)
	}

	hook, err := NewHook(
		e,
		"/hook",
		manager,
		log.WithField("system", "hook"),
	)
	if err != nil {
		return nil, err
	}

	return &HTTPServer{
		cfg:  cfg,
		e:    e,
		log:  log,
		hook: hook,
	}, nil
}

// Start starts echo server
func (s *HTTPServer) Start() error {
	s.log.Infof("http server listening on %s", s.cfg.Addr)
	return s.e.Start(s.cfg.Addr)
}

// Stop does nothing
func (s *HTTPServer) Stop() error {
	s.log.Infof("stopping http server")
	return nil
}

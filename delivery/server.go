package delivery

import (
	"fmt"
	"usecase-1/config"
	"usecase-1/delivery/controller"
	"usecase-1/delivery/middleware"
	"usecase-1/manager"
	"usecase-1/utils/exceptions"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	log            *logrus.Logger
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))

	// semua controller disini
	controller.NewU1Controller(s.useCaseManager.NewU1UseCase(), s.engine)
	controller.NewU2Controller(s.useCaseManager.NewU2UseCase(), s.engine)
	controller.NewU3Controller(s.useCaseManager.NewU3UseCase(), s.engine)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckError(err)
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
		log:            logrus.New(),
	}
}

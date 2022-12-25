package application

import (
	"vikishptra/domain_todo-core/controller/todoapi"
	"vikishptra/domain_todo-core/gateway/withgorm"
	"vikishptra/domain_todo-core/usecase/getalltodo"
	"vikishptra/domain_todo-core/usecase/runtodochecked"
	"vikishptra/domain_todo-core/usecase/runtodocreate"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/server"
	"vikishptra/shared/infrastructure/token"
)

type todoapp struct{}

func NewTodoapp() gogen.Runner {
	return &todoapp{}
}

func (todoapp) Run() error {

	const appName = "todoapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := todoapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		getalltodo.NewUsecase(datasource),
		runtodochecked.NewUsecase(datasource),
		runtodocreate.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}

package todoapi

import (
	"github.com/gin-gonic/gin"

	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/token"
)

type selectedRouter = gin.IRouter

type ginController struct {
	*gogen.BaseController
	log      logger.Logger
	cfg      *config.Config
	jwtToken token.JWTToken
}

func NewGinController(log logger.Logger, cfg *config.Config, tk token.JWTToken) gogen.RegisterRouterHandler[selectedRouter] {
	return &ginController{
		BaseController: gogen.NewBaseController(),
		log:            log,
		cfg:            cfg,
		jwtToken:       tk,
	}
}

func (r *ginController) RegisterRouter(router selectedRouter) {

	resource := router.Group("/api/v1", r.authentication())
	resource.POST("/todo/message", r.authorization(), r.runTodoCreateHandler())

	resource.PUT("/todo/:id", r.authorization(), r.runTodoCheckedHandler())
	resource.GET("/todo", r.authorization(), r.getAllTodoHandler())
	resource.DELETE("/todo/message/:id", r.authorization(), r.runTodoDeleteByIDHandler())
	resource.GET("/todo/:id", r.authorization(), r.getTodoByIDHandler())
	resource.PUT("/todo/message/:id", r.authorization(), r.runUpdateMessageTodoByIDHandler())
}

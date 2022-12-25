package todoapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_todo-core/model/entity"
	"vikishptra/domain_todo-core/model/vo"
	"vikishptra/domain_todo-core/usecase/runtodochecked"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runTodoCheckedHandler() gin.HandlerFunc {

	type InportRequest = runtodochecked.InportRequest
	type InportResponse = runtodochecked.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		TodoId vo.TodoID `uri:"id"`
	}

	type response struct {
		Todo *entity.Todo `json:"todo"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindUri(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.TodoId = jsonReq.TodoId

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusNotFound, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Todo = res.Todo

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}

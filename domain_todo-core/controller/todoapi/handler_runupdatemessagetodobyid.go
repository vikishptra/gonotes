package todoapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_todo-core/model/errorenum"
	"vikishptra/domain_todo-core/usecase/runupdatemessagetodobyid"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runUpdateMessageTodoByIDHandler() gin.HandlerFunc {

	type InportRequest = runupdatemessagetodobyid.InportRequest
	type InportResponse = runupdatemessagetodobyid.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
		InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReqUri request
		if err := c.BindUri(&jsonReqUri); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonReqJSON request
		if err := c.BindJSON(&jsonReqJSON); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}
		var req InportRequest
		req.ID = jsonReqUri.ID
		req.Title = jsonReqJSON.Title
		req.Description = jsonReqJSON.Description

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			if err == errorenum.DataNull {
				c.JSON(http.StatusNotFound, payload.NewErrorResponse(err, traceID))
				return
			}
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Todo = res.Todo

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}

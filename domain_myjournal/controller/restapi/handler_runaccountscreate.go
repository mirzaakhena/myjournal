package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"myjournal/domain_myjournal/usecase/runaccountscreate"
	"myjournal/shared/infrastructure/logger"
	"myjournal/shared/infrastructure/util"
	"myjournal/shared/model/payload"
)

// runAccountsCreateHandler ...
func (r *Controller) runAccountsCreateHandler() gin.HandlerFunc {

	type request struct {
		runaccountscreate.InportRequest
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req runaccountscreate.InportRequest
		req = jsonReq.InportRequest

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunAccountsCreateInport.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}

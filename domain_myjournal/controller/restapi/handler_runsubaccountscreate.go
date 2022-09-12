package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"myjournal/domain_myjournal/usecase/runsubaccountscreate"
	"myjournal/shared/infrastructure/logger"
	"myjournal/shared/infrastructure/util"
	"myjournal/shared/model/payload"
)

// runSubAccountsCreateHandler ...
func (r *Controller) runSubAccountsCreateHandler() gin.HandlerFunc {

	type request struct {
		runsubaccountscreate.InportRequest
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

		var req runsubaccountscreate.InportRequest
		req = jsonReq.InportRequest

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunSubAccountsCreateInport.Execute(ctx, req)
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

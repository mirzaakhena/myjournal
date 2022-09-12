package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"myjournal/domain_myjournal/usecase/runwalletcreate"
	"myjournal/shared/infrastructure/logger"
	"myjournal/shared/infrastructure/util"
	"myjournal/shared/model/payload"
)

// runWalletCreateHandler ...
func (r *Controller) runWalletCreateHandler() gin.HandlerFunc {

	type request struct {
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

		var req runwalletcreate.InportRequest

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunWalletCreateInport.Execute(ctx, req)
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

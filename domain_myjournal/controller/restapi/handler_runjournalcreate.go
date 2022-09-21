package restapi

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"myjournal/domain_myjournal/usecase/runjournalcreate"
	"myjournal/shared/infrastructure/logger"
	"myjournal/shared/model/payload"
	"myjournal/shared/util"
)

// runJournalCreateHandler ...
func (r *Controller) runJournalCreateHandler() gin.HandlerFunc {

	type request struct {
		runjournalcreate.InportRequest
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

		var req runjournalcreate.InportRequest
		req = jsonReq.InportRequest
		req.Date = time.Now()
		req.WalletId = entity.WalletID(c.Param("walletId"))
		req.UserId = entity.UserID(c.GetString("userId"))

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunJournalCreateInport.Execute(ctx, req)
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

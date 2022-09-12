package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"myjournal/domain_myjournal/usecase/runjournalrollback"
	"myjournal/shared/infrastructure/logger"
	"myjournal/shared/infrastructure/util"
	"myjournal/shared/model/payload"
)

// runJournalRollbackHandler ...
func (r *Controller) runJournalRollbackHandler() gin.HandlerFunc {

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

		var req runjournalrollback.InportRequest

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunJournalRollbackInport.Execute(ctx, req)
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

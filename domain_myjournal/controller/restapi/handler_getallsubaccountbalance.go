package restapi

import (
	"context"
	"github.com/gin-gonic/gin"
	"myjournal/domain_myjournal/model/entity"
	"net/http"

	"myjournal/domain_myjournal/usecase/getallsubaccountbalance"
	"myjournal/shared/infrastructure/logger"
	"myjournal/shared/model/payload"
	"myjournal/shared/util"
)

// getAllSubaccountBalanceHandler ...
func (r *Controller) getAllSubaccountBalanceHandler() gin.HandlerFunc {

	type request struct {
		getallsubaccountbalance.InportRequest
	}

	type response struct {
		Count int64 `json:"count"`
		Items []any `json:"items"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.Bind(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req getallsubaccountbalance.InportRequest
		req.FindAllSubAccountBalanceRequest = jsonReq.FindAllSubAccountBalanceRequest
		req.WalletID = entity.WalletID(c.Param("walletId"))

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.GetAllSubaccountBalanceInport.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Count = res.Count
		jsonRes.Items = util.ToSliceAny(res.Items)

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}

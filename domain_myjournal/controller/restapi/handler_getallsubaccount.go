package restapi

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
	"net/http"

	"github.com/gin-gonic/gin"

	"myjournal/domain_myjournal/usecase/getallsubaccount"
	"myjournal/shared/infrastructure/logger"
	"myjournal/shared/model/payload"
	"myjournal/shared/util"
)

// getAllSubAccountHandler ...
func (r *Controller) getAllSubAccountHandler() gin.HandlerFunc {

	type request struct {
		getallsubaccount.InportRequest
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

		var req getallsubaccount.InportRequest
		req.FindAllSubAccountRequest = jsonReq.FindAllSubAccountRequest
		req.WalletID = entity.WalletID(c.Param("walletId"))
		req.Page = jsonReq.Page
		req.Size = jsonReq.Size

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.GetAllSubAccountInport.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Count = res.Count
		jsonRes.Items = res.Items

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}

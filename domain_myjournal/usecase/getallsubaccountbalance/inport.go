package getallsubaccountbalance

import (
	"context"
	"myjournal/domain_myjournal/model/entity"

	"myjournal/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

// InportRequest is request payload to run the usecase
type InportRequest struct {
	WalletId entity.WalletID
	Page     int64
	Size     int64
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	Count int64
	Items []any
}

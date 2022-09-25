package getallaccount

import (
	"context"
	"myjournal/domain_myjournal/model/repository"

	"myjournal/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

// InportRequest is request payload to run the usecase
type InportRequest struct {
	repository.FindAllAccountRequest
	Page int64 `form:"page,omitempty,default=1"`
	Size int64 `form:"size,omitempty,default=30"`
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	Count int64
	Items []any
}

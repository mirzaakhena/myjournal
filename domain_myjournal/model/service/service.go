package service

import "context"

type GetRandomStringService interface {
	GetRandomString(ctx context.Context) string
}

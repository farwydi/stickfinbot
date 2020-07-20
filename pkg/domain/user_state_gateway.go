package domain

import "context"

type UserStatGateway interface {
	GetCurrentState(ctx context.Context, userID int) (*UserState, error)
}

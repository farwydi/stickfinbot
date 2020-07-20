package userstate

import "github.com/farwydi/stickfinbot/pkg/domain"

func toUserState(m UserStat) *domain.UserState {
	return &domain.UserState{
		UserID: m.UserID,
	}
}

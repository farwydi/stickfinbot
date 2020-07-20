// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	telegram2 "github.com/farwydi/stickfinbot/internal/endpoint/telegram"
	"github.com/farwydi/stickfinbot/internal/gateway/gcp/userstate"
	"github.com/farwydi/stickfinbot/internal/gateway/telegram"
	"github.com/farwydi/stickfinbot/pkg/domain"
	"github.com/farwydi/stickfinbot/pkg/usecase/general"
)

// Injectors from bootstrap.go:

func bootstrap(contextContext context.Context) (domain.Endpoint, func(), error) {
	telegramGateway, err := telegram.NewTelegramGateway()
	if err != nil {
		return nil, nil, err
	}
	userStatGateway, cleanup, err := userstate.NewGCPUseStateGateway(contextContext)
	if err != nil {
		return nil, nil, err
	}
	generalGeneral := general.NewGeneralUseCase(telegramGateway, userStatGateway)
	endpoint := telegram2.NewTelegramEndpoint(generalGeneral)
	return endpoint, func() {
		cleanup()
	}, nil
}

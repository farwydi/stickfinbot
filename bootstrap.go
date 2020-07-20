// +build wireinject

package main

import (
	"context"
	tgendpoint "github.com/farwydi/stickfinbot/internal/endpoint/telegram"
	"github.com/farwydi/stickfinbot/internal/gateway/gcp/userstate"
	telegramgateway "github.com/farwydi/stickfinbot/internal/gateway/telegram"
	"github.com/farwydi/stickfinbot/pkg/domain"
	"github.com/farwydi/stickfinbot/pkg/usecase/general"
	"github.com/google/wire"
)

func bootstrap(context.Context) (domain.Endpoint, func(), error) {
	panic(wire.Build(
		userstate.NewGCPUseStateGateway,
		telegramgateway.NewTelegramGateway,
		general.NewGeneralUseCase,
		tgendpoint.NewTelegramEndpoint,
	))
}

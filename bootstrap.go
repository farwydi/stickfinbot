// +build wireinject

package main

import (
	tgendpoint "github.com/farwydi/stickfinbot/internal/endpoint/telegram"
	telegramgateway "github.com/farwydi/stickfinbot/internal/gateway/telegram"
	"github.com/farwydi/stickfinbot/pkg/domain"
	"github.com/farwydi/stickfinbot/pkg/usecase/general"
	"github.com/google/wire"
)

func bootstrap(token domain.TelegramBotToken) (domain.Endpoint, func(), error) {
	panic(wire.Build(
		telegramgateway.NewTelegramGateway,
		general.NewGeneralUseCase,
		tgendpoint.NewTelegramEndpoint,
	))
}

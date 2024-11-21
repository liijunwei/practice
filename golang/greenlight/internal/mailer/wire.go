//go:build wireinject
// +build wireinject

package mailer

import (
	"greenlight/internal/config"

	"github.com/google/wire"
)

func SetupMailer(cfg config.SMTP) Mailer {
	wire.Build(
		New,
	)

	return Mailer{}
}

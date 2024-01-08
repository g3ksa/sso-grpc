package auth

import (
	"context"
	"log/slog"
)

type Auth struct {
	log *slog.Logger
}

type UserStorage interface {
	SaveUser(
		ctx context.Context,
		email,
		password string,
	)
}

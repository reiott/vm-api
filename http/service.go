package http

import (
	"time"

	"go.uber.org/zap"
)

type Service struct {
	Store DataStore
	Now   func() time.Time
	Log   *zap.Logger
}

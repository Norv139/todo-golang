package utils

import (
	"context"
	"time"
)

func GetCtx() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	sex := 10 * time.Second
	return context.WithTimeout(ctx, sex)
}

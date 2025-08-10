package cache

import (
	"context"
	"time"
)


type Cache interface{
	SetValue(ctx context.Context, key string, value string, expiration time.Duration) error
}


package cacheredis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/cache"
	"github.com/go-redis/redis/v8"
)

type RedisCartCacheOptions struct {
	UniversalClient redis.UniversalClient
}

type RedisCartCache struct {
	o RedisCartCacheOptions
}

func NewRedisCartCache(options RedisCartCacheOptions) *RedisCartCache {
	return &RedisCartCache{
		o: options,
	}
}

func (c RedisCartCache) Get(ctx context.Context, userId string) (cache.Cart, error) {
	key := getFullKey(userId)

	valueData, err := c.o.UniversalClient.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = preseterrors.ErrNotFound
		}
		return cache.Cart{}, err
	}

	var cart cache.Cart
	if err := json.Unmarshal(valueData, &cart); err != nil {
		return cache.Cart{}, err
	}

	return cart, nil
}

func (c RedisCartCache) Set(ctx context.Context, userId string, cart cache.Cart) error {
	key := getFullKey(userId)

	data, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	if err := c.o.UniversalClient.Set(ctx, key, data, DurationTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (c RedisCartCache) Delete(ctx context.Context, userId string) error {
	key := getFullKey(userId)

	if err := c.o.UniversalClient.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}

func getFullKey(userId string) string {
	return fmt.Sprintf(FormatKey, userId)
}

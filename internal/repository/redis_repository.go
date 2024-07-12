package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"test-international-trade/internal/domain"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Save(data domain.EncryptedData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx := context.Background()
	return r.client.Set(ctx, data.Input, jsonData, 0).Err()
}

func (r *RedisRepository) Get(key string) (*domain.EncryptedData, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var data domain.EncryptedData
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

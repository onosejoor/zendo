package redis

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	*redis.Client
}

var (
	client   *RedisStore
	ttl      = time.Minute * 5
	clientMu sync.Mutex
)

func GetRedisClient() *RedisStore {
	clientMu.Lock()
	defer clientMu.Unlock()
	if client != nil {
		return client
	}

	opts, err := redis.ParseURL(os.Getenv("REDIS_UPSTASH_URL"))
	if err != nil {
		log.Fatalln("error parsing redis_upstash_url: ", err.Error())
		return nil
	}

	store := redis.NewClient(opts)
	if err := store.Ping(context.TODO()).Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("connected to redis client")

	client = &RedisStore{
		Client: store,
	}

	return client
}

func (c *RedisStore) GetCacheData(key string, ctx context.Context) ([]byte, error, bool) {
	data, err := c.Client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil || len(data) < 1 {
			return nil, err, true
		} else {
			return nil, err, false
		}

	}

	return data, nil, false
}

func (c *RedisStore) SetCacheData(key string, ctx context.Context, value string) error {
	err := c.Client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (c *RedisStore) GetCacheData(key string, ctx context.Context, v any) (error, bool) {
	data, err := c.Client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil || len(data) < 1 {
			return err, true
		} else {
			return err, false
		}

	}

	_ = json.Unmarshal(data, &v)

	return nil, false
}

func (c *RedisStore) SetCacheData(key string, ctx context.Context, value any) error {
	jsonValue, _ := json.Marshal(&value)

	err := c.Client.Set(ctx, key, jsonValue, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func DeleteTaskCache(ctx context.Context, id string, taskId string) error {
	redisClient := GetRedisClient()
	key := []string{fmt.Sprintf("user:%s:task:%s", id, taskId), fmt.Sprintf("user:%s:tasks", id), fmt.Sprintf("user:%s:stats", id)}

	if err := redisClient.Client.Del(ctx, key...).Err(); err != nil {
		log.Println("Error deleting data from cache: ", err.Error())
		return err
	}
	return nil
}

func DeleteProjectCache(ctx context.Context, id string, taskId string) error {
	redisClient := GetRedisClient()
	key := []string{fmt.Sprintf("user:%s:project:%s", id, taskId), fmt.Sprintf("user:%s:projects", id), fmt.Sprintf("user:%s:stats", id)}

	if err := redisClient.Client.Del(ctx, key...).Err(); err != nil {
		log.Println("Error deleting data from cache: ", err.Error())
		return err
	}
	return nil
}

func (redisClient *RedisStore) GetCacheHandler(ctx *fiber.Ctx, result any, key string, name string) bool {

	err, isEmpty := redisClient.GetCacheData(key, ctx.Context(), &result)
	if err != nil && !isEmpty {
		log.Println("Error querying cached data: ", err.Error())
		ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
		return true
	}

	if !isEmpty {
		ctx.Status(200).JSON(bson.M{
			"success": true,
			name:      result,
		})
		log.Println("Returning from cache with key: ", key)
		return true
	}

	return false

}

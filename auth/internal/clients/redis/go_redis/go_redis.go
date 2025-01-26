package go_redis

import (
	"context"
	"fmt"
	"github.com/MentalMentos/techTaskMetr/auth/internal/clients/redis"
	"github.com/goccy/go-json"
	goRedis "github.com/redis/go-redis/v9"
	"time"
)

// Ensure GoRedisClient implements the RedisClient interface
var _ redis.IRedis = (*GoRedisClient)(nil)

// GoRedisClient is a wrapper around the go-redis client, providing methods
// for interacting with a Redis data store.
type GoRedisClient struct {
	Client *goRedis.Client // Underlying go-redis client instance
}

// NewGoRedisClient initializes a new GoRedisClient with the provided Redis configuration.
// It creates a new Redis client using the configuration's address.
func NewGoRedisClient(config redis.IRedisConfig) *GoRedisClient {
	redisClient := &GoRedisClient{
		Client: goRedis.NewClient(&goRedis.Options{
			Addr:     config.Address(),
			Password: "", // Password is set to an empty string for no authentication
			DB:       0,  // Using the default database
		}),
	}

	return redisClient
}

// Set stores a string value in Redis under the specified key with a given duration.
// The duration determines how long the value will be stored before expiring.
func (g *GoRedisClient) Set(ctx context.Context, key string, value string, duration time.Duration) error {
	return g.Client.Set(ctx, key, value, duration).Err()
}

// Get retrieves the string value associated with the specified key from Redis.
// If the key does not exist or an error occurs, it returns an error.
func (g *GoRedisClient) Get(ctx context.Context, key string) (string, error) {
	return g.Client.Get(ctx, key).Result()
}

// SetObject serializes a given object into JSON format and stores it in Redis under the specified key
// with a given duration. If marshaling fails, an error is returned.
func (g *GoRedisClient) SetObject(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	noteBytes, err := json.Marshal(value) // Serialize the object to JSON
	if err != nil {
		return fmt.Errorf("failed to marshal note: %v", err) // Return an error if marshaling fails
	}
	g.Client.Set(ctx, key, string(noteBytes), duration) // Store the JSON string in Redis with expiration
	return nil
}

// GetObject retrieves the JSON string associated with the specified key from Redis,
// unmarshals it into the provided value parameter. If retrieval or unmarshaling fails,
// an error is returned.
func (g *GoRedisClient) GetObject(ctx context.Context, key string, value any) error {
	const mark = "Clients.Redis.GoRedis.GetObject"

	val, err := g.Client.Get(ctx, key).Result() // Retrieve the value from Redis
	if err != nil {
		return fmt.Errorf("failed to get note from redis: %v", err) // Return an error if retrieval fails
	}

	if err = json.Unmarshal([]byte(val), &value); err != nil {
		return fmt.Errorf("failed to unmarshal note: %v", err) // Return an error if unmarshaling fails
	}
	return nil
}

// Delete removes the specified key and its associated value from Redis.
// It returns an error if the deletion fails.
func (g *GoRedisClient) Delete(ctx context.Context, key string) error {
	return g.Client.Del(ctx, key).Err()
}

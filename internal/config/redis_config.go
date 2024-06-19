package config

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

/*
ConnectRedis establishes a connection to a Redis server using the configuration provided. It creates
a new Redis client with the specified server address retrieved from the configuration. The function
then attempts to ping the Redis server to ensure a successful connection.

Parameters:
  - None.

Returns:
  - A pointer to the Redis client if the connection is successful.
  - An error, if any, encountered during the connection process.
*/
func ConnectRedis() (*redis.Client, error) {
	// Create a redis client instance
	redisClient := redis.NewClient(&redis.Options{
		Addr: viper.GetString(`redis.address`),
		//Password: viper.GetString(`redis.password`),
		//DB: viper.GetInt(`redis.database`),
	})

	// Ping the Redis server to check if connection is successful
	_, err := redisClient.Ping().Result()
	// Return an error message if connection to Redis fails
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}

/*
SetToRedis sets a key-value pair in Redis with the specified key and value.

Parameters:
  - redisClient: A pointer to the Redis client instance.
  - key: The key to associate with the specified value in Redis.
  - val: The value to be stored in Redis.
*/
func SetToRedis(redisClient *redis.Client, key, val string) {
	// Set the key-value pair in Redis with a time-to-live (TTL) of 0 (no expiration)
	err := redisClient.Set(key, val, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

/*
GetFromRedis retrieves a value from Redis using the specified key.

Parameters:
  - redisClient: A pointer to the Redis client instance.
  - key: The key associated with the value to be retrieved from Redis.

Returns:
  - The value associated with the key in Redis.
  - An error, if any, encountered during the retrieval process.
*/
func GetFromRedis(redisClient *redis.Client, key string) (string, error) {
	// Attempt to retrieve the value from Redis using the specified key
	val, err := redisClient.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val, err
}

/*
DeleteFromRedis attempts to retrieve a value from Redis using the specified key and deletes it if
present.

Parameters:
  - redisClient: A pointer to the Redis client instance.
  - key: The key associated with the value to be deleted from Redis.

Returns an error if the key is not found or has expired, indicating an invalid or expired token.
*/
func DeleteFromRedis(redisClient *redis.Client, key string) error {
	// Retrieve the value from Redis using the specified key
	_, err := redisClient.Get(key).Result()

	// If the key is not found, return an error indicating an invalid or expired access token
	if err == redis.Nil {
		return errors.New("invalid or expired token")
	}

	// Delete the key from Redis
	redisClient.Del(key)

	return err
}

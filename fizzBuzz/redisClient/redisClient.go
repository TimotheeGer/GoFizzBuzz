package redisclient

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"training.go/fizzBuzz/models"
)

var ctx = context.Background()

func StartRedis() (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),     // adresse et port du serveur Redis
		Password: os.Getenv("REDIS_PASSWORD"), // le mot de passe, si nécessaire
		DB:       0,                           // numéro de la base de données à utiliser
	})

	// Envoi d'une commande PING pour tester la connexion
	var ctx = context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(pong, "La connexion à Redis a été établie avec succès.")
	return rdb, nil
}

func StoreEvent(ClientRedis *redis.Client, key string, dataJson []byte) error {

	// Récupère et imprime la valeur associée à la clé "key".
	val, err := ClientRedis.Get(ctx, key).Result()
	if err == redis.Nil {

		if err = ClientRedis.Set(ctx, key, dataJson, 0).Err(); err != nil {
			return err
		}
		fmt.Println("nouvelle key enregister sur redis:", key)

	} else if err != nil {
		return err
	} else {

		var structFizzBuzz models.QueryFizzBuzz

		if err = json.Unmarshal([]byte(val), &structFizzBuzz); err != nil {
			return err
		}

		structFizzBuzz.Count++

		jsonFizzBuzz, err := json.Marshal(structFizzBuzz)
		if err != nil {
			return err
		}

		if err = ClientRedis.Set(ctx, key, jsonFizzBuzz, 0).Err(); err != nil {
			return err
		}
		fmt.Println("key update sur redis:", key)
	}
	return nil
}



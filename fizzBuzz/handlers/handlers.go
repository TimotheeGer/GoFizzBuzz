package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo"
	"training.go/fizzBuzz/models"
	redisclient "training.go/fizzBuzz/redisClient"
	"training.go/fizzBuzz/services"
)

var ctx = context.Background()

func HashData(data []byte) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func CheckQueryParam(c echo.Context) error {

	if params := len(c.QueryParams()); params != 5 {
		return fmt.Errorf(fmt.Sprintf("Nombre de paramètres incorrect. 5 paramètres attendus, mais %d reçus.", params))
	}

	int1, err := strconv.Atoi(c.QueryParam("int1"))
	if err != nil || int1 <= 0 {
		return errors.New("paramètre int1 invalide ou manquant")
	}

	int2, err := strconv.Atoi(c.QueryParam("int2"))
	if err != nil || int2 <= 0 {
		return errors.New("paramètre int2 invalide ou manquant")
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		return errors.New("paramètre limit invalide ou manquant")
	}

	str1 := c.QueryParam("str1")
	if str1 == "" {
		return errors.New("paramètre str1 invalide ou manquant")
	}

	str2 := c.QueryParam("str2")
	if str2 == "" {
		return errors.New("paramètre str2 invalide ou manquant")
	}

	return nil
}

func FizzBuzzHandler(c echo.Context, clientRedis *redis.Client) error {

	if err := CheckQueryParam(c); err != nil {
		return models.ResponseFizzBuzz(c, http.StatusBadRequest, nil, false, err.Error())
	}

	queryFizzBuzz := new(models.QueryFizzBuzz)
	if err := c.Bind(queryFizzBuzz); err != nil {
		return models.ResponseFizzBuzz(c, http.StatusBadRequest, nil, false, err.Error())
	}
	queryFizzBuzz.Count = 1;
	jsonFizzBuzz, err := json.Marshal(queryFizzBuzz)
	if err != nil {
		return models.ResponseFizzBuzz(c, http.StatusInternalServerError, nil, false, err.Error())
	}

	hashJson := HashData(jsonFizzBuzz)

	if err := redisclient.StoreEvent(clientRedis, hashJson, jsonFizzBuzz); err != nil {
		return models.ResponseFizzBuzz(c, http.StatusInternalServerError, nil, false, err.Error())
	}

	result := services.CalculateFizzBuzz(
		queryFizzBuzz.Int1,
		queryFizzBuzz.Int2,
		queryFizzBuzz.Limit,
		queryFizzBuzz.Str1,
		queryFizzBuzz.Str2)

	if len(result) > 0 {
		message := fmt.Sprintf("Résultats de FizzBuzz pour les valeurs Int1: %v, Int2: %v, Str1: %v, Str2: %v et limite: %v.",
		queryFizzBuzz.Int1,
		queryFizzBuzz.Int2,
		queryFizzBuzz.Str1,
		queryFizzBuzz.Str2,
		queryFizzBuzz.Limit)
		
		return models.ResponseFizzBuzz(c, http.StatusOK, result, true, message)
	}

	return c.NoContent(http.StatusNoContent)
}

func StatisticsHandler(c echo.Context, clientRedis *redis.Client) error {

	find := false
	structFizzBuzz := new(models.QueryFizzBuzz)

	keys, err := clientRedis.Keys(ctx, "*").Result()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for _, key := range keys {

		structFizzBuzzTmp := new(models.QueryFizzBuzz)

		val, err := clientRedis.Get(ctx, key).Result()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if err = json.Unmarshal([]byte(val), &structFizzBuzzTmp); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if structFizzBuzzTmp.Count > structFizzBuzz.Count {
			structFizzBuzz = structFizzBuzzTmp
		}

		find = true
	}

	if find {
		return models.ResponseStatistics(c, http.StatusOK, *structFizzBuzz, true,
			fmt.Sprintf("La requête la plus utilisée a été appelée %v fois.", structFizzBuzz.Count))
	}

	return c.NoContent(http.StatusNoContent)
}

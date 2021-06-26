package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

var ctx = context.Background()

type HashingService struct {
	rdb *redis.Client
}

func NewHashing(rdb *redis.Client) *HashingService {
	return &HashingService{
		rdb: rdb,
	}
}

func (h *HashingService) GetBodyFromHash(urlStr string) (bool, []byte) {
	val, err := h.rdb.Get(ctx, urlStr).Result()
	if err == redis.Nil{
		return false, nil
	} else if err != nil {
		log.Fatalf("hash get failed: %s", err.Error())
	}

	return true, []byte(val)
}

func (h *HashingService) SetBodyInHash(urlStr string, body []byte) {
	hashRecordsLimit := viper.GetInt("hashRecordsLimit")
	hashRecords, err := h.rdb.LLen(ctx, "hashList").Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	if int(hashRecords) >= hashRecordsLimit {
		if int(hashRecords) > hashRecordsLimit{
			h.removeExcess()
		}

		oldUrl, err := h.rdb.LPop(ctx, "hashList").Result()
		if err == redis.Nil {
			fmt.Println("\"hashList\" does not exist")
		} else if err != nil {
			log.Fatal(err.Error())
		} else {
			h.rdb.Del(ctx, oldUrl)
		}
	}

	err = h.rdb.Set(ctx, urlStr, body, 0).Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = h.rdb.RPush(ctx, "hashList", urlStr).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *HashingService) removeExcess() {
	hashRecordsLimit := viper.GetInt("hashRecordsLimit")
	hashRecords, err := h.rdb.LLen(ctx, "hashList").Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	difference := int(hashRecords) - hashRecordsLimit
	for i := difference; i > 0; i-- {
		oldUrl, err := h.rdb.LPop(ctx, "hashList").Result()
		if err == redis.Nil {
			fmt.Println("\"hashList\" does not exist")
		} else if err != nil {
			log.Fatal(err.Error())
		} else {
			h.rdb.Del(ctx, oldUrl)
		}
	}
}

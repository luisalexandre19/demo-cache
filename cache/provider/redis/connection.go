package redis

import (
	"context"
	"errors"
	"time"

	boot "demo.cache/boostrap"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var (
	ErrNil = errors.New("Nenhum registro encontrado no cache!")
	Ctx    = context.TODO()
)

const (
	KEY_APPLICATIONS = "APPLICATIONS"
)

func ConnAuth() (client *redis.Client, err error) {

	log.Info("Try to connect Redis")

	client = redis.NewClient(&redis.Options{
		Addr:        boot.REDIS_CONFIG.Host + ":" + boot.REDIS_CONFIG.Port,
		Username:    boot.REDIS_CONFIG.User,
		Password:    boot.REDIS_CONFIG.Password,
		DB:          boot.REDIS_CONFIG.Db,
		PoolSize:    boot.REDIS_CONFIG.PoolSize,
		MaxConnAge:  time.Second * time.Duration(boot.REDIS_CONFIG.MaxConnAge),
		IdleTimeout: time.Second * time.Duration(boot.REDIS_CONFIG.IdleTimeout),
		DialTimeout: time.Second * time.Duration(boot.REDIS_CONFIG.DialTimeout),
	})

	_, err = client.Ping(Ctx).Result()

	if err != nil {
		log.Error("--------- Error PING redis ", err.Error())
		client.Close()
		return nil, err
	}

	log.Info("Connected on Redis")

	return
}

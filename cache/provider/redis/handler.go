package redis

import (
	"encoding/json"
	"time"

	boot "demo.cache/boostrap"
	"demo.cache/cache/domain"
	log "github.com/sirupsen/logrus"
)

func (rp *RedisProvider) Get(key string) domain.CacheResponse {
	val, err := rp.Client.Get(Ctx, key).Bytes()

	var obj domain.ResponseOperation
	if err := json.Unmarshal(val, &obj); err != nil {
		log.Info("Error parse Json ", err.Error())
	}

	resp := domain.CacheResponse{}
	resp.SetStatus(err)

	resp.SetData(obj.Data)
	resp.SetHeader(obj.Header)
	return resp

}

func (rp *RedisProvider) Set(key string, data interface{}) domain.CacheResponse {

	resp := domain.CacheResponse{}

	jsonParsed, err := json.Marshal(data)

	if err != nil {
		resp.SetStatus(err)
		log.Debugf("Error add data in redis: %s", resp)
	}

	statusCmd := rp.Client.Set(Ctx, key, jsonParsed, time.Minute*time.Duration(boot.REDIS_CONFIG.Ttl)).Err()

	resp.SetStatus(statusCmd)

	if statusCmd != nil {
		log.Error("Unkhow rrror on add data in Redis : ", statusCmd)
		rp.ConneError = statusCmd
		return resp
	}

	return resp
}

func (rp *RedisProvider) Del(key string) domain.CacheResponse {

	error := rp.Client.Del(Ctx, key).Err()
	resp := domain.CacheResponse{}
	resp.SetStatus(error)

	if error != nil {
		rp.ConneError = error
		return resp
	}

	return domain.CacheResponse{}
}

func (rp *RedisProvider) Connect() domain.CacheConnection {

	rp.Client, rp.ConneError = ConnAuth()

	if rp.ConneError != nil {
		log.Errorf("--------- Fail to connect Redis:  HOST %s PORT %d | ERROR : %s ", boot.REDIS_CONFIG.Host, boot.REDIS_CONFIG.Port, rp.ConneError)
	} else {
		log.Info(" --------------------------  Connected to cache provider --------------------------")
	}

	return domain.CacheConnection{Err: rp.ConneError, Conn: rp.Client}
}

func (rp *RedisProvider) Error() error {
	return rp.ConneError
}

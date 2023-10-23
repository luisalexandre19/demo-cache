package cache

import (
	"errors"

	"demo.cache/cache/domain"
	log "github.com/sirupsen/logrus"
)

var RepositoryCache domain.CacheRepository

func Initialize() {
	RepositoryCache = FactoryProvider()
	RepositoryCache.Connect()
}

//Faz a busca de dados no redis, recebe os parametros do request, e executa a func que gera a key no cache
func FindDataCache(parametersRequest string) domain.CacheResponse {
	var cacheResponse domain.CacheResponse

	if !tryConnectionIsError() {
		cacheResponse.SetStatus(errors.New("Erro ao conectar no provider de cache "))
		return cacheResponse
	}
	key := GenerateCacheKey(parametersRequest)
	cacheResponse = RepositoryCache.Get(key)
	if cacheResponse.Data() == "" {
		log.Debugf("Not found data in cache provider for key [ %s ]  | params [ %s ]  | MSG  [ %s ] ", key, parametersRequest, cacheResponse.Error())
		return cacheResponse
	} else {
		log.Debugf("Found data in cache provider :  %s with headers %s ", cacheResponse.Data(), cacheResponse.Header())
		log.Infof("Found data in cache provider for key [ %s ]  | params [ %s ]  ", key, parametersRequest)
		return cacheResponse
	}
}

func SetDataCache(parametersRequest string, data interface{}) error {

	key := GenerateCacheKey(parametersRequest)

	if !tryConnectionIsError() {
		return errors.New("Error to connect with provider")
	}
	log.Debugf(" Add data from bussines api to cache provider: [%s]  :  key [ %s ] | params [ %s ] : ", key, parametersRequest)
	status := RepositoryCache.Set(key, data)
	return status.Status.Err
}

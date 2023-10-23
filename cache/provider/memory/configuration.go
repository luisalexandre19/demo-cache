package memory

import (
	"time"

	boot "demo.cache/boostrap"
	"github.com/allegro/bigcache/v3"

	log "github.com/sirupsen/logrus"
)

var bigCacheInstance *bigcache.BigCache

func (cache *MemoryProvider) Build() *bigcache.BigCache {

	if bigCacheInstance == nil {
		bigCacheInstance = NewInstance()
	}
	return bigCacheInstance
}

func NewInstance() *bigcache.BigCache {

	bigCache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Duration(boot.REDIS_CONFIG.Ttl) * time.Minute))

	if err != nil {
		log.Errorf("Erro ao criar instancia cache memory  : %s ", err.Error())
		return nil
	}
	return bigCache
}

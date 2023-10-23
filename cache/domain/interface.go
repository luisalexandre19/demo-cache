package domain

type CacheRepository interface {
	Get(key string) CacheResponse
	Set(key string, data interface{}) CacheResponse
	Del(key string) CacheResponse
	Connect() CacheConnection
	Error() error
}

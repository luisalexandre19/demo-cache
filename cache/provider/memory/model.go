package memory

import (
	"encoding/json"
	"errors"

	"demo.cache/cache/domain"
	log "github.com/sirupsen/logrus"
)

type MemoryProvider struct {
	ConneError error
}

func (mp *MemoryProvider) Get(key string) (resp domain.CacheResponse) {
	byteData, err := mp.Build().Get(key)
	if err != nil {
		resp.SetStatus(err)
		log.Debugf("Erro recuperar cache in memory: %s", resp)
		return
	}

	var obj domain.ResponseOperation
	if err := json.Unmarshal(byteData, &obj); err != nil {
		log.Info("Erro parse JSon ", err.Error())
	}

	resp.SetData(obj.Data)
	resp.SetHeader(obj.Header)
	log.Debugf("Ger data cache in memory: %s", resp)
	return

}
func (mp *MemoryProvider) Set(key string, data interface{}) domain.CacheResponse {

	resp := domain.CacheResponse{}

	jsonParsed, err := json.Marshal(data)

	if err != nil {
		resp.SetStatus(errors.New("Error on create bytes to cache"))
		log.Debugf("Erro add data cache in memory: %s", resp)
	} else {

		err := mp.Build().Set(key, jsonParsed)

		if err != nil {
			resp.SetStatus(err)
			log.Debugf("Erro add data cache in memory: %s", resp)
		}

	}

	return resp
}
func (mp *MemoryProvider) Del(key string) (resp domain.CacheResponse) {
	err := mp.Build().Delete(key)
	resp.SetStatus(err)
	if err != nil {
		log.Debugf("Erro ao deletar data cache in memory: %s", resp)
	}
	return
}
func (rp *MemoryProvider) Connect() domain.CacheConnection {

	return domain.CacheConnection{}
}

func (rp *MemoryProvider) Error() error {
	return nil
}

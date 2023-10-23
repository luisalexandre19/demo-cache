package web

import (
	"net/http"

	boot "demo.cache/boostrap"
	"demo.cache/cache"
	"demo.cache/cache/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var GetRouter *mux.Router

func init() {
	GetRouter = mux.NewRouter()
}

func CreateEndpoints() {
	GetRouter.PathPrefix("/").HandlerFunc(findData).Methods(http.MethodGet)
	return
}

func findData(response http.ResponseWriter, request *http.Request) {

	extractedParamnsCache, err := ExtractParamsKey(request)

	if err == nil {

		var responseData domain.CacheResponse

		response.Header().Set("Content-Type", "application/json")

		responseData = cache.FindDataCache(extractedParamnsCache)
		if responseData.Data() != "" {
			response.Header().Set("x-cached", "true")
			response.Header().Set("x-cache-provider", boot.APP_CONFIG.Provider)
			writeHeaders(response, responseData.Header())
			response.Write([]byte(responseData.Data()))
		}

		if responseData.Data() == "" {
			responseCall := findBussinesData(request, extractedParamnsCache)
			writeHeaders(response, responseCall.HttpResponse.Header)
			response.WriteHeader(responseCall.HttpResponse.StatusCode)
			response.Write([]byte(responseCall.DataContent.Data))
		}

	} else {

		responseModel := CallServiceMain(request)
		writeHeaders(response, responseModel.HttpResponse.Header)
		response.WriteHeader(responseModel.HttpResponse.StatusCode)
		response.Write([]byte(responseModel.DataContent.Data))
	}

}

func writeHeaders(response http.ResponseWriter, header http.Header) {
	for key, value := range header {
		response.Header().Set(key, value[0])
	}
}

func findBussinesData(request *http.Request, paramnsCache string) ResponseModel {

	var cacheErr error
	responseModel := CallServiceMain(request)

	if len(responseModel.Error) > 0 {
		log.Error("Error on call business api ", responseModel.Error)
	}

	if responseModel.HttpStatusCode >= 200 && responseModel.HttpStatusCode < 300 {
		cacheErr = cache.SetDataCache(paramnsCache, responseModel.DataContent)
	}

	if cacheErr != nil {
		log.Errorf("Error:   [ %s ] ",
			cacheErr.Error())
	}
	return responseModel
}

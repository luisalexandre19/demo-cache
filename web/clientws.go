package web

//define os metodos de chamadas para o serviço principal, container que contem os dados de negocio, ou chamada para um serviço externo
//caso tenha sido parametrizado para chamada externa
import (
	"net/http"

	boot "demo.cache/boostrap"
	log "github.com/sirupsen/logrus"
)

func CallServiceMain(request *http.Request) ResponseModel {
	var responseModel ResponseModel
	resp, err := DoRequest(request)

	if err != nil {
		parseResponseData, err := ParseHttpResponseJson(resp)
		log.Error("Received an unkhown answer from business api :  ", parseResponseData)
		responseModel.Error = append(responseModel.Error, err)
		return responseModel
	}

	responseModel = ResponseModel{HttpResponse: resp, HttpStatusCode: resp.StatusCode}

	parseResponseData, err := ParseHttpResponseJson(resp)

	responseCacheModel := ResponseCacheModel{Data: parseResponseData, Header: resp.Header}

	responseModel.DataContent = responseCacheModel

	return responseModel
}

func DoRequest(request *http.Request) (*http.Response, error) {
	var concat ConcatString

	urlRequest := concat.Add(boot.APP_CONFIG.BussinesContainerAddr).Build()
	log.Debugf("Call business api:  [ %s ] ", urlRequest)
	req, err := http.NewRequest(request.Method, urlRequest, request.Body)

	if err != nil {
		log.Errorf("Error on create request to call business api [ %s ] ", err)
		return nil, err
	}
	for key, value := range request.Header {
		req.Header.Set(key, value[0])
		log.Debugf("\nHeaders added to call business api :  [ %s:%s ] ", key, value[0])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error on execute HTTP [%s]  |  \nHost [ %s ] |\n Body [ %s ] |\n Headers [ %s ] ", request.Method, req.Host, req.Body, req.Header)
	} else {

		for key, value := range resp.Header {
			log.Debugf("\nResponse headers from business api :  [ %s:%s ] ", key, value[0])
		}
	}

	return resp, err
}

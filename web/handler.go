package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"sort"

	"encoding/base64"

	"github.com/Jeffail/gabs"
	log "github.com/sirupsen/logrus"
)

//extrai os parametros do requests do path e da query ,e retorna uma string concatenada para usar como chave de cache no provider
func ExtractParamsKey(r *http.Request) (string, error) {
	queryStringParams := getQueryStringParams(r)

	bodyParams := getBodyParams(r)
	log.Debugf("-----------query => [%s] body => [%s] ", queryStringParams, bodyParams)

	return queryStringParams + bodyParams, nil
}

func getBodyParams(r *http.Request) string {

	if r.Body != nil {

		responseData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("  Error on read response body : %s ", err)
			printError(responseData)
			return ""
		}

		r.Body = io.NopCloser(bytes.NewBuffer(responseData))

		return base64.StdEncoding.EncodeToString([]byte(responseData))

	} else {
		return ""
	}

}

func getQueryStringParams(request *http.Request) (allParamns string) {

	params := request.URL.Query()
	var concat ConcatString
	keysOrder := orderKeysMap(params)

	for _, key := range keysOrder {
		concat.Add(key).Add("=").Add(params[key][0])
	}

	return concat.Build()
}

//faz o cast do http response e retornar uma string no formato json
func ParseHttpResponseJson(reqApi *http.Response) (string, error) {

	if reqApi.Body == nil {
		return "", errors.New("Request body is empty")
	}
	responseData, err := ioutil.ReadAll(reqApi.Body)
	if err != nil {
		log.Error(" Error to read response from business api: ", err)
		return "", err
	}
	jsonParsed, err := gabs.ParseJSON(responseData)

	if err != nil {
		log.Error("  ParseHttpResponseJson Error : %s ", err)
		printError(responseData)
		return "", err
	}
	return jsonParsed.String(), nil
}

func printError(responseData []byte) {

	var obj json.RawMessage
	if err := json.Unmarshal(responseData, &obj); err != nil {
		log.Info("Error parse JSon ", err.Error())
	}

	log.Infof("%+v\n", obj)
	log.Infof("JSON: %s\n", obj)

}

func ParseStringToJson(jsonString string) (string, error) {

	jsonParsed, err := gabs.ParseJSON([]byte(jsonString))

	if err != nil {
		log.Error(" Error to parse json : ", err)
		return "", err
	}
	return jsonParsed.String(), nil
}

func orderKeysMap(values map[string][]string) []string {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

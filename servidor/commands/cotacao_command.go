package commands

import (
	"context"
	"encoding/json"
	"github.com/alacerda/gotation/entities"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const externalAPI = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

func GetCotacao(w http.ResponseWriter, r *http.Request) {
	var httpClient http.Client
	var actual entities.Cotacao

	requestContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	request, err := http.NewRequestWithContext(requestContext, "GET", externalAPI, nil)
	if err != nil {
		log.Errorln(err)
	}

	response, err := httpClient.Do(request)
	if err != nil {
		log.Errorln(err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorln(err)
	}

	err = json.Unmarshal(responseBody, &actual)
	if err != nil {
		log.Errorln(err)
	}

	// loga a cotacao recebida
	newLog := LogQuote{
		quote: &actual,
	}
	go newLog.logQuote()

	// envia o cotacao para o solicitante
	_, err = w.Write([]byte("{\"quote\":\"" + actual.USDBRL.Bid + "\"}"))
	if err != nil {
		log.Errorln(err)
	}
}

package main

import (
	"context"
	"github.com/alacerda/gotation/cliente/commands"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func init() {
	log.SetReportCaller(true)
	log.SetLevel(log.ErrorLevel)
}

var url string = "http://127.0.0.1:8080/cotacao"

func main() {
	contexto := context.Background()
	contexto, cancel := context.WithTimeout(contexto, time.Millisecond*300)
	defer cancel()

	request, err := http.NewRequestWithContext(contexto, "GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	logger := commands.LogResponse{Response: string(responseBody)}
	logger.Save()

	log.Debugln("Response: ", string(responseBody))
}

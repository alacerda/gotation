package main

import (
	. "github.com/alacerda/gotation/servidor/commands"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	log.SetReportCaller(true)
	log.SetLevel(log.ErrorLevel)
}

var err error

func main() {
	for {
		http.HandleFunc("/cotacao", GetCotacao)
		err = http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Errorln(err)
		}
	}
}

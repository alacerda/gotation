package commands

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type LogResponse struct {
	Response string
}

func (self *LogResponse) Save() {
	logFile, err := os.OpenFile("client.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Errorln(err)
	}
	defer logFile.Close()

	logFile.WriteString(self.Response + " at " + time.DateTime + "\n")
}

package commands

import (
	"context"
	"database/sql"
	"github.com/alacerda/gotation/entities"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"time"
)

var err error

type LogQuote struct {
	quote *entities.Cotacao
}

func (self *LogQuote) logQuote() {
	db, err := sql.Open("sqlite3", "server.db")
	if err != nil {
		log.Errorln(err)
	} else {
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS quote(valor varchar(10), created_at DATETIME DEFAULT CURRENT_TIMESTAMP)")
		if err != nil {
			log.Errorln(err)
		}
		contexto := context.Background()
		contexto, cancel := context.WithTimeout(contexto, time.Millisecond*10)
		defer cancel()

		_, err = db.ExecContext(contexto, "INSERT INTO quote(valor) values(?)", self.quote.USDBRL.Bid)
		if err != nil {
			log.Errorln(err)
		}
	}
}

package main

import (
	"framework/web/ping"
	"net/http"

	"github.com/AlexStocks/log4go"
)

func main() {

	r := ping.Default()
	log := log4go.NewDefaultLogger(log4go.INFO)
	defer log.Close()
	r.GET("/panic", func(c *ping.Context) {
		names := []string{"ping"}
		c.String(http.StatusOK, names[100])
	})
	r.GET("/ok", func(c *ping.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.Run(":9999")
}

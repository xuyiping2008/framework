package ping

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()

		/*logServer := &provider.LogServiceProvider{
			Driver: "rotate",
			CtxFielder: func(par interface{}) map[string]interface{} {
				return map[string]interface{}{"att": "ping"}
			},
		}
		newLog := logServer.Register()
		params := logServer.Params()

		logInter, err := newLog(params)
		if err != nil {
			log.Printf("[code = %d],[path=%s] Logger error: %s", c.StatusCode, c.Req.RequestURI, err.Error())
		}

		logPing, ok := logInter.(provider.PingLog)
		if ok {
			logPing.Info("test", map[string]interface{}{"console": 111111})
		} else {
			fmt.Println("log print error")
		}*/

		log.Printf("[code = %d],[path=%s] Logger time %d", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

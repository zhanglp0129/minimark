package main

import (
	"fmt"
	"minimark/config"
	"minimark/routers"
)

func main() {
	cfg := config.GetConfig()
	r := routers.GetRouters()
	err := r.Run(fmt.Sprintf("%s:%d", cfg.Server.IP, cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}

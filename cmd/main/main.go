package main

import (
	"fmt"
	"minimark/config"
	"minimark/dao"
	"minimark/routers"
)

func main() {
	cfg := config.GetConfig()
	dao.GetDB()
	r := routers.GetRouters()
	err := r.Run(fmt.Sprintf("%s:%d", cfg.Server.IP, cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}

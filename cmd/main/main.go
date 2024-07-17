package main

import (
	"fmt"
	"minimark/config"
	"minimark/dao"
)

func main() {
	cfg := config.GetConfig()
	fmt.Printf("%+v\n", cfg)
	dao.GetDB()
	//r := routers.GetRouters()
	//err := r.Run(fmt.Sprintf("%s:%d", cfg.Server.IP, cfg.Server.Port))
	//if err != nil {
	//	panic(err)
	//}
}

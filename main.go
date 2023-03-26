/**
 * @Author think
 * @Date 2023/3/26$
 * @Note
 **/

package main

import (
	"server-resource-probe/config"
	"server-resource-probe/router"
	"server-resource-probe/service"
	"strconv"
)

func main() {
	startParam := config.GetStartParam()
	switch startParam.M {
	case "wa":
		router.SetupRouter(strconv.Itoa(startParam.Ap))
	case "ws":
		service.StartWS(strconv.Itoa(startParam.Sp))
	case "all":
		go router.SetupRouter(strconv.Itoa(startParam.Ap))
		service.StartWS(strconv.Itoa(startParam.Sp))
	}
}

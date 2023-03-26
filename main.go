/**
 * @Author think
 * @Date 2023/3/26$
 * @Note
 **/

package main

import (
	"server-resource-probe/router"
	"server-resource-probe/service"
)

func main() {
	go router.SetupRouter("8082")
	service.StartWS("8888")
}

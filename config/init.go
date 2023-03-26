/**
 * @Author root$
 * @Date 2023/3/26$
 * @Note 配置初始化文件
 **/

package config

import (
	"flag"
	"fmt"
	"os"
	"server-resource-probe/model"
)

var (
	version string
	h       bool
	v, V    bool
	m       string
	ap      int
	sp      int
)

func usage() {
	fmt.Fprintf(os.Stderr, `server resource probe version: `+version+`
Usage: server resource probe [-m start mode] [-ap WebApiServer start port] [-sp WebSocketServer start port]

Options:
`)
	flag.PrintDefaults()
}
func GetStartParam() model.ParamModel {
	version = "0.0.1"
	flag.BoolVar(&h, "h", false, "show help")

	flag.BoolVar(&v, "v", false, "show version and configure options then exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")
	flag.StringVar(&m, "m", "all", "StartMode:\n \twa: Start WebApiServer\n \tws: Start WebSocketServer\n \tall: Start WebApiServer and WebSocketServer)")
	flag.IntVar(&ap, "ap", 8082, "WebApiServer start port")
	flag.IntVar(&sp, "sp", 8083, "WebSocketServer start port")
	flag.Parse()
	if h {
		flag.Usage = usage
		flag.Usage()
		os.Exit(0)
	} else if v || V {
		fmt.Fprintf(os.Stderr, `server resource probe version: `+version)
		os.Exit(0)
	}
	if m != "wa" && m != "ws" && m != "all" {
		fmt.Fprintf(os.Stderr, `Error start model! please use -h show help. `)
		os.Exit(0)
	}
	return model.ParamModel{
		M:  m,
		Ap: ap,
		Sp: sp,
	}
}

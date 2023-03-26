/**
 * @Author root$
 * @Date 2023/3/26$
 * @Note 启动参数模型
 **/

package model

type ParamModel struct {
	M  string `json:"m"`  //启动模式
	Ap int    `json:"ap"` //WebApiServer启动端口
	Sp int    `json:"sp"` //WebSocket启动端口
}

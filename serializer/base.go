package serializer

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

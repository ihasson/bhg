package main

type sessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

func main() {

}

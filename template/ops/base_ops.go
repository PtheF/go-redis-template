package ops

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
)

type Ops struct {
	Pool *redigo.Pool
}

func (ops *Ops) Conn() redigo.Conn {
	return ops.Pool.Get()
}

func (ops *Ops) CloseConn(conn redigo.Conn) {
	_ = conn.Close()
}

func (ops *Ops) mergeArgs(key string, value ...interface{}) []interface{} {
	fmt.Printf("len(args)=%v\n", len(value))
	//args := make([]interface{}, len(value)+1)
	//args := []interface{}
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, value...)

	return args
}

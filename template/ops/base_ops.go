package ops

import (
	redigo "github.com/gomodule/redigo/redis"
)

type Ops struct {
	Pool *redigo.Pool
}

func (ops *Ops) conn() redigo.Conn {
	return ops.Pool.Get()
}

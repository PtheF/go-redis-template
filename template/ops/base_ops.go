package ops

import (
	redigo "github.com/gomodule/redigo/redis"
)

type Ops struct {
	Pool *redigo.Pool
}

func (ops *Ops) Conn() redigo.Conn {
	return ops.Pool.Get()
}

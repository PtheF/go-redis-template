package ops

import (
	"github.com/gomodule/redigo/redis"
)

type SetOps struct {
	Ops
}

func (o *SetOps) Add(key string, value ...interface{}) error {
	conn := o.Conn()

	defer o.CloseConn(conn)

	args := o.mergeArgs(key, value...)

	_, err := conn.Do("sadd", args...)

	return err
}

func (o *SetOps) IsMember(key, member string) (bool, error) {
	conn := o.Conn()

	defer o.CloseConn(conn)

	return redis.Bool(conn.Do("SISMEMBER", key, member))
}

func (o *SetOps) Members(key string) ([]string, error) {
	conn := o.Conn()

	defer o.CloseConn(conn)

	return redis.Strings(conn.Do("SMEMBERS", key))
}

func (o *SetOps) Rem(key string, members ...interface{}) error {
	conn := o.Conn()

	defer o.CloseConn(conn)

	_, err := conn.Do("SREM", o.mergeArgs(key, members...)...)

	return err
}

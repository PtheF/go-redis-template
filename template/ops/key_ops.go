package ops

import (
	redigo "github.com/gomodule/redigo/redis"
)

type KeyOps struct {
	Ops
}

func (ops *KeyOps) Expire(key string, expiredSec int64) error {
	conn := ops.Conn()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", key, expiredSec)
	return err
}

func (ops *KeyOps) Delete(keys ...interface{}) (bool, error) {
	conn := ops.Conn()
	defer conn.Close()

	v, err := redigo.Bool(conn.Do("DEL", keys...))
	return v, err
}

func (ops *KeyOps) Exist(key string) (v bool, err error) {
	conn := ops.Conn()
	defer conn.Close()

	return redigo.Bool(conn.Do("EXISTS", key))
}

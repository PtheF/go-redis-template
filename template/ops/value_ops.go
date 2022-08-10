package ops

import (
	redigo "github.com/gomodule/redigo/redis"
)

type ValueOps struct {
	Ops
}

const (
	SET = "SET"
	GET = "GET"
)

func (ops *ValueOps) Set(key string, value string) (err error) {
	conn := ops.conn()
	defer conn.Close()

	_, err = conn.Do(SET, key, value)

	return err
}

func (ops *ValueOps) Setnx(key string, value string) (err error) {
	conn := ops.conn()
	defer conn.Close()

	_, err = conn.Do("SETNX", key, value)

	return err
}

func (ops *ValueOps) Get(key string) (v string, err error) {
	conn := ops.conn()
	defer conn.Close()

	v, err = redigo.String(conn.Do(GET, key))

	return v, err
}

func (ops *ValueOps) Incr(key string) (v int, err error) {
	conn := ops.conn()
	defer conn.Close()

	return redigo.Int(conn.Do("INCR", key))
}

func (ops *ValueOps) IncrBy(key string, incr int) (v int, err error) {
	conn := ops.conn()
	defer conn.Close()

	return redigo.Int(conn.Do("INCRBY", key, incr))
}

func (ops *ValueOps) SetEx(key string, value string, exp int64) (bool, error) {
	conn := ops.conn()
	defer conn.Close()

	return redigo.Bool(conn.Do(SET, key, value, "ex", exp))
}

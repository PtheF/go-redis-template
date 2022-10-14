package ops

import "github.com/gomodule/redigo/redis"

type HashOps struct {
	Ops
}

func (o *HashOps) Set(key, field, value string) error {
	conn := o.Conn()

	defer conn.Close()

	_, err := conn.Do("hset", key, field, value)

	return err
}

func (o *HashOps) Get(key, field string) (string, error) {
	conn := o.Conn()

	defer conn.Close()

	return redis.String(conn.Do("hget", key, field))
}

func (o *HashOps) GetAll(key string) {
	conn := o.Conn()

	defer o.CloseConn(conn)

	conn.Do("HGETALL", key)
}

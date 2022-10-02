package ops

import redigo "github.com/gomodule/redigo/redis"

type ListOps struct {
	Ops
}

func (ops *ListOps) Lpush(key string, values ...interface{}) (len int, err error) {
	conn := ops.Conn()
	defer conn.Close()

	return redigo.Int(conn.Do("lpush", values))
}

func (ops *ListOps) Lpop(key string) (interface{}, error) {
	conn := ops.Conn()

	defer conn.Close()

	return conn.Do("LPOP", key)
}

func (ops *ListOps) Rpush(key string, values ...interface{}) (len int, err error) {
	conn := ops.Conn()
	defer conn.Close()

	return redigo.Int(conn.Do("RPUSH", key, values))
}

func (ops *ListOps) Rpop(key string) (interface{}, error) {
	conn := ops.Conn()
	defer conn.Close()

	return conn.Do("RPOP", key)
}

func (ops *ListOps) GetAll(key string) ([]interface{}, error) {
	conn := ops.Conn()
	defer conn.Close()

	return redigo.Values(conn.Do("LRANGE", key, 0, -1))
}

func (ops *ListOps) Len(key string) (int, error) {
	conn := ops.Conn()
	defer conn.Close()

	return redigo.Int(conn.Do("LLEN", key))
}

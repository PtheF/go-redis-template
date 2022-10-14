package ops

import redigo "github.com/gomodule/redigo/redis"

type ListOps struct {
	Ops
}

func (ops *ListOps) Lpush(key string, values ...interface{}) error {
	conn := ops.Conn()
	defer conn.Close()

	_, err := conn.Do("lpush", ops.mergeArgs(key, values...)...)

	return err
}

func (ops *ListOps) Lpop(key string) (string, error) {
	conn := ops.Conn()

	defer conn.Close()

	return redigo.String(conn.Do("LPOP", key))
}

func (ops *ListOps) Rpush(key string, values ...interface{}) error {
	conn := ops.Conn()
	defer conn.Close()

	_, err := conn.Do("RPUSH", ops.mergeArgs(key, values...)...)

	return err
}

func (ops *ListOps) Rpop(key string) (string, error) {
	conn := ops.Conn()
	defer conn.Close()

	return redigo.String(conn.Do("RPOP", key))
}

func (ops *ListOps) GetAll(key string) ([]string, error) {
	conn := ops.Conn()
	defer conn.Close()

	return redigo.Strings(conn.Do("LRANGE", key, 0, -1))
}

func (ops *ListOps) Len(key string) (int, error) {
	conn := ops.Conn()
	defer conn.Close()

	return redigo.Int(conn.Do("LLEN", key))
}

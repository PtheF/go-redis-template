package template

import (
	"errors"
	"github.com/PtheF/go-redis-template/template/ops"
	redigo "github.com/gomodule/redigo/redis"
)

var redTpl *RedTemp

type RedTemp struct {
	pool *redigo.Pool

	keyOps *ops.KeyOps

	valueOps *ops.ValueOps

	listOps *ops.ListOps

	//scriptOps *script.RedScriptOps
}

type RedTempBuilder struct {
	maxIdle     int
	maxActive   int
	idleTimeout int
	server      string
	auth        string
	db          int
}

func BuildRedTemp() *RedTempBuilder {
	return &RedTempBuilder{0, 0, 0, "127.0.0.1:6379", "", 0}
}

func (r *RedTempBuilder) Server(server string) *RedTempBuilder {
	r.server = server
	return r
}

func (r *RedTempBuilder) Auth(auth string) *RedTempBuilder {
	r.auth = auth
	return r
}

func (r *RedTempBuilder) DB(db int) *RedTempBuilder {
	r.db = db
	return r
}

func (r *RedTempBuilder) MaxIdle(maxIdle int) *RedTempBuilder {
	r.maxIdle = maxIdle
	return r
}

func (r *RedTempBuilder) MaxActive(maxActive int) *RedTempBuilder {
	r.maxActive = maxActive
	return r
}

func (r *RedTempBuilder) IdleTimeout(idleTimeout int) *RedTempBuilder {
	r.idleTimeout = idleTimeout
	return r
}

func (r *RedTempBuilder) Build() error {
	template := &RedTemp{}

	template.pool = GetRedisPool(r.server, r.auth, r.db, r.maxIdle, r.idleTimeout, r.maxActive)

	conn := template.pool.Get()

	_, err := conn.Do("PING")

	if err != nil {
		return errors.New("ConnectionError,CanNotGetPONG")
	}

	template.keyOps = &ops.KeyOps{}
	template.keyOps.Pool = template.pool

	template.valueOps = &ops.ValueOps{}
	template.valueOps.Pool = template.pool

	template.listOps = &ops.ListOps{}
	template.listOps.Pool = template.pool

	redTpl = template

	return nil
}

func OpsForValue() *ops.ValueOps {
	return redTpl.valueOps
}

func OpsForKey() *ops.KeyOps {
	return redTpl.keyOps
}

func OpsForList() *ops.ListOps {
	return redTpl.listOps
}

func Conn() redigo.Conn {
	return redTpl.pool.Get()
}

func Execute(redisScript string, keys []interface{}, args []interface{}) (interface{}, error) {
	conn := redTpl.pool.Get()
	lua := redigo.NewScript(len(keys), redisScript)

	keyLen := len(keys)
	argLen := len(args)

	r := make([]interface{}, keyLen+argLen)

	for i := 0; i < keyLen; i++ {
		r[i] = keys[i]
	}

	for i := 0; i < argLen; i++ {
		r[i+keyLen] = args[i]
	}

	return lua.Do(conn, r...)
	//exe, err := lua.Do(conn, r...)

	//return string(exe[:]), err
}

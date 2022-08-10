package template

import (
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

type RedisPoolError struct {
	msg       string
	auth      bool
	dbError   bool
	connError bool
}

func (e *RedisPoolError) Error() string {
	return e.msg
}

func (e *RedisPoolError) AuthError() bool {
	return e.auth
}

func (e *RedisPoolError) DBError() bool {
	return e.dbError
}

func (e *RedisPoolError) ConnError() bool {
	return e.connError
}

// GetRedisPool 初始化连接池
/*
  server: redis 地址
  pass: 密码
  db: 数据库
  maxIdle: 最大空闲
  idleTimeout: 最大空闲存活时间
  maxActive: 最大活动 -> 救急连接数 = maxActive - maxIdle
*/
func GetRedisPool(
	server string,
	pass string,
	db int,
	maxIdle int,
	idleTimeout int,
	maxActive int,
) *redigo.Pool {

	// 创建 Redigo 连接池，固定写法
	return &redigo.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		Wait:        true,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,

		/*
			  Dial: 连接工厂
				这个东西可以理解成线程池里面的线程工厂，用来返回连接的，
				文档可以直接去看这个 Pool 的源码，里面的注释很清楚了，要求返回的 conn 必须没有异常状态。
		*/
		Dial: func() (redigo.Conn, error) {

			// 这个才是核心的连接语句，只是我们在这里连接完以后再进行一轮测试
			c, err := redigo.Dial("tcp", server)

			if err != nil {
				return nil, &RedisPoolError{err.Error(), false, false, true}
			}

			// 带有密码的连接，固定写法，记住
			if pass != "" {
				_, err = c.Do("AUTH", pass)
				if err != nil {
					_ = c.Close()
					return nil, &RedisPoolError{err.Error(), true, false, false}
				}
			}

			// 选择一下数据库
			_, err = c.Do("SELECT", db)

			if err != nil {
				return nil, &RedisPoolError{err.Error(), false, true, false}
			}

			return c, nil
		},

		// 这个 TestOnBorrow 是空闲连接复用的时候要做的检查，检查复用的空闲连接是否健康
		// 如果出现了 error，则说明这个空闲连接完犊子了，可以关闭了，再创建个新的连接
		// 两个参数分别是即将复用的连接，和最大等待时间
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			// 可以看到这里验证健康还是很简单的，就是直接 PING 一下就可以了。
			_, err := c.Do("PING")
			return err
		},
	}
}

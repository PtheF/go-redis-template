package main

import (
	"fmt"
	"github.com/PtheF/go-redis-template/template"
	"reflect"
)

const (
	RedisScript = `
		local k1 = KEYS[1]
		local k2 = KEYS[2]

		local a1 = ARGV[1]
		local a2 = ARGV[2]

		redis.call('set', k1, a1)
		redis.call('set', k2, a2)

		local r1 = redis.call('get', k1)
		local r2 = redis.call('get', k2)

		-- return k1 .. ':' .. r1 .. '; ' .. k2 .. ':' .. r2

		-- return {a=1, b=2}

		return 1
	`
)

func main() {

	//a1 := []interface{}{1, 2, 3, 4, 5}
	//a2 := []interface{}{6, 7, 8, 9, 0}
	//
	//a3 := make([]interface{}, len(a1)+len(a2))
	//
	//i := 0
	//
	//for ; i < len(a1); i++ {
	//	a3[i] = a1[i]
	//}
	//
	//for i = 0; i < len(a2); i++ {
	//	a3[i+len(a1)] = a2[i]
	//}
	//
	//fmt.Printf("%v", a3)

	err := template.BuildRedTemp().MaxActive(20).MaxIdle(10).IdleTimeout(600).Build()

	if err != nil {
		fmt.Println("CreateRedTemplateError", err.Error())
	}

	keys := []interface{}{"name", "age"}

	execute, err := template.Execute(RedisScript, keys, []interface{}{"jack", 10})

	fmt.Printf("execute: %v, type: %t, err: %v\n", execute, execute, err)
	//fmt.Println(execute.(type))

	et := reflect.TypeOf(execute)

	println(et.Kind())
	println(et.Name())

	//bytes := execute.([]byte)
	//
	//res := string(bytes[:])
	//
	//fmt.Println(res)
}

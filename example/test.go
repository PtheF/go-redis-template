package main

import (
	"fmt"
	"go-redis-template/template"
)

func main() {
	err := template.BuildRedTemp().MaxActive(20).MaxIdle(10).IdleTimeout(600).Build()

	if err != nil {
		fmt.Println("CreateRedTemplateError", err.Error())
	}

	err = template.OpsForValue().Set("go-redis-template", "233333")

	if err != nil {
		fmt.Println("SetStringError: ", err)
	}

	err = template.OpsForKey().Expire("go-redis-template", 10)

	if err != nil {
		fmt.Println("SetExpiredError: ", err)
	}

	v, err := template.OpsForValue().Get("go-redis-template")

	if err != nil {
		fmt.Println("GetRedisKeyError: ", err)
	} else {
		fmt.Println("get value: ", v)
	}
}

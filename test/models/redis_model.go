package models

import (
	. "fmt"
	"github.com/PtheF/go-redis-template/template"
)

func init() {
	err := template.BuildRedTemp().MaxActive(20).MaxIdle(10).IdleTimeout(600).Build()

	if err != nil {
		Printf("init redis error: %v", err.Error())
	}
}

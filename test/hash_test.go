package test

import (
	. "fmt"
	"github.com/PtheF/go-redis-template/template"
	_ "github.com/PtheF/go-redis-template/test/models"
	"testing"
)

func TestHashOps(t *testing.T) {
	err := template.OpsForHash().Set("user:0001", "name", "jack")
	err = template.OpsForHash().Set("user:0001", "age", "20")

	if err != nil {
		Printf("hset err: %v", err)
	}

	name, _ := template.OpsForHash().Get("user:0001", "name")
	age, _ := template.OpsForHash().Get("user:0001", "age")

	Printf("user: name=%v, age=%v\n", name, age)

}

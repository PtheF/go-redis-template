package test

import (
	. "fmt"
	"github.com/PtheF/go-redis-template/template"
	_ "github.com/PtheF/go-redis-template/test/models"
	"testing"
)

func TestSetOps(t *testing.T) {
	if err := template.OpsForSet().Add("user", "u1", "u2", "u3"); err != nil {
		Printf("error %v", err)
	}

	isMember, _ := template.OpsForSet().IsMember("user", "u1")

	Printf("isMember: %v\n", isMember)

	members, _ := template.OpsForSet().Members("user")

	for i := range members {
		Printf("%v ", members[i])
	}

	Println()

	if err := template.OpsForSet().Rem("user", "u1"); err != nil {
		Printf("delete member error: %v", err)
	}

	Println("after delete")

	members, _ = template.OpsForSet().Members("user")

	for i := range members {
		Printf("%v ", members[i])
	}

	Println()
}

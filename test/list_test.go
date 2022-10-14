package test

import (
	. "fmt"
	"github.com/PtheF/go-redis-template/template"
	_ "github.com/PtheF/go-redis-template/test/models"
	"testing"
)

func TestListOps(t *testing.T) {
	if err := template.OpsForList().Lpush("ops_list", 1, 2, 3, 4); err != nil {
		Printf("lpush error: %v\n", err)
		t.Error(err)
	}

	if i, err := template.OpsForList().Len("ops_list"); err != nil {
		Printf("get len error: %v\n", err)
		t.Error(err)
	} else {
		Printf("len=%v\n", i)
	}

	res, _ := template.OpsForList().Lpop("ops_list")

	//res := lpop.(string)

	Printf("lpop = %v\n", res)

	all, _ := template.OpsForList().GetAll("ops_list")

	for i := range all {
		Printf("%v\n", all[i])
	}

}

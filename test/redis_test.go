package test

import (
	"github.com/PtheF/go-redis-template/template"
	_ "github.com/PtheF/go-redis-template/test/models"
	"testing"
)

func TestRedis(t *testing.T) {
	_ = template.OpsForValue().Set("name", "jack")
}

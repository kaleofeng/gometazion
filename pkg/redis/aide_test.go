package redis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAide_Do(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	var config Config
	err := config.Load("../../temp/redis.yaml")
	ast.NoError(err)

	aide := NewAide(config)
	err = aide.Open()
	ast.NoError(err)

	reply, err := aide.Do("info", "memory")
	ast.NoError(err)

	err = aide.Close()
	ast.NoError(err)

	fmt.Println(string(reply.([]byte)))
}

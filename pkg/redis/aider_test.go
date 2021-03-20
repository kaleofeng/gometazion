package redis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAider_Do(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	var config Config
	err := config.Load("../../temp/redis.yaml")
	ast.NoError(err)

	aider := NewAider(config)
	err = aider.Open()
	ast.NoError(err)

	reply, err := aider.Do("info", "memory")
	ast.NoError(err)

	err = aider.Close()
	ast.NoError(err)

	fmt.Println(string(reply.([]byte)))
}

package mysql

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestAider_Query(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	var config Config
	err := config.Load("../temp/mysql.yaml")
	ast.NoError(err)

	aider := NewAider(config)
	err = aider.Open()
	ast.NoError(err)

	rows, err := aider.Query("select * from user where user=?", "root")
	ast.NoError(err)

	err = aider.Close()
	ast.NoError(err)

	fmt.Println(rows.Columns())
}

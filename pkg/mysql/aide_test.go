package mysql

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestAide_Query(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	var config Config
	err := config.Load("../../temp/mysql.yaml")
	ast.NoError(err)

	aide := NewAide(config)
	err = aide.Open()
	ast.NoError(err)

	rows, err := aide.Query("select * from user where user=?", "root")
	ast.NoError(err)

	err = aide.Close()
	ast.NoError(err)

	fmt.Println(rows.Columns())
}

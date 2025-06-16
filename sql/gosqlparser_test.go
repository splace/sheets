package table

import (
	"fmt"
	sql "github.com/krasun/gosqlparser"
)

func Example_sql() {
	query, err := sql.Parse("SELECT col1, col2 FROM table1 WHERE col1 == \"abc\" AND col3 == 5 LIMIT 10")
	if err != nil {
		fmt.Printf("unexpected error: %s", err)
		return
	}
	fmt.Printf("%+v\n",query)
	// Output:
	// {{"Table":"table1","Columns":["col1","col2"],"Where":{"Expr":{"Left":{"Left":{"Name":"col1"},"Operator":0,"Right":{"Value":"\"abc\""}},"Operator":1,"Right":{"Left":{"Name":"col3"},"Operator":0,"Right":{"Value":"5"}}}},"Limit":"10"}

}

package table

import (
	"fmt"
)

func Example_sqlre() {
	fmt.Printf("%q\n",sqlre.FindStringSubmatch("SELECT * FROM foo_bar1 WHERE id < 3"))
	fmt.Printf("%q\n",sqlre.FindStringSubmatch("DELETE * FROM foo WHERE id < 3 AND id>1 LIMIT 10"))
	// Output:
	//
}





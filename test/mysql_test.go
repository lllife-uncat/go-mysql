package test

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "testing"
import "fmt"

func TestSimpleQuery(test *testing.T) {
	db, err := sql.Open("mysql", "root:1234@/db_aiadoc")

	if err != nil {
		fmt.Print(err)
		test.Fail()
	}
	rows , err := db.Query("select ID, user_name from user_table where user_name != ?", 5)
	if err != nil {
		fmt.Print(err)
		test.Fail()
	}

	for rows.Next() {
		var id int;
		var name string;
		rows.Scan(&id, &name)
		test.Log(id)
		test.Log(name)
	}
}


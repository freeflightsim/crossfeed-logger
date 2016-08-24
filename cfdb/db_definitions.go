
package cfdb


import (
	"fmt"


	///"database/sql"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var Dbx *sqlx.DB

// Initialise connection
func Init(user, password, database string) error {

	cred := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, database)

	var err error
	Dbx, err = sqlx.Connect("postgres", cred)
	//err := Dbx.Ping()
	if err != nil {
		fmt.Println("err", err)
	}
	return err
}


func CreateTables() {

	res, err := Dbx.Exec(SCHEMA_STAGING)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}





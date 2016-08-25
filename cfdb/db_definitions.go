
package cfdb


import (
	"fmt"
	"strings"
	"errors"

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

func CreateTables() error {

	var errs []string
	var e error

	e = CreateTable("staging", SCHEMA_STAGING, true)
	if e != nil {
		errs = append(errs, e.Error())
	}

	e = CreateTable("callsign", SCHEMA_CALLSIGN, true)
	if e != nil {
		errs = append(errs, e.Error())
	}

	e = CreateTable("flight", SCHEMA_FLIGHT, true)
	if e != nil {
		errs = append(errs, e.Error())
	}

	if len(errs) > 0 {
		return errors.New( strings.Join(errs, "\n"))
	}
	return nil
}

// Creates a database table
func CreateTable(table_name string, schema_sql string, drop bool) error {
	fmt.Println("t=", table_name)
	if drop {
		_, errd := Dbx.Exec("drop table if exists " +  table_name)
		if errd != nil {
			fmt.Println("errrd=", errd)
		} else {
			//fmt.Println(resd)
		}
	}

	_, errc := Dbx.Exec(schema_sql)
	if errc != nil {
		fmt.Println("errc=", errc)
	} else {
		//fmt.Println(res)
	}
	return errc
}





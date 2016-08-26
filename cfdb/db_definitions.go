
package cfdb


import (
	"fmt"
	//"strings"
	//"errors"

	///"database/sql"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var Dbx *sqlx.DB

// Initialise connection
func Init(user, password, database string) error {

	cred := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=require", user, password, database)

	var err error
	Dbx, err = sqlx.Connect("postgres", cred)
	//err := Dbx.Ping()
	if err != nil {
		fmt.Println("err", err)
	}
	return err
}

func DBCreateTables() []string {

	var errs []string
	var e error

	_, e = Dbx.Exec("CREATE EXTENSION postgis;")
	if e != nil {
		errs = append(errs, e.Error())
	}

	e = DBCreateTable("staging", SCHEMA_STAGING, true)
	if e != nil {
		errs = append(errs, e.Error())
	}
	DBIndexStaging()

	e = DBCreateTable("callsign", SCHEMA_CALLSIGN, true)
	if e != nil {
		errs = append(errs, e.Error())
	}
	DBIndexCallsign()

	e = DBCreateTable("aero", SCHEMA_AERO, true)
	if e != nil {
		errs = append(errs, e.Error())
	}
	DBIndexAero()

	e = DBCreateTable("flight", SCHEMA_FLIGHT, true)
	if e != nil {
		errs = append(errs, e.Error())
	}

	if len(errs) > 0 {
		fmt.Println("ERRORS=", errs)
		return errs
		//return errors.New( strings.Join(errs, "\n"))
	}
	return nil
}

// Creates a database table
func DBCreateTable(table_name string, schema_sql string, drop bool) error {
	fmt.Println("t=----------\n", table_name)
	if drop {
		_, errd := Dbx.Exec("DROP TABLE IF EXISTS " +  table_name)
		if errd != nil {
			fmt.Println("errrd=", errd)
		} else {
			//fmt.Println("Dropeed=" + table_name)
		}
	}

	_, errc := Dbx.Exec(schema_sql)
	if errc != nil {
		fmt.Println("CREATE ERROR=", errc)
	} else {
		//fmt.Println(res)
	}
	return errc
}

func DBInfo() map[string]string {

	ret := make(map[string]string)

	row := Dbx.QueryRow("select postgis_full_version();")

	var st string
	err := row.Scan(&st)
	if err != nil {
		fmt.Println("INFO+", err)
	}
	ret["postgis"] = st
	fmt.Println("INFO+", row)
	return ret
}

func DBAddIndex(table, col, name string){

	if name == "" {
		name = "idx_" + table + "_" + col
	} else {
		name = "idx_" + name
	}
	sql := "create index " + name + " on " + table + "(" + col + ")"
	res, errc := Dbx.Exec(sql)
	if errc != nil {
		fmt.Println("Index Err=", errc)
	} else {
		fmt.Println(sql, res)
	}

}


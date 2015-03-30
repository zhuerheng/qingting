package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
)

var db *sql.DB

func init() {
	//get password
	f, err := os.Open("../psd.in")
	if err != nil {
		panic(err)
	}
	psd, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	temp := string(psd)
	n := len(psd) - 1
	for temp[n] == '\n' {
		n--
	}
	defer f.Close()

	//connect to sql
	db, err = sql.Open("mysql", "root:"+string(psd[:n+1])+"@tcp(adserver1:3306)/ad_statistics")
	if err != nil {
		panic(err)
	}
	fmt.Println("database connect!")
}

func DbQueryFromAd_name(key string, fieldKnown string, fieldWant string) (value string, found bool, err error) {
	found = false
	rows, err := db.Query("select " + fieldWant + " from ad_name where " + fieldKnown + " = \"" + key + "\"")
	if err != nil {
		return "", found, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&value)
		found = true
	}

	return value, found, nil
}

func DbAdd(name string, url string) error {
	stmt, err := db.Prepare("insert into ad_name values(?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, url)
	if err != nil {
		return err
	}
	return nil
}

func DbQueryNum(datetime string, name string) (int, bool, error) {
	var found = false
	var num int
	rows, err := db.Query("select num from ad_name,ad_num where ad_num.url=ad_name.url and datetime=? and name=?", datetime, name)
	if err != nil {
		return 0, found, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&num)
		found = true
	}

	return num, found, nil

}

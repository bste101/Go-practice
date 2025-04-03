package main

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

type Cover struct {
	Id   int
	Name string
}

func main() {
	db, err := sql.Open("sqlserver", "sqlserver//sa:P@ssw0rd@localhost:1433?database=master")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	query := "select * from cover"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&Cover.Id, &Cover.Name)
		if err != nil {
			panic(err)
		}
		covers = append(covers, cover)
	}
	fmt.Printf("%#v", covers)
	// id := 0
	// name := ""
	// ok := rows.Next()
	// if ok {
	// 	rows.Scan(&id, &name)
	// }
	// rows.Close()

	fmt.Println(id, name)

}

func GetCovers() ([]Cover, error){
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	
}
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	id       int
	name     string
	address  string
	position string
}

func connectDB() (*sql.DB, error) {
	// template connection db
	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/golang_basic_sql")

	// jika tanpa password
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_basic_sql")

	db, err := sql.Open("mysql", "root:H3ru@mysql@tcp(127.0.0.1:3306)/golang_basic_sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("connection Sukses")
	return db, nil
}

func query() {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, name, address, position FROM employees")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var resp []employee

	for rows.Next() {
		var a = employee{}
		var err = rows.Scan(&a.id, &a.name, &a.address, &a.position)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// fmt.Println(a.id, a.name, a.address, a.position)

		resp = append(resp, a)
	}

	for _, r := range resp {
		fmt.Println(r.id, r.name, r.address, r.position)
	}
}

func queryRow() {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var employee = employee{}

	var id = "staff"

	err = db.QueryRow("Select id, name, address, position FROM employees WHERE position = ?", id).
		Scan(&employee.id, &employee.name, &employee.address, &employee.position)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(employee.id, employee.name, employee.address, employee.position)

}

func prepare() {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	q, err := db.Prepare("select id,name, address, position FROM employees where id = ?")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var employee1 = employee{}
	q.QueryRow(1).Scan(&employee1.id, &employee1.name, &employee1.address, &employee1.position)
	fmt.Println("Where ID 1", employee1.name, employee1.address, employee1.position)

	var employee2 = employee{}
	q.QueryRow(2).Scan(&employee2.id, &employee2.name, &employee2.address, &employee2.position)
	fmt.Println("Where ID 2", employee2.name, employee2.address, employee2.position)
}

func execInsert() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	name := "TrialInsertName"
	address := "TrialInsertAddress"
	position := "TrialInsertPosition"

	_, err = db.Exec("Insert into employees (name, position, address ) values (?, ?, ?)", name, position, address)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Insert Success")
}

func execUpdate() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	id := 4
	name := "TrialUpdateName"
	address := "TrialUpdateAddress"
	position := "TrialUpdatePosition"

	_, err = db.Exec("UPDATE employees SET name = ?, address = ?, position = ? WHERE id = ?", name, address, position, id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Update Success")
}

func execDelete() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	id := 4

	_, err = db.Exec("DELETE FROM employees WHERE id = ?", id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Delete Success")
}

func main() {
	// query()
	// queryRow()
	// prepare()
	// execInsert()
	// execUpdate()
	execDelete()
}

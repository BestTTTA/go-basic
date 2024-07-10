package main

import "fmt"

type (
	Database interface {
		Insert() error
	}

	RealDB struct{}
	MockDB struct{}
)

func (r RealDB) Insert() error {
	fmt.Println("Insert RealDB")
	return nil
}

func (m MockDB) Insert() error {
	fmt.Println("Insert MockDB")
	return nil
}

func InsertDB(db Database) error {
	return db.Insert()
}


func main() {

	r := RealDB{}
	m := MockDB{}

	InsertDB(r)
	InsertDB(m)



}
package main

import "fmt"

//easy to unit testing
type (
	Database interface {
		Insert() error
		Update() error
	}

	PostgresDb struct{} 

	MockDb struct{}
)


func NewProgresdb() Database{
	return &PostgresDb{}
} 
func NewMockdb() Database{
	return &MockDb{}
} 

func (p *PostgresDb) Insert() error {
	fmt.Println("Real Insert!!!")
	return nil
}

func (p *PostgresDb) Update() error {
	fmt.Println("Real Update!!!")
	return nil
}

func (m *MockDb) Insert() error {
	fmt.Println("Mock Insert!!!")
	return nil
}

func (m *MockDb) Update() error {
	fmt.Println("Mock Update!!!")
	return nil
}

func InsertPlayerItem(db Database) error {
	return db.Insert()
}

func UpdatePlayerItem(db Database) error {
	return db.Update()
}

func main() {
	postgres := &PostgresDb{}
	mock := &MockDb{}

	InsertPlayerItem(postgres)
	InsertPlayerItem(mock)

	UpdatePlayerItem(postgres)
	UpdatePlayerItem(mock)
}

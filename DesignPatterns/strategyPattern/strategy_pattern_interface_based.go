package main

import "fmt"

type IDBConnection interface {
	Connect()
}

type DBConnection struct {
	DB IDBConnection
}

func (d DBConnection) DBConnect() {
	d.DB.Connect()
}

type PostgreSQLConn struct {
	ConnectionString string
}

func (p PostgreSQLConn) Connect() {
	fmt.Println("PostgreSQL", p.ConnectionString)
}

type MongoDBConn struct {
	ConnectionString string
}

func (m MongoDBConn) Connect() {
	fmt.Println("MongoDB ", m.ConnectionString)
}

func main() {
	postgres := &PostgreSQLConn{ConnectionString: "Conected to PostgreSQL"}
	conn := &DBConnection{DB: postgres}
	conn.DB.Connect()

	mongodb := &MongoDBConn{ConnectionString: "Connected to MongoDB"}
	conn = &DBConnection{DB: mongodb}
	conn.DB.Connect()

}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
}

type PostgresDB struct {
	db *sql.DB
}

func newStorage() (*PostgresDB, error) {
	connPoint := "postgres://postgres:bankit@10.10.10.30/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connPoint)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func (s *PostgresDB) Init() error {
	return s.createAccountTable()
}

func (s *PostgresDB) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		balance real,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresDB) CreateAccount(acc *Account) error {
	query := `insert into account
	(first_name, last_name, created_at)
	values ($1, $2, $3)`

	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

package Persist

import (
	"fmt"
	"strings"
	"database/sql"
	_ "github.com/lib/pq"
)

type DBC struct{
	client *sql.DB
}	


func Initialize_db_connection() DBC{
	connStr := "user=james dbname=market_event sslmode=require"
	dbc, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error connection to timescale DB %v", err)
	}

	return DBC{
		client: dbc,
	}
}

func (dbc DBC) insert_market_event(time float64, price float64, quantity float64, side string, exchange_name string) {

	sqlStatement := `INSERT INTO btc_usd(time, price, quantity, side, exchange_id, contract_type) values($1, $2, $3, $4, $5, $6)`

	s := strings.Split(exchange_name, "_")
	exchange_id := s[0]
	contract_type := s[1]

	_, err := dbc.client.Exec(sqlStatement, time, price, quantity, side, exchange_id, contract_type)
	if err != nil {
		fmt.Printf("Errors writing to postgres %v\n", err)
	} 
}
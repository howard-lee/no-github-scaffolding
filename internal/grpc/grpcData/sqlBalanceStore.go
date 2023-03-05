package grpcData

import (
	"database/sql"
	"log"

	grpcApi "api"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLStore is a MongoDB data store which implements the Store interface
type MySQLStore struct {
	session *sql.DB
}

// NewMySQLStore creates an instance of MySQLStore with the given connection string
func NewMySQLStore(connection string) (*MySQLStore, error) {
	log.Println("Opening connection to:", connection)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Check connection is up
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &MySQLStore{session: db}, nil
}

// Search returns Balances from the MySQL instance which have the name name
func (m *MySQLStore) GetBalance(account string) []grpcApi.Balance {
	log.Println("GetBalance , account:", account)
	var results []grpcApi.Balance

	rows, err := m.session.Query("SELECT Account, Balance FROM Balances WHERE Account=?", account)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		bal := grpcApi.Balance{}
		rows.Scan(&bal.Account, &bal.Balance)
		results = append(results, bal)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	return results
}

// DeleteAllBalances deletes all the Balances from the datastore
// func (m *MySQLStore) DeleteAllBalances() {
// 	m.session.Exec("DELETE FROM Balances")
// }

// InsertBalances inserts a slice of Balances into the datastore
func (m *MySQLStore) InsertBalances(balances []grpcApi.Balance) error {
	for _, bal := range balances {
		_, err := m.session.Exec("INSERT INTO Balances (Account, Balance) VALUES (?, ?)",
			bal.Account,
			bal.Balance,
		)

		if err != nil {
			return err
		}
	}
	return nil
}

// CreateSchema creates the initial datbase schema
func (m *MySQLStore) CreateSchema() {
	m.session.Exec("DROP TABLE Balances")
	m.session.Exec("CREATE TABLE Balances (Account varchar(60),  Balance int)")
}

func (m *MySQLStore) GetWallets(account string) []grpcApi.Wallet {
	log.Println("Get Wallets , account:", account)
	var results []grpcApi.Wallet

	rows, err := m.session.Query("SELECT WalletID, URL FROM Wallets WHERE Account=?", account)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		bal := grpcApi.Wallet{}
		rows.Scan(&bal.WalletID, &bal.URL)
		results = append(results, bal)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	return results
}

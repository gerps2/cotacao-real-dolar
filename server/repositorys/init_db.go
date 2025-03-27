package repositorys

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"server/domain/errors"
)

func InitDB(dataSourceName string) (*sql.DB, *errors.CustomError) {
	db, err := sql.Open("sqlite3", dataSourceName)

	if err != nil {
		errConnection := errors.NewDatabaseConnetionError(err)
		return nil, &errConnection
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS cotacoes (
		        id INTEGER PRIMARY KEY AUTOINCREMENT,
		        code TEXT,
		        codein TEXT,
		        name TEXT,
		        high TEXT,
		        low TEXT,
		        varBid TEXT,
		        pctChange TEXT,
		        bid TEXT,
		        ask TEXT,
		        timestamp TEXT,
		        create_date TEXT
		    );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		errCreateTable := errors.NewDatabaseCreateTableError(err)
		return nil, &errCreateTable
	}

	return db, nil
}

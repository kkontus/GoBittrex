package databases

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	gbtError "GoBittrex/error"
)

type DB struct {
	DB *sql.DB
}

// example connect: gbtCryptoDb.Connect(gbtConfig.MYSQL_HOST, gbtConfig.MYSQL_PORT, gbtConfig.MYSQL_USERNAME, gbtConfig.MYSQL_PASSWORD, gbtConfig.MYSQL_DBNAME)
func Connect(host string, port string, username string, password string, dbname string) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	db, err := sql.Open("mysql", dataSource)
	return db, err
}

// import confing and db connector
//gbtConfig "GoBittrex/config"
//gbtCryptoDb "GoBittrex/databases"

//db, err := gbtCryptoDb.Connect(gbtConfig.MYSQL_HOST, gbtConfig.MYSQL_PORT, gbtConfig.MYSQL_USERNAME, gbtConfig.MYSQL_PASSWORD, gbtConfig.MYSQL_DBNAME)
//defer db.Close()
//gbtCryptoDb.Insert(db, err)
////gbtCryptoDb.FindAll(db, err)
//gbtCryptoDb.Find(db, err, 1)
//gbtCryptoDb.Find(db, err, 2)

func Find(db *sql.DB, err error, id int) {
	if err != nil {
		gbtError.ShowError(err)
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		gbtError.ShowError(err)
	}

	// prepare statement for reading data
	rows, err := db.Prepare("SELECT * FROM names WHERE id = ?")
	if err != nil {
		gbtError.ShowError(err)
	}

	defer rows.Close()

	var name string // we "scan" the result in here

	// Query the id of 13
	err = rows.QueryRow(1).Scan(&id, &name) // name WHERE id = 1
	if err != nil {
		gbtError.ShowError(err)
	}
	fmt.Printf("The name that we are looking for with id: %d is %s \n", id, name)

	// Query another number.. 2 maybe?
	err = rows.QueryRow(2).Scan(&id, &name) // name WHERE id = 2
	if err != nil {
		gbtError.ShowError(err) // proper error handling instead of panic in your app
	}
	fmt.Printf("The name that we are looking for with id: %d is %s \n", id, name)
}

func FindAll(db *sql.DB, err error) {
	if err != nil {
		gbtError.ShowError(err)
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		gbtError.ShowError(err)
	}

	// prepare statement for reading data
	rows, err := db.Query("SELECT id, name FROM names")
	if err != nil {
		gbtError.ShowError(err)
	}

	defer rows.Close()

	var id int
	var name string

	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Printf("%d : %s \n", id, name)
	}
}

func Insert(db *sql.DB, err error) {
	if err != nil {
		gbtError.ShowError(err)
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		gbtError.ShowError(err)
	}

	// we have two fields in table id and name so we need two placeholders or else we will get an error
	stmtIns, err := db.Prepare("INSERT INTO names VALUES( ?, ? )")
	if err != nil {
		gbtError.ShowError(err)
	}

	// close the statement when we leave main() // the program terminates
	defer stmtIns.Close()

	// our id field auto increments so we don't need to pass actual value for it.
	_, err = stmtIns.Exec(1, "Test")
	if err != nil {
		gbtError.ShowError(err)
	}
}

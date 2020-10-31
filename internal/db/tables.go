/*
 * Creates all necessary tables with alters and drops.
 */

package db

import (
	"fmt"
)

func SyncTables() {
	fmt.Println("Synchronizing database tables.")
	user()
}

func user() {
	var sql []string
	sql = append(sql, "CREATE TABLE IF NOT EXISTS account (username VARCHAR(50) PRIMARY KEY, password VARCHAR(1000));")
	ExecuteStatements(&sql)
}

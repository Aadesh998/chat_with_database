package services

import (
	"database/sql"
	"fmt"
	"strings"
)

func ExcecuteSQL(db *sql.DB, query string) bool {
	query = strings.TrimSpace(query)
	if strings.HasPrefix(strings.ToLower(query), "select") {
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println("Error executing select query:", err)
			return false
		}
		defer rows.Close()

		cols, err := rows.Columns()
		if err != nil {
			fmt.Println("Error getting columns:", err)
			return false
		}

		for rows.Next() {
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i := range columns {
				columnPointers[i] = &columns[i]
			}

			if err := rows.Scan(columnPointers...); err != nil {
				fmt.Println("Error scanning row:", err)
				continue
			}

			for i, colName := range cols {
				val := columns[i]
				switch v := val.(type) {
				case []byte:
					fmt.Printf("%s: %s\t", colName, string(v))
				default:
					fmt.Printf("%s: %v\t", colName, v)
				}
			}
			fmt.Println()
		}
		return true

	}

	results, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return false
	}
	lastInsertID, _ := results.LastInsertId()
	rowsAffected, _ := results.RowsAffected()

	fmt.Printf("Last Insert ID: %d, Rows Affected: %d\n", lastInsertID, rowsAffected)

	return true
}

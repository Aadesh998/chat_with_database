package dbfunction

import (
	"database/sql"
	"log"

	"github.com/Aadesh-lab/views"
)

func GetFullSchema(sqlDB *sql.DB, schema string) ([]views.TableSchema, error) {
	rows, err := sqlDB.Query(`
		SELECT
			table_name,
			column_name,
			data_type,
			is_nullable
		FROM information_schema.columns
		WHERE table_schema = $1
		ORDER BY table_name, ordinal_position
	`, schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	log.Println(rows)

	tableMap := make(map[string][]views.ColumnSchema)

	for rows.Next() {
		var table, column, dataType, nullable string
		if err := rows.Scan(&table, &column, &dataType, &nullable); err != nil {
			return nil, err
		}

		tableMap[table] = append(tableMap[table], views.ColumnSchema{
			ColumnName: column,
			DataType:   dataType,
			IsNullable: nullable,
		})
	}

	var result []views.TableSchema
	for table, cols := range tableMap {
		result = append(result, views.TableSchema{
			TableName: table,
			Columns:   cols,
		})
	}

	return result, nil
}

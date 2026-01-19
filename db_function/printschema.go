package dbfunction

import (
	"fmt"

	"github.com/Aadesh-lab/views"
)

func PrintSchema(schema []views.TableSchema) {
	for _, table := range schema {
		fmt.Println("-------------------------------------------------")
		fmt.Println("TABLE:", table.TableName)
		for _, col := range table.Columns {
			fmt.Printf("  - %s (%s) nullable=%s\n",
				col.ColumnName,
				col.DataType,
				col.IsNullable,
			)
		}
	}
}

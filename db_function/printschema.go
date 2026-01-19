package dbfunction

import (
	"fmt"
	"strings"

	"github.com/Aadesh-lab/views"
)

func PrintSchema(schema []views.TableSchema) {
	fmt.Println(SchemaToString(schema))
}

func SchemaToString(schema []views.TableSchema) string {
	var builder strings.Builder
	for _, table := range schema {
		builder.WriteString("-------------------------------------------------\n")
		builder.WriteString(fmt.Sprintf("TABLE: %s\n", table.TableName))
		for _, col := range table.Columns {
			builder.WriteString(fmt.Sprintf("  - %s (%s) nullable=%s\n",
				col.ColumnName,
				col.DataType,
				col.IsNullable,
			))
		}
	}
	return builder.String()
}

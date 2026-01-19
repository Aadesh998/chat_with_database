package views

type TableSchema struct {
	TableName string
	Columns   []ColumnSchema
}

type ColumnSchema struct {
	ColumnName string
	DataType   string
	IsNullable string
}

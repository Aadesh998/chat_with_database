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

type WSResponse struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type QueryResult struct {
	SQL           string                   `json:"sql"`
	Columns       []string                 `json:"columns"`
	Rows          []map[string]interface{} `json:"rows"`
	Visualization string                   `json:"visualization"`
}

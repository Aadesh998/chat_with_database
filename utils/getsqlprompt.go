package utils

func GetSQLPrompt(schema, userQuery string) string {
	return `
You are an expert SQL generator.

Rules:
- Output ONLY the SQL query
- Do NOT add explanations, comments, or text
- Do NOT include markdown or code fences
- Do NOT include the word "SQL"
- Output must be a single valid SQL statement

Schema:
` + schema + `

User request:
` + userQuery
}

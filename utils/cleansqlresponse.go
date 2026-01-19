package utils

import "strings"

func CleanSQLResponse(response string) string {
	sql := strings.TrimSpace(response)

	// Remove starting ```sql or ```
	if strings.HasPrefix(sql, "```sql") {
		sql = strings.TrimPrefix(sql, "```sql")
	} else if strings.HasPrefix(sql, "```") {
		sql = strings.TrimPrefix(sql, "```")
	}

	// Remove ending ```
	if strings.HasSuffix(sql, "```") {
		sql = strings.TrimSuffix(sql, "```")
	}

	return strings.TrimSpace(sql)
}

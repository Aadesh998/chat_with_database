package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Aadesh-lab/db"
	dbfunction "github.com/Aadesh-lab/db_function"
	"github.com/Aadesh-lab/envloader"
	"github.com/Aadesh-lab/views"
)

func main() {
	envloader.LoadConfig()
	db.InitDB()

	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB:", err)
	}

	schema, err := dbfunction.GetFullSchema(sqlDB, envloader.AppConfig.DBSchema)
	printSchema(schema)
}

func printSchema(schema []views.TableSchema) {
	db_table_name := []string{
		"odh_aop_record",
		"odh_breach_record",
		"odh_cod_risk_record",
		"odh_conversion_record",
		"odh_cv_adherence_record",
		"odh_cxns_record",
		"odh_elk_fwd_ofd_record",
		"odh_elk_fwd_record",
		"odh_elk_rev_ofd_record",
		"odh_elk_rev_record",
		"odh_emo_record",
		"odh_indicator",
		"odh_ipd_record",
		"odh_lps_record",
		"odh_promise_record",
		"odh_ranking_config",
	}

	for _, tablename := range db_table_name {
		fmt.Println(schemaToString(tablename, schema))
	}
}

func schemaToString(tablename string, schema []views.TableSchema) string {
	var builder strings.Builder
	for _, table := range schema {
		if table.TableName == tablename {
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
	}
	return builder.String()
}

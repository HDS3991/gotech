package sql2gorm

import (
	"database/sql"
	"github.com/pkg/errors"
)

func GetTableFromDB(dsn, tableName string) (string, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return "", errors.WithMessage(err, "open db error")
	}
	defer db.Close()
	rows, err := db.Query("SHOW CREATE TABLE " + tableName)
	if err != nil {
		return "", errors.WithMessage(err, "query show create table error")
	}
	defer rows.Close()
	if !rows.Next() {
		return "", errors.Errorf("table(%s) not found", tableName)
	}

	var table, createSql string
	err = rows.Scan(&table, &createSql)
	if err != nil {
		return "", err
	}
	return createSql, nil
}

func ParseSqlFromDB(dsn, tableName string, options ...Option) (*ModelCode, error) {
	createSql, err := GetTableFromDB(dsn, tableName)
	if err != nil {
		return nil, err
	}
	return Parse(createSql, options...)
}

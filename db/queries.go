package db

import (
	"fmt"
	"strconv"

	"github.com/omrfrkazt/lotr-vanilla/models"
)

func (model *DBModel) GetAllQuery(tableName string) string {
	return `SELECT * FROM ` + tableName
}

func (model *DBModel) GetAllCustomSelectQuery(tableName string, selectFields string) string {
	return `SELECT ` + selectFields + ` FROM ` + tableName + `WHERE is_active = true` + ` ORDER BY id DESC`
}

func (model *DBModel) GetByIdQuery(tableName string, id int) string {
	fmt.Println("calisti id")
	return `SELECT * FROM ` + tableName + ` WHERE id = ` + strconv.Itoa(id) + ` and is_active = true` + ` LIMIT 1`
}

func (model *DBModel) GetByIdCustomSelectQuery(tableName string, id int, selectFields string) string {
	return `SELECT ` + selectFields + ` FROM ` + tableName + ` WHERE id = ` + string(rune(id)) + ` and is_active = true` + ` ORDER BY id DESC`
}

func (model *DBModel) AddQuery(tableName string, base models.Base) string {
	return `INSERT INTO ` + tableName + ` (` + base.GetColumns() + `) VALUES (` + base.GetValues() + `) RETURNING *`
}

func (model *DBModel) UpdateQuery(tableName string, id int, columns string, values string) string {
	return `UPDATE ` + tableName + ` SET ` + columns + ` = ` + values + ` WHERE id = ` + string(rune(id)) + ` RETURNING *`
}

func (model *DBModel) DeleteQuery(tableName string, id int) string {
	return `UPDATE ` + tableName + ` SET is_active = false WHERE id = ` + string(rune(id)) + ` RETURNING *`
}

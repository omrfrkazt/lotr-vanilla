package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/omrfrkazt/lotr-vanilla/models"
)

func (model *DBModel) GetFellowship(id string) (result []models.FellowshipMember, err error) {
	var query string
	if id == "" {
		query = model.GetAllQuery(models.Fellowship)
	} else if intId, err := strconv.Atoi(id); err == nil {
		query = model.GetByIdQuery(models.Fellowship, intId)
	}
	fmt.Println(query)
	rows, err := model.DB.QueryContext(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()
	var fellowship models.FellowshipMember
	for rows.Next() {
		err = rows.Scan(&fellowship.Id, &fellowship.Name, &fellowship.Surname, &fellowship.CreatedAt, &fellowship.IsActive)
		if err != nil {
			return
		}
		result = append(result, fellowship)
	}
	return
}

func (model *DBModel) AddFellowship(reader io.ReadCloser) sql.Result {
	var fellowship models.FellowshipMember

	err := json.NewDecoder(reader).Decode(&fellowship)
	if err != nil {
		panic(err)
	}
	query := model.AddQuery(models.Fellowship, &fellowship)
	fmt.Println(query)
	result, err := model.DB.ExecContext(context.Background(), query)
	if err != nil {
		panic(err)
	}
	return result
}

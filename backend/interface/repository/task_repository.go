package repository

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	. "github.com/jinzhu/gorm"
)

func FetchTasksCount() (int, error){
	var cnt int
	db := database.NewDB()
	qs := db.Model(&model.Task{})
	err := qs.Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func FindTasks(scopes []func(db *DB) *DB) ([]*model.Task, error) {
	var tasks []*model.Task
	db := database.NewDB()
	rows, err := db.Model(&model.Task{}).Scopes(scopes...).Rows()
	if err != nil {
		return []*model.Task{}, err
	}
	for rows.Next() {
		task := &model.Task{}
		err := db.ScanRows(rows, task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

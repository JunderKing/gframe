package service

import (
	"gframe/model"
)

type Demo struct {
}

func (s *Demo) CreateUser(name string, age int, sex int) *model.Demo {
	userModel := &model.Demo{
		Name: name,
		Age:  age,
		Sex:  sex,
	}
	model.Db.Create(userModel)
	return userModel
}

func (s *Demo) FindUserById(userId int) *model.Demo {
	userModel := &model.Demo{}
	model.Db.First(userModel, userId)
	return userModel
}

func (s *Demo) UpdateUserName(userId int, name string) int {
	userModel := &model.Demo{Id: userId}
	tx := model.Db.Model(userModel).Update("Name", name)
	return int(tx.RowsAffected)
}

func (s *Demo) DeleteUser(userId int) {
	model.Db.Delete(&model.Demo{}, userId)
}

package usecase

import (
	"fmt"
	"usecase-1/model"
	"usecase-1/repository"
)

type Usecase1UseCase interface {
	RegisterNewU1(payload model.Usecase1Model) (model.Usecase1Model, error)
	FindAllU1() []model.Usecase1Model
	FindByIdU1(id int) model.Usecase1Model
	DeleteByIdU1(id int) string
}

type usecase1UseCase struct {
	repo repository.Usecase1Repository
}

func (e *usecase1UseCase) RegisterNewU1(payload model.Usecase1Model) (model.Usecase1Model, error) {
	//pengecekan nama tidak boleh kosong
	if payload.Task == "" {
		return payload, fmt.Errorf("name required fields")
	}

	data := e.repo.Create(payload)
	return data, nil
}

func (e *usecase1UseCase) FindAllU1() []model.Usecase1Model {
	return e.repo.List()
}

func (e *usecase1UseCase) FindByIdU1(id int) model.Usecase1Model {
	return e.repo.GetByID(id)
}

func (e *usecase1UseCase) DeleteByIdU1(id int) string {
	return e.repo.Delete(id)
}

func NewU1UseCase(repo repository.Usecase1Repository) Usecase1UseCase {
	return &usecase1UseCase{repo: repo}
}

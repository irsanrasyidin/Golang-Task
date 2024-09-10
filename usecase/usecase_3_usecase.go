package usecase

import (
	"usecase-1/repository"
)

type Usecase3UseCase interface {
	RegisterNewU3(filename string, fileData []byte) error
	FindAllU3() ([]string, error)
	FindByIdU3(filename string) ([]byte, error)
}

type usecase3UseCase struct {
	repo repository.Usecase3Repository
}

func (e *usecase3UseCase) RegisterNewU3(filename string, fileData []byte) error {
	return e.repo.Create(filename, fileData)
}

func (e *usecase3UseCase) FindAllU3() ([]string, error) {
	return e.repo.List()
}

func (e *usecase3UseCase) FindByIdU3(filename string) ([]byte, error) {
	return e.repo.GetByID(filename)
}

func NewU3UseCase(repo repository.Usecase3Repository) Usecase3UseCase {
	return &usecase3UseCase{repo: repo}
}

package usecase

import (
	"fmt"
	"usecase-1/model"
	"usecase-1/repository"
	"usecase-1/utils/security"
)

type Usecase2UseCase interface {
	RegisterNewU1(payload model.Usecase2RegisterModel) (string, error)
	LoginNewU1(payload model.Usecase2LoginModel) (string, error)
	FindByUsernameU2(username string) model.Usecase2Model
	DeleteByIdU1(username string) string
}

type usecase2UseCase struct {
	repo repository.Usecase2Repository
}

func (e *usecase2UseCase) RegisterNewU1(payload model.Usecase2RegisterModel) (string, error) {
	return e.repo.Create(payload), nil
}

func (e *usecase2UseCase) LoginNewU1(payload model.Usecase2LoginModel) (string, error) {
	status := e.repo.FindByUsernamePassword(payload)
	if !status {
		return "", fmt.Errorf("invalid username:" + payload.Username)
	}

	// mekanisme jika user itu ada akan membalikan sebuah token
	token, err := security.CreateAccessToken(payload)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return token, nil
}

func (e *usecase2UseCase) FindByUsernameU2(username string) model.Usecase2Model {
	return e.repo.GetByUsername(username)
}

func (e *usecase2UseCase) DeleteByIdU1(username string) string {
	return e.repo.Delete(username)
}

func NewU2UseCase(repo repository.Usecase2Repository) Usecase2UseCase {
	return &usecase2UseCase{repo: repo}
}

package manager

import (
	"sync"
	"usecase-1/model"
	"usecase-1/repository"
)

type RepoManager interface {
	NewU1Repository() repository.Usecase1Repository
	NewU2Repository() repository.Usecase2Repository
	NewU3Repository() repository.Usecase3Repository
}

type repoManager struct {
	infra     InfraManager
	DataTask  *[]model.Usecase1Model
	DataUser  *[]model.Usecase2Model
	DataLogin *[]model.Usecase2LoginModel

	U1Repo repository.Usecase1Repository
	U2Repo repository.Usecase2Repository
	U3Repo repository.Usecase3Repository
}

var onceLoadUsecase1Repo sync.Once
var onceLoadUsecase2Repo sync.Once
var onceLoadUsecase3Repo sync.Once

// TransactionLeaveRepo implements RepoManager.
func (r *repoManager) NewU1Repository() repository.Usecase1Repository {
	onceLoadUsecase1Repo.Do(func() {
		r.U1Repo = repository.NewU1Repository(r.DataTask)
	})
	return r.U1Repo
}

func (r *repoManager) NewU2Repository() repository.Usecase2Repository {
	onceLoadUsecase2Repo.Do(func() {
		r.U2Repo = repository.NewU2Repository(r.DataUser, r.DataLogin)
	})
	return r.U2Repo
}

func (r *repoManager) NewU3Repository() repository.Usecase3Repository {
	onceLoadUsecase3Repo.Do(func() {
		r.U3Repo = repository.NewU3Repository()
	})
	return r.U3Repo
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra:     infra,
		DataTask:  &[]model.Usecase1Model{},
		DataUser:  &[]model.Usecase2Model{},
		DataLogin: &[]model.Usecase2LoginModel{},
	}
}

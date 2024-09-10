package manager

import (
	"sync"
	"usecase-1/usecase"
)

type UseCaseManager interface {
	NewU1UseCase() usecase.Usecase1UseCase
	NewU2UseCase() usecase.Usecase2UseCase
	NewU3UseCase() usecase.Usecase3UseCase
}

type useCaseManager struct {
	repoManager RepoManager

	U1Usecase usecase.Usecase1UseCase
	U2Usecase usecase.Usecase2UseCase
	U3Usecase usecase.Usecase3UseCase
}

var onceLoadUsecase1Usecase sync.Once
var onceLoadUsecase2Usecase sync.Once
var onceLoadUsecase3Usecase sync.Once

// TransactionUseCase implements UseCaseManager.
func (u *useCaseManager) NewU1UseCase() usecase.Usecase1UseCase {
	onceLoadUsecase1Usecase.Do(func() {
		u.U1Usecase = usecase.NewU1UseCase(u.repoManager.NewU1Repository())
	})
	return u.U1Usecase
}

func (u *useCaseManager) NewU2UseCase() usecase.Usecase2UseCase {
	onceLoadUsecase2Usecase.Do(func() {
		u.U2Usecase = usecase.NewU2UseCase(u.repoManager.NewU2Repository())
	})
	return u.U2Usecase
}

func (u *useCaseManager) NewU3UseCase() usecase.Usecase3UseCase {
	onceLoadUsecase3Usecase.Do(func() {
		u.U3Usecase = usecase.NewU3UseCase(u.repoManager.NewU3Repository())
	})
	return u.U3Usecase
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}

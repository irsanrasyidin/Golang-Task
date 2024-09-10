package manager

import (
	"fmt"
	"sync"
	"usecase-1/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	Conn() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg *config.Config
}

var onceLoadDb sync.Once

func (i *infraManager) initDb() error {
	var initErr error
	onceLoadDb.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			initErr = err // Simpan error di variabel luar
			return
		}
		i.db = db
	})
	return initErr
}

func (i *infraManager) Conn() *gorm.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

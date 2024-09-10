package repository

import (
	"math/rand"
	"strconv"
	"time"
	"usecase-1/model"
)

type Usecase1Repository interface {
	Create(payload model.Usecase1Model) model.Usecase1Model
	List() []model.Usecase1Model
	GetByID(id int) model.Usecase1Model
	Delete(id int) string
}

type usecase1Repository struct {
	DataTask *[]model.Usecase1Model
}

func (e *usecase1Repository) Create(payload model.Usecase1Model) model.Usecase1Model {
	Length := len(*e.DataTask)
	payload.ID = Length + 1
	payload.Mulai = time.Now()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomHours := r.Intn(5) + 1
	payload.Selesai = payload.Mulai.Add(time.Duration(randomHours) * time.Hour)
	payload.Durasi = payload.Selesai.Sub(payload.Mulai)

	*e.DataTask = append(*e.DataTask, payload)
	return payload
}

func (e *usecase1Repository) List() []model.Usecase1Model {
	return *e.DataTask
}

func (e *usecase1Repository) GetByID(id int) model.Usecase1Model {
	var DataTask model.Usecase1Model
	for _, v := range *e.DataTask {
		if v.ID == id {
			DataTask = v
		}
	}
	return DataTask
}

func (e *usecase1Repository) Delete(id int) string {
	var DataTask []model.Usecase1Model
	var Status bool
	for _, v := range *e.DataTask {
		Status = false
		if v.ID != id {
			DataTask = append(DataTask, v)
		} else {
			Status = true
		}
	}

	e.DataTask = &DataTask

	if Status {
		return "Data dengan ID:" + strconv.Itoa(id) + " berhasil di hapus"
	} else {
		return "Tidak ada data yang terhapus"
	}
}

func NewU1Repository(DataTask *[]model.Usecase1Model) Usecase1Repository {
	return &usecase1Repository{DataTask: DataTask}
}

package repository

import (
	"fmt"
	"log"
	"time"
	"usecase-1/model"

	"golang.org/x/crypto/bcrypt"
)

type Usecase2Repository interface {
	Create(payload model.Usecase2RegisterModel) string
	FindByUsernamePassword(payload model.Usecase2LoginModel) bool
	GetByUsername(username string) model.Usecase2Model
	Delete(username string) string
}

type usecase2Repository struct {
	DataUser  *[]model.Usecase2Model
	DataLogin *[]model.Usecase2LoginModel
}

func (e *usecase2Repository) Create(payload model.Usecase2RegisterModel) string {
	var DataUser model.Usecase2Model
	var DataLogin model.Usecase2LoginModel
	DataUser.ID = payload.ID
	DataUser.Username = payload.Username
	DataUser.Nama = payload.Nama
	t, _ := time.Parse("2006-01-02", payload.Tgl_Lahir)
	DataUser.Tgl_Lahir = t
	DataLogin.ID = payload.ID
	DataLogin.Username = payload.Username
	DataLogin.Password = payload.Password
	*e.DataUser = append(*e.DataUser, DataUser)
	*e.DataLogin = append(*e.DataLogin, DataLogin)
	return payload.Username
}

func (e *usecase2Repository) FindByUsernamePassword(payload model.Usecase2LoginModel) bool {
	var status bool
	for _, v := range *e.DataLogin {
		status = false
		if v.Username == payload.Username {
			err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(payload.Password))
			if err != nil {
				if err == bcrypt.ErrMismatchedHashAndPassword {
					fmt.Println("Password does not match")
					status = false
					break
				} else {
					log.Fatalf("Failed to compare password: %v", err)
					status = false
					break
				}
			} else {
				fmt.Println("Password matches")
				status = true
				break
			}
		}
	}
	return status
}

func (e *usecase2Repository) GetByUsername(username string) model.Usecase2Model {
	var DataUser model.Usecase2Model
	for _, v := range *e.DataUser {
		if v.Username == username {
			DataUser = v
		}
	}
	return DataUser
}

func (e *usecase2Repository) Delete(username string) string {
	var DataUser []model.Usecase2Model
	var Status bool
	for _, v := range *e.DataUser {
		Status = false
		if v.Username != username {
			DataUser = append(DataUser, v)
		} else {
			Status = true
		}
	}

	e.DataUser = &DataUser

	if Status {
		return "Data dengan Username:" + username + " berhasil di hapus"
	} else {
		return "Tidak ada data yang terhapus"
	}
}

func NewU2Repository(DataUser *[]model.Usecase2Model, DataLogin *[]model.Usecase2LoginModel) Usecase2Repository {
	return &usecase2Repository{
		DataUser:  DataUser,
		DataLogin: DataLogin,
	}
}

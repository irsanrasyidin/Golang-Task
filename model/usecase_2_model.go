package model

import "time"

type Usecase2Model struct {
	ID        string
	Username  string
	Nama      string
	Tgl_Lahir time.Time
}

type Usecase2LoginModel struct {
	ID       string
	Username string
	Password string
}

type Usecase2RegisterModel struct {
	ID        string
	Username  string
	Password  string
	Nama      string
	Tgl_Lahir string
}

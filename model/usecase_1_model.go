package model

import "time"

type Usecase1Model struct {
	ID      int
	Task    string
	Mulai   time.Time
	Selesai time.Time
	Durasi  time.Duration
}

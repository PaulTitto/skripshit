package progress

import "time"

type ProgressUser struct {
	IdProgress int
	UserId     int
	IdMaterial int
	Status     string
	DateStatus time.Time
}

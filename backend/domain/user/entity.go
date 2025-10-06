package user

import "time"

type User struct {
	Id               int
	Username         string
	Email            string
	PasswordHash     string
	Streak           int
	Avatar_file_name string
	LastDateLearn    string
	Role             string
	Created_at       time.Time
	Updated_at       time.Time
}

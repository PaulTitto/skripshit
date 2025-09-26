package material

import "time"

type Material struct {
	IdMaterial      int
	IdChapter       int
	NameMaterial    string
	XP              int
	Status          string
	OrderMaterialAt time.Time
}

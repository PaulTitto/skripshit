package chapter

import "time"

type Chapter struct {
	IdChapter      int
	NameChapter    string
	Description    string
	OrderChapterAt time.Time
}

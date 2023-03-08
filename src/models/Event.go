package models

type Event struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	ShortDesc string `json:"short_desc"`
	LongDesc  string `json:"long_desc"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Organizer string `json:"organizer"`
	Place     string `json:"place"`
	Status    string `json:"status"`
}

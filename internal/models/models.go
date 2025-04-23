package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"user_id"`
	Username  string     `gorm:"column:username;size:32;not null; unique" json:"username"`
	Password  string     `gorm:"column:password;size:128;not null" json:"password"`
	Phone     string     `gorm:"column:phone;size:15" json:"phone"`
	First     string     `gorm:"column:first;size:50" json:"first"`
	Last      string     `gorm:"column:last;size:50" json:"last"`
	Address   string     `gorm:"column:address;size:100" json:"address"`
	City      string     `gorm:"column:city;size:50" json:"city"`
	Country   string     `gorm:"column:country;size:50" json:"country"`
	Reminders []Reminder `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}

type Reminder struct {
	ID          uint      `gorm:"primaryKey" json:"primaryKey"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	Title       string    `gorm:"column:title;size 100;not null" json:"title"`
	Description string    `gorm:"column:description;size:255" json:"description"`
	Address     string    `gorm:"column:address;size:100" json:"address"`
	City        string    `gorm:"column:city;size:50" json:"city"`
	Zip         string    `gorm:"column:zip;size:10" json:"zip"`
	Country     string    `gorm:"column:country;size:50" json:"country"`
	Time        time.Time `gorm:"column:time;not null" json:"time"`
	RemindAt    time.Time `gorm:"column:remind_at" json:"remind_at"`
	TravelTime  time.Time `gorm:"column:travel_time" json:"travel_time"`
	Repeating   bool      `gorm:"column:repeating" json:"repeating"`
	EndDate     time.Time `gorm:"column:end_date" json:"end_date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Reminder) TableName() string {
	return "reminders"
}

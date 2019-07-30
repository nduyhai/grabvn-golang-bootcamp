package feedbackserver

import "github.com/jinzhu/gorm"

type Feedback struct {
	gorm.Model
	ID          string
	PassengerID string
	BookingCode string
	Feedback    string
}

package feedbackserver

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
	"log"
)

type server struct {
	DB *gorm.DB
}

func (s *server) Add(ctx context.Context, in *feedback.CreateFeedbackRequest) (*feedback.FeedbackResponse, error) {
	res := Feedback{
		ID:          uuid.New().String(),
		PassengerID: in.PassengerID,
		BookingCode: in.BookingCode,
		Feedback:    in.Feedback,
	}

	err := s.DB.Create(&res).Error
	if err == nil {
		return &feedback.FeedbackResponse{ID: res.ID}, nil
	} else {
		log.Print(err)
		return &feedback.FeedbackResponse{ID: res.ID}, nil
	}
}

func (s *server) GetById(ctx context.Context, in *feedback.FeedbackRequest) (*feedback.PassengerFeedback, error) {
	var res Feedback
	err := s.DB.Where("ID = ?", in.ID).Find(&res).GetErrors()
	if len(err) == 0 {
		fb := feedback.PassengerFeedback{
			ID:          res.ID,
			PassengerID: res.PassengerID,
			BookingCode: res.BookingCode,
			Feedback:    res.Feedback,
		}
		return &fb, nil

	} else {
		return &feedback.PassengerFeedback{ID: in.ID}, errors.New("not found")
	}
}
func (s *server) GetByBookingCode(ctx context.Context, in *feedback.BookingRequest) (*feedback.PassengerFeedback, error) {
	var res Feedback
	err := s.DB.Where("BookingCode = ?", in.Code).Find(&res).GetErrors()
	if len(err) == 0 {
		fb := feedback.PassengerFeedback{
			ID:          res.ID,
			PassengerID: res.PassengerID,
			BookingCode: res.BookingCode,
			Feedback:    res.Feedback,
		}
		return &fb, nil

	} else {
		return &feedback.PassengerFeedback{BookingCode: in.Code}, errors.New("not found")
	}

}
func (s *server) Delete(ctx context.Context, in *feedback.FeedbackRequest) (*feedback.FeedbackResponse, error) {
	err := s.DB.Where("ID = ?", in.ID).Delete(Feedback{}).GetErrors()
	if len(err) == 0 {
		return &feedback.FeedbackResponse{ID: in.ID}, nil
	} else {
		return &feedback.FeedbackResponse{ID: in.ID}, errors.New("internal found")
	}

}

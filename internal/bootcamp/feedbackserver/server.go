package feedbackserver

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/status"
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

	var check Feedback
	notFound := s.DB.Where("booking_code = ?", in.BookingCode).Find(&check).RecordNotFound()

	if notFound {
		err := s.DB.Create(&res).Error
		if err == nil {
			return &feedback.FeedbackResponse{ID: res.ID}, nil
		} else {
			log.Print(err)
			return &feedback.FeedbackResponse{ID: res.ID}, errors.New("cannot create feedback")
		}
	} else {
		return &feedback.FeedbackResponse{ID: res.ID}, status.Error(409, "conflict")
	}

}

func (s *server) GetByPassengerId(ctx context.Context, in *feedback.PassengerRequest) (*feedback.PassengerResponse, error) {
	var resArr []Feedback
	err := s.DB.Where("passenger_id = ?", in.PassengerID).Find(&resArr).GetErrors()
	if len(err) == 0 {
		var fbs [] *feedback.PassengerFeedback
		for _, res := range resArr {
			pfb := feedback.PassengerFeedback{
				ID:          res.ID,
				PassengerID: res.PassengerID,
				BookingCode: res.BookingCode,
				Feedback:    res.Feedback,
			}

			fbs = append(fbs, &pfb)
		}
		return &feedback.PassengerResponse{PassengerID: in.PassengerID, Feedback: fbs}, nil

	} else {
		log.Print(err)
		return &feedback.PassengerResponse{PassengerID: in.PassengerID}, status.Error(404, "not found")
	}
}
func (s *server) GetByBookingCode(ctx context.Context, in *feedback.BookingRequest) (*feedback.PassengerFeedback, error) {
	var res Feedback
	found := s.DB.Where("booking_code = ?", in.Code).Find(&res).RecordNotFound()
	if found {
		fb := feedback.PassengerFeedback{
			ID:          res.ID,
			PassengerID: res.PassengerID,
			BookingCode: res.BookingCode,
			Feedback:    res.Feedback,
		}
		return &fb, nil

	} else {
		return &feedback.PassengerFeedback{BookingCode: in.Code}, status.Error(404, "not found")
	}

}
func (s *server) Delete(ctx context.Context, in *feedback.FeedbackRequest) (*feedback.FeedbackResponse, error) {

	if err := s.DB.Where("ID = ?", in.ID).First(&Feedback{}).Error; gorm.IsRecordNotFoundError(err) {
		return &feedback.FeedbackResponse{ID: in.ID}, status.Error(404, "not found")
	} else {
		err := s.DB.Where("ID = ?", in.ID).Delete(Feedback{}).GetErrors()
		if len(err) == 0 {
			return &feedback.FeedbackResponse{ID: in.ID}, nil
		} else {
			log.Print(err)
			return &feedback.FeedbackResponse{ID: in.ID}, status.Error(500, "internal server error")
		}
	}
}

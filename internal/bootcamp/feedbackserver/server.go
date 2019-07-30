package feedbackserver

import (
	"context"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *feedback.PassengerFeedback) (*feedback.FeedbackResponse, error) {

	return &feedback.FeedbackResponse{}, nil
}

func (s *server) GetById(ctx context.Context, in *feedback.FeedbackRequest) (*feedback.PassengerFeedback, error) {

	return &feedback.PassengerFeedback{}, nil
}
func (s *server) GetByBookingCode(ctx context.Context, in *feedback.BookingRequest) (*feedback.ListFeedbackResponse, error) {
	return &feedback.ListFeedbackResponse{}, nil
}
func (s *server) Delete(ctx context.Context, in *feedback.FeedbackRequest) (*feedback.FeedbackResponse, error) {

	return &feedback.FeedbackResponse{}, nil
}

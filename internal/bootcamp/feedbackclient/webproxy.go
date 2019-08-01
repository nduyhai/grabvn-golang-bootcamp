package feedbackclient

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
	"log"
	"time"
)

type WebProxy struct {
	client *feedback.FeedbackServiceClient
}

func (w *WebProxy) addFeedback(ctx *gin.Context) {

	var argument feedback.CreateFeedbackRequest

	err := ctx.BindJSON(&argument)
	if err != nil {
		ctx.String(400, "invalid param")
		return
	} else {
		c, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := (*w.client).Add(c, &argument)
		if err == nil {
			ctx.JSON(201, res)
		} else {
			log.Print(err)
			convert := status.Convert(err)
			ctx.String(int(convert.Code()), convert.Message())
		}
	}

}

func (w *WebProxy) getFeedbackById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(400, "invalid param")
		return
	} else {
		c, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := (*w.client).GetById(c, &feedback.FeedbackRequest{ID: id})
		if err == nil {
			ctx.JSON(200, res)
		} else {
			log.Print(err)
			convert := status.Convert(err)
			ctx.String(int(convert.Code()), convert.Message())
		}
	}
}

func (w *WebProxy) deleteFeedbackById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(400, "invalid param")
		return
	} else {
		c, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := (*w.client).Delete(c, &feedback.FeedbackRequest{ID: id})
		if err == nil {
			ctx.JSON(200, res)
		} else {
			log.Print(err)
			convert := status.Convert(err)
			ctx.String(int(convert.Code()), convert.Message())
		}
	}
}

func (w *WebProxy) getFeedbackByCode(ctx *gin.Context) {
	code := ctx.Param("id")
	if code == "" {
		ctx.String(400, "invalid param")
		return
	} else {
		c, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := (*w.client).GetByBookingCode(c, &feedback.BookingRequest{Code: code})
		if err == nil {
			ctx.JSON(200, res)
		} else {
			log.Print(err)
			convert := status.Convert(err)
			ctx.String(int(convert.Code()), convert.Message())
		}
	}
}

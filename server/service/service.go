package service

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

type Service struct {
	orderCollection *mongo.Collection
}

func New(orderCollection *mongo.Collection) *Service {
	return &Service{
		orderCollection: orderCollection,
	}
}

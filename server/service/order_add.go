package service

import (
	"context"
	"net/http"
	"time"

	"github.com/alexadastra/hisho/server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// add an order
func (s *Service) AddOrder(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(c.Request.Context(), 100*time.Second)
	defer cancel()

	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(order)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	order.ID = primitive.NewObjectID()

	result, insertErr := s.orderCollection.InsertOne(ctx, order)
	if insertErr != nil {
		msg := "order item was not created"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result)
}

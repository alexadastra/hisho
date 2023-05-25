package service

import (
	"context"
	"net/http"
	"time"

	"github.com/alexadastra/hisho/server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// update a waiter's name for an order
func (s *Service) UpdateWaiter(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	orderID := c.Params.ByName("id")

	docID, _ := primitive.ObjectIDFromHex(orderID)

	var waiter models.Waiter
	if err := c.BindJSON(&waiter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := s.orderCollection.UpdateOne(ctx, bson.M{"_id": docID},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   "server",
						Value: waiter.Server,
					},
				},
			},
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result.ModifiedCount)
}

package internal

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/helpers"
	"github.com/hanshal101/ppSM/models"
)

func UserbyID(c *gin.Context) {
	id := c.Param("user_id")
	user_id, err := helpers.STRtoUINT(id)
	if err != nil {
		log.Fatalf(" ", err)
	}
	var mess string
	var user models.User
	user, mess = helpers.FindUserbyID(user_id)
	if mess != "" {
		c.JSON(400, gin.H{"error": mess})
		return
	}

	c.JSON(200, user)
}

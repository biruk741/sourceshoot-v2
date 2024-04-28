package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleGetIndustries handles the GET request to retrieve all industries.
func (h Handler) HandleGetIndustries(c *gin.Context) {
	allIndustries, err := h.IndustryService.GetAllIndustries()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, allIndustries)
}

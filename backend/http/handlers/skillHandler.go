package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) HandleGetSkills(c *gin.Context) {
	allSkills, err := h.SkillsService.GetSkills()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, allSkills)
}

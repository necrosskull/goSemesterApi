package controllers

import (
	"github.com/gin-gonic/gin"
	"goSemesterApi/internal/logic"
	"goSemesterApi/internal/model"
	"strconv"
	"time"
)

// GetWeek godoc
// @Summary Get current week number
// @Description Get current week number
// @Tags Semester
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Week
// @Failure 404 {object} model.ErrorResponse
// @Router /week [get]
func GetWeek(c *gin.Context) {
	week := logic.GetCurrentWeekNumber()

	if week.Week < 0 {
		message := "Семестр еще не начался, до начала семестра осталось недель: " + strconv.Itoa(-week.Week)
		response := model.ErrorResponse{Detail: message}

		c.JSON(404, response)
		return
	}
	c.JSON(200, week)

}

// GetSemester godoc
// @Summary Get current semester info
// @Description Get current semester info
// @Tags Semester
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Period
// @Router /semester [get]
func GetSemester(c *gin.Context) {
	period := logic.GetPeriod(time.Now())
	c.JSON(200, period)
}

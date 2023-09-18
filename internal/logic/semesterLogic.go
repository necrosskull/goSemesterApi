package logic

import (
	"goSemesterApi/internal/model"
	"time"
)

func GetPeriod(date time.Time) model.Period {
	var yearStart, yearEnd, semester int

	if date.Month() >= 9 {
		yearStart = date.Year()
		yearEnd = date.Year() + 1
		semester = 1
	} else {
		yearStart = date.Year() - 1
		yearEnd = date.Year()
		semester = 2
	}

	startDate := getSemesterStartDate(yearStart, semester)

	return model.Period{
		YearStart: yearStart,
		YearEnd:   yearEnd,
		Semester:  semester,
		StartDate: startDate,
	}
}

func getSemesterStartDate(yearStart, semester int) time.Time {
	var startDay, startMonth int

	if semester == 1 {
		startDay = 1
		startMonth = 9
	} else {
		startDay = 1
		startMonth = 2
	}

	startDate := time.Date(yearStart, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)

	if startDate.Weekday() == time.Sunday {
		startDate = startDate.AddDate(0, 0, 1)
	}

	return startDate
}

func GetCurrentWeekNumber() model.Week {
	currentDate := time.Now()
	period := GetPeriod(currentDate)
	semesterStartDate := period.StartDate

	if currentDate.Before(semesterStartDate) {
		return model.Week{
			Week: 1,
		}
	}

	_, currentWeek := currentDate.ISOWeek()
	_, semesterWeek := semesterStartDate.ISOWeek()

	week := currentWeek - semesterWeek + 1

	return model.Week{
		Week: week,
	}
}

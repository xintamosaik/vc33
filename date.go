package main

import (
	"errors"
	"strconv"
	"time"
)

type Month int
type Year int
type MonthDate struct {
	Month Month `json:"month"`
	Year  Year  `json:"year"`
}

const maxCareerHistoryYears = 40

func earliestJobYear() int {
	return thisYear() - maxCareerHistoryYears
}
func earliestEducationYear() int {
	return thisYear() - maxCareerHistoryYears
}

func (d MonthDate) Before(other MonthDate) bool {
	if d.Year != other.Year {
		return d.Year < other.Year
	}

	return d.Month < other.Month
}
func parseMonth(value string) (Month, error) {
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("month must be a number")
	}

	if n < 1 || n > 12 {
		return 0, errors.New("month must be between 1 and 12")
	}

	return Month(n), nil
}

func parseYear(value string) (Year, error) {
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("year must be a number")
	}

	if n < earliestJobYear() || n > thisYear() {
		return 0, errors.New("invalid year")
	}

	return Year(n), nil
}

func parseMonthDate(monthValue, yearValue string) (MonthDate, error) {
	month, err := parseMonth(monthValue)
	if err != nil {
		return MonthDate{}, err
	}

	year, err := parseYear(yearValue)
	if err != nil {
		return MonthDate{}, err
	}

	return MonthDate{
		Month: month,
		Year:  year,
	}, nil
}
func thisYear() int {
	return time.Now().Year()
}
func valueOrEmptyString(number int) string {
	if number == 0 {
		return ""
	}
	return strconv.Itoa(number)
}

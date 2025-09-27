package services

import "time"

var monthSeasonMap = map[int]string{
	1: "winter",
	2: "winter",
	3: "spring",
	4: "spring",
	5: "spring",
	6: "summer",
	7: "summer",
	8: "summer",
	9: "fall",
	10: "fall",
	11: "fall",
	12: "winter",
}

func findCurrentSeason() string {
	month := time.Now().Month()

	return monthSeasonMap[int(month)]
}

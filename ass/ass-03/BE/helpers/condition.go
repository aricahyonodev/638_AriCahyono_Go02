package helpers

func CheckWaterStatus(water int) string  {
	var waterStatus = "aman"
	if water >=  6 && water <= 8 {
		waterStatus = "siaga"
	}else if water > 8 {
		waterStatus = "bahaya"
	}

	return waterStatus
}

func CheckWindStatus(wind int) string  {
	var waterStatus = "aman"
	if wind >=  7 && wind <= 15 {
		waterStatus = "siaga"
	}else if wind > 15 {
		waterStatus = "bahaya"
	}

	return waterStatus
}
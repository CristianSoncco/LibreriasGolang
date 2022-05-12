//Funcion que recibe formato de hora 20:20:20 y devuelve 08:20 p.m.
func FormatHour24(phour string) (hourpm string) {
	hourReturn := ""
	ampm := "am"
	if phour == "" {
		phour = "00:00:00"
	}
	firstBin := strings.Index(phour, ":")
	i, err := strconv.Atoi(phour[0:firstBin])
	hour := i
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	j, err_ := strconv.Atoi(phour[firstBin+1 : firstBin+3])
	minute := j
	if err_ != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if hour < 12 {
		ampm = "am"
	} else {
		ampm = "pm"
	}

	hour = hour % 12
	if hour == 0 {
		hour = 12
	}

	hourReturn = fmt.Sprintf("%02d", hour) + ":" + fmt.Sprintf("%02d", minute) + " " + ampm
	return hourReturn
}

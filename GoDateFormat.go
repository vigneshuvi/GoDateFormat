package GoDateFormat

import (
         "time"
         "fmt"
         "strings"
)

const yyyy = "2006"
const yy = "06"

const mm = "01"
const mmm = "Jan"
const mmmm = "January"

const dd = "02"
const ddd = "Mon"
const dddd = "Monday"

const HH = "15"
const MM = "04"
const SS = "05"
const ss = "05"

const tt = "tt"

const Z = "MST"
const ZZZ = "MST"

func ConvertFormat(format string) (string){
	var goFormate = format;
	if strings.Contains(goFormate, "yyyy") {
		goFormate = strings.Replace(goFormate, "yyyy", yyyy, -1)
	} else if strings.Contains(goFormate, "yy") {
		goFormate = strings.Replace(goFormate, "yy", yy, -1)
	}

	if strings.Contains(goFormate, "mmmm") {
		goFormate = strings.Replace(goFormate, "mmmm", mmmm, -1)
	} else if strings.Contains(goFormate, "mmm") {
		goFormate = strings.Replace(goFormate, "mmm", mmm, -1)
	} else if strings.Contains(goFormate, "mm") {
		goFormate = strings.Replace(goFormate, "mm", mm, -1)
	}

	if strings.Contains(goFormate, "dddd") {
		goFormate = strings.Replace(goFormate, "dddd", dddd, -1)
	} else if strings.Contains(goFormate, "ddd") {
		goFormate = strings.Replace(goFormate, "ddd", ddd, -1)
	} else if strings.Contains(goFormate, "dd") {
		goFormate = strings.Replace(goFormate, "dd", dd, -1)
	}

	if strings.Contains(goFormate, "HH") {
		goFormate = strings.Replace(goFormate, "HH", HH, -1)
	}
	if strings.Contains(goFormate, "MM"){
		goFormate = strings.Replace(goFormate, "MM", MM, -1)
	}
	if strings.Contains(goFormate, "SS"){
		goFormate = strings.Replace(goFormate, "SS", SS, -1)
	} else if strings.Contains(goFormate, "ss"){
		goFormate = strings.Replace(goFormate, "ss", ss, -1)
	}

	if strings.Contains(goFormate, "ZZZ"){
		goFormate = strings.Replace(goFormate, "ZZZ", ZZZ, -1)
	} else if strings.Contains(goFormate, "Z"){
		goFormate = strings.Replace(goFormate, "Z", Z, -1)
	}
	return (goFormate)
}

func GetToday(format string) (todayString string){
	
	today := time.Now()
	var goFormate = ConvertFormat(format)

	fmt.Println("Old Format: %s", format)
	fmt.Println("New Format: %s", goFormate)
	todayString = today.Format(goFormate);
	fmt.Println(todayString)
	return
}


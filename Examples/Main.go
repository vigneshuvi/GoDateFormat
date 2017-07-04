package main


import (
    "fmt"
    "time"
    "github.com/vigneshuvi/GoDateFormat"
)

func main() {
    var goFormat = GoDateFormat.ConvertFormat("yyyy-MM-dd HH:mm:ss Z")
    fmt.Println("Go Long Date Format(Today): ", GetToday(goFormat))
}

func GetToday(format string) (todayString string){
    today := time.Now()
    todayString = today.Format(format);
    return
}



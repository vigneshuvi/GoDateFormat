package main


import (
    "fmt"
    "time"
    "github.com/vigneshuvi/GoDateFormat"
)

func main() {
    fmt.Println("Go Date Format(Today - 'yyyy-MM-dd HH:mm:ss Z'): ", GetToday(GoDateFormat.ConvertFormat("yyyy-MM-dd HH:mm:ss Z")))
    fmt.Println("Go Date Format(Today - 'yyyy-MMM-dd'): ", GetToday(GoDateFormat.ConvertFormat("yyyy-MMM-dd")))
    fmt.Println("Go Time Format(NOW - 'HH:mm:ss'): ", GetToday(GoDateFormat.ConvertFormat("HH:MM:SS")))
}

func GetToday(format string) (todayString string){
    today := time.Now()
    todayString = today.Format(format);
    return
}



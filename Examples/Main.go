package main


import (
    "fmt"
    "github.com/vigneshuvi/GoDateFormat"
)

func main() {
    var goFormat = GoDateFormat.ConvertFormat("yyyy-MM-dd HH:mm:ss Z")
    fmt.Println("Go Long Date Formate: %s", goFormat)
}



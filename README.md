[![GoDateFormat](https://img.shields.io/travis/rust-lang/rust.svg)](https://github.com/vigneshuvi/GoDateFormat)
[![Language Go](https://img.shields.io/badge/Go-Lang-green.svg?style=shields)](https://golang.org/)



# GoDateFormat

Convert normal date format into Go lang date format

## Installation

First thing is to get your GoDateFormat package into your machine.

```go

go get "github.com/vigneshuvi/GoDateFormat"

```

## Importing packages

Import all necessary packages.("fmt" - Print, "time" - Getting time from machine) 

```go

import (
    "fmt"
    "time"
    "github.com/vigneshuvi/GoDateFormat"
)

```


### Examples:

```go

func main() {
    var goFormat = GoDateFormat.ConvertFormat("yyyy-MM-dd HH:mm:ss Z")
    fmt.Println("Go Long Date Format(Today): ", GetToday(goFormat))
}

func GetToday(format string) (todayString string){
    today := time.Now()
    todayString = today.Format(format);
    return
}


```

### Write Output:

```go

Output: 

Go Long Date Format(Today): 2017-09-04 17:07:28 IST

```

## License

GoDateFormat is licensed under the Apache License.

## Contact

### Vignesh Kumar
* http://vigneshuvi.github.io
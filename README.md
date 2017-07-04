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
    fmt.Println("Go Date Format(Today - 'yyyy-MM-dd HH:mm:ss Z'): ", GetToday(GoDateFormat.ConvertFormat("yyyy-MM-dd HH:mm:ss Z")))
    fmt.Println("Go Date Format(Today - 'yyyy-MMM-dd'): ", GetToday(GoDateFormat.ConvertFormat("yyyy-MMM-dd")))
    fmt.Println("Go Time Format(NOW - 'HH:mm:ss'): ", GetToday(GoDateFormat.ConvertFormat("HH:mm:ss")))
}

func GetToday(format string) (todayString string){
    today := time.Now()
    todayString = today.Format(format);
    return
}


```

### Output:

```go

Output: 

Go Date Format(Today - 'yyyy-MM-dd HH:mm:ss Z'):  2017-33-04 17:07:10 IST
Go Date Format(Today - 'yyyy-MMM-dd'):  2017-33M-04
Go Time Format(NOW - 'HH:mm:ss'):  17:07:10

```

## License

GoDateFormat is licensed under the Apache License.

## Contact

### Vignesh Kumar
* http://vigneshuvi.github.io
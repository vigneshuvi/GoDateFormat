[![GoDateFormat](https://img.shields.io/travis/rust-lang/rust.svg)](https://github.com/vigneshuvi/GoDateFormat)
[![Language Go](https://img.shields.io/badge/Language-Go-orange.svg?style=shields)](https://golang.org/)



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

### Constants and Placeholders

```go

const  (
	yyyy = "2006"
	yy = "06"
	mmmm = "January"
	mmm = "Jan"
	mm = "01"
	dddd = "Monday"
	ddd = "Mon"
	dd = "02"

	HHT = "03"
	HH = "15"
	MM = "04"
	SS = "05"
	ss = "05"
	tt = "PM"
	Z = "MST"
	ZZZ = "MST"
)

```

### Importing packages

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
    fmt.Println("Go Time Format(NOW - 'HH:MM:SS'): ", GetToday(GoDateFormat.ConvertFormat("HH:MM:SS")))
    fmt.Println("Go Time Format(NOW - 'HH:MM:SS tt'): ", GetToday(GoDateFormat.ConvertFormat("HH:MM:SS tt")))
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

Go Date Format(Today - 'yyyy-MM-dd HH:mm:ss Z'):  2017-19-04 18:07:05 IST
Go Date Format(Today - 'yyyy-MMM-dd'):  2017-Jul-04
Go Time Format(NOW - 'HH:MM:SS'):  18:19:05
Go Time Format(NOW - 'HH:MM:SS tt'):  06:19:05 PM

```

## License

GoDateFormat is licensed under the Apache License.

## Contact

### Vignesh Kumar
* http://vigneshuvi.github.io
[![Language Go](https://img.shields.io/badge/Go-Lang-green.svg?style=shields)](https://golang.org/)



# GoDateFormat
Convert normal date format into Go lang date format


## Importing packages

First thing is to import the package. 

```go

import ( 
	"github.com/vigneshuvi/GoDateFormat"
	"fmt"
	)

```


## Examples:

```go

import "github.com/vigneshuvi/GoDateFormat"

var goFormate = GoDateFormat.ConvertFormat("yyyy-MM-dd HH:mm:ss Z")
fmt.Println("Go Long Date Formate: %s", goFormate)


```

### Write Output:

```go

Output: 

Go Long Date Formate: 2017-58-04 16:07:31 IST

```

## License

SwiftCSVExport is licensed under the Apache License.

## Contact

### Vignesh Kumar
* http://vigneshuvi.github.io
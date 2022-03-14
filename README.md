# A simple torrent download manager in golang

## Installation
To install the package, use the command below

```sh
go get github.com/milinches/download-manager
```

Import the package

```go
import "github.com/milinches/download-manager"
```

`Download a torrent`

```go
package main

import (
    "log"
    "fmt"
    "github.com/milinches/download-manager"
)

torrent, err := donwload.Download("{torrentLink}", "{fileName}", {section})
// handle error
if err != nil {
    log.Fatal(err.Error())
}
fmt.Println(torrent)
```

Used [Muhammed Usman](https://www.youtube.com/c/MuhammadUsmanH) tutorial as my guide. Please, a sub to his youtube channel would be amazing! 🥰
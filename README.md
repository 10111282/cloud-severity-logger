# JSON logs with severity
This module provides a set of verbose functions to output logs formatted to the JSON with a severity info appended.
Such formatting allows logs to be correctly recorded in StackDriver when deployed to a cloud infrastructure, e.g. to Google CloudRun.

## Installation
```
$ go get github.com/10111282/cloud-severity-logger
```

## Usage
```
import (
	logger "github.com/10111282/cloud-severity-logger"
)

logger.Error("It's an error!")
\\ {"message":"It's an error!","severity":"ERROR"}

logger.Warning("It's an warning!")
\\ {"message":"It's a warning!","severity":"WARNING"}
```



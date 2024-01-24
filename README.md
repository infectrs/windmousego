# windmousego
Implementation of WindMouse in Golang

## Installation

`go get github.com/infectrs/windmousego`
## Usage

```golang
package main

import (
	"fmt"
	"github.com/infectrs/windmousego"
	"math"
	"math/rand"
)

func main() {
	windmouse := windmousego.MouseSettings{
		StartX:     math.Ceil(rand.Float64() * 1920),
		StartY:     math.Ceil(rand.Float64() * 1080),
		EndX:       math.Ceil(rand.Float64() * 1920),
		EndY:       math.Ceil(rand.Float64() * 1080),
		Gravity:    math.Ceil(rand.Float64() * 10),
		Wind:       math.Ceil(rand.Float64() * 10),
		MinWait:    2,
		MaxWait:    math.Ceil(rand.Float64() * 5),
		MaxStep:    math.Ceil(rand.Float64() * 3),
		TargetArea: math.Ceil(rand.Float64() * 10),
	}

	result := windmouse.GeneratePoints()

	fmt.Println(result)
}
```
## Credits

https://github.com/arevi/wind-mouse

package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func Message(text interface{}) {
	color.Blue(fmt.Sprint(text))
}

func Warn(text interface{}) {
	color.Yellow(fmt.Sprint(text))
}

func Error(text interface{}) {
	color.Red(fmt.Sprint(text))
}

func Success(text interface{}) {
	color.Green(fmt.Sprint(text))
}

package utils

import (
	"github.com/fatih/color"
)

func Message(text string) {
	color.Blue(text)
}

func Warn(text string) {
	color.Yellow(text)
}

func Error(text string) {
	color.Red(text)
}

func Success(text string) {
	color.Green(text)
}

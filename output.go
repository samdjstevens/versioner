package main

import (
	"github.com/fatih/color"
)

// helper functions for printing coloured text lines

func YellowLn(text string) {
	color.New(color.FgYellow).Println(text)
}

func BlueLn(text string) {
	color.New(color.FgBlue).Println(text)
}

func RedLn(text string) {
	color.New(color.FgRed).Println(text)
}

func GreenLn(text string) {
	color.New(color.FgGreen).Println(text)
}
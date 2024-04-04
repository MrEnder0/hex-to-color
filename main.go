package main

import (
	"fmt"
	"math"
)

type Color struct {
	red, green, blue int
}

type ColorDistance struct {
	distance int
	color    Color
}

var Colors = map[Color]string{
	{255, 0, 0}:     "red",
	{0, 255, 0}:     "green",
	{0, 0, 255}:     "blue",
	{255, 255, 0}:   "yellow",
	{0, 255, 255}:   "cyan",
	{255, 0, 255}:   "magenta",
	{255, 130, 0}:   "orange",
	{255, 0, 210}:   "pink",
	{128, 128, 128}: "gray",
	{255, 255, 255}: "white",
	{0, 0, 0}:       "black",
}

func main() {
	for {
		var given_hex string

		fmt.Println("Enter a color in hex format (#RRGGBB):")
		fmt.Scan(&given_hex)

		var converted_hex = hex_to_rgb(given_hex)

		smallest_distance := ColorDistance{distance: math.MaxInt64}

		for k := range Colors {
			if converted_hex == k {
				fmt.Printf("The color is %s\n", Colors[k])
				return
			} else {
				distance := math.Abs(float64(converted_hex.red-k.red)) + math.Abs(float64(converted_hex.green-k.green)) + math.Abs(float64(converted_hex.blue-k.blue))

				if int(distance) < smallest_distance.distance {
					smallest_distance.distance = int(distance)
					smallest_distance.color = k
				}
			}
		}

		fmt.Printf("The closest color is %s\n", Colors[smallest_distance.color])
	}
}

func hex_to_rgb(hex string) Color {
	var r, g, b int
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	return Color{r, g, b}
}

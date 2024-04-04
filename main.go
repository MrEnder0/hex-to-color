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
	var r, g, b int
	for {
		fmt.Println("Enter an r value:")
		fmt.Scan(&r)
		fmt.Println("Enter a g value:")
		fmt.Scan(&g)
		fmt.Println("Enter a b value:")
		fmt.Scan(&b)

		given_color := Color{r, g, b}
		smallest_distance := ColorDistance{distance: math.MaxInt64}

		for k := range Colors {
			if given_color == k {
				fmt.Printf("The color is %s\n", Colors[k])
				return
			} else {
				distance := math.Abs(float64(r-k.red)) + math.Abs(float64(g-k.green)) + math.Abs(float64(b-k.blue))

				if int(distance) < smallest_distance.distance {
					smallest_distance.distance = int(distance)
					smallest_distance.color = k
				}
			}
		}

		fmt.Printf("The closest color is %s\n", Colors[smallest_distance.color])
	}
}

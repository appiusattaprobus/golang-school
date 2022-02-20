package main

import (
	"fmt"

	square "./homework-2"
	"github.com/kyokomi/emoji"
)

// http://www.unicode.org/emoji/charts/emoji-list.html
func main() {
	fmt.Println(emoji.Sprint("Square of :red_circle: = "), square.CalcSquare(10.0, square.SidesCircle))
	fmt.Println(emoji.Sprint("Square of :red_triangle_pointed_up: = "), square.CalcSquare(10.0, square.SidesTriangle))
	fmt.Println(emoji.Sprint("Square of :red_square: = "), square.CalcSquare(10.0, square.SidesSquare))
}

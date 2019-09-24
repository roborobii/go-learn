package main

import (
	"fmt"
	"math"
	
	// import a package that you've created
	"github.com/roborobii/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.8))
	fmt.Println(math.Ceil(2.8))
	fmt.Println(math.Sqrt(2.8))

	fmt.Println(strutil.Reverse("Hello!"))
}
package main
import "fmt"
import "math"
func quad() {
    var a float64
    var b float64
    var c float64
    var item1 float64
    var item2 float64
    fmt.Print("What is a?: ")
    _,err := fmt.Scan(&a)
    fmt.Print("What is b?: ")
    _,err2 := fmt.Scan(&b)
    fmt.Print("What is c?: ")
    _,err3 := fmt.Scan(&c)
    item1 = ((-b)+math.Sqrt(math.Pow(b, 2)-4*a*c))/2*a
    item2 = ((-b)-math.Sqrt(math.Pow(b, 2)-4*a*c))/2*a
    _=err
    _=err2
    _=err3
    fmt.Println("X =",item1,"and X =",item2)
}
func algeb() {
	var input string
	fmt.Print("Welcome to the algebra calculator available calculators include: quadratic equasion(quad) what calc do you want?: ")
	_,err := fmt.Scan(&input)
	_=err
	if input == "quad" {
		quad()
	}
}
func geom() {
	fmt.Println("wip")
}
func chem() {
	fmt.Println("wip")
}
func phys() {
	fmt.Println("wip")
}
func misc() {
	fmt.Println("wip")
}
func main() {
	var input string
	fmt.Print("Welcome to Gocalc available calculators include algebra(algeb), geometry(geom), chemistry(chem), misc(misc), physics(phys)\nWhat calculator do you want?: ")
	_,err := fmt.Scan(&input)
	_=err
	if input == "chem" {
		chem()
	} else if input == "algeb" {
		algeb()
	} else if input == "geom" {
		geom()
	} else if input == "misc" {
		misc()
	} else if input == "phys" {
		phys()
	} else {
		fmt.Println("\n\nERROR! Something went wrong please try again!\n\n")
		main()
	}
}
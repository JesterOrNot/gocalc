package main
import "fmt"
import "math"
import "github.com/Arafatk/glot"
func exGraph() {
	dimensions := 2
	// The dimensions supported by the plot
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	pointGroupName := "Simple Circles"
	style := "circle"
	points := [][]float64{{7, 3, 13, 5.6, 11.1}, {12, 13, 11, 1,  7}}
        // Adding a point group
	plot.AddPointGroup(pointGroupName, style, points)
	// A plot type used to make points/ curves and customize and save them as an image.
	plot.SetTitle("Example Plot")
	// Optional: Setting the title of the plot
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("Y-Axis")
	// Optional: Setting label for X and Y axis
	plot.SetXrange(-2, 18)
	plot.SetYrange(-2, 18)
	// Optional: Setting axis ranges
	plot.SavePlot("2.png")
}
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
	exGraph()
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
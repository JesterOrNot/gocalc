package main

import (
	"fmt"
)

func velocity() {
	var displacement float64
	var time_interval float64
	var velocity float64
	fmt.Print("What is the displacment?: ")
	err, _ := fmt.Scan(&displacement)
	fmt.Print("What is the time interval?: ")
	err2, _ := fmt.Scan(&time_interval)
	_ = err
	_ = err2
	velocity = displacement / time_interval
	fmt.Println("Velocity =", velocity)
}
func acceleration() {
	var initial_velocity float64
	var final_velocity float64
	var initial_time float64
	var final_time float64
	var delta_t float64
	var acceleration float64
	var delta_v float64
	fmt.Print("What is the initial time?: ")
	err, _ := fmt.Scan(&initial_time)
	fmt.Print("What is the final time?: ")
	err1, _ := fmt.Scan(&final_time)
	fmt.Print("What is the initial velocity?: ")
	err2, _ := fmt.Scan(&initial_velocity)
	fmt.Print("What is the final velocity?: ")
	err3, _ := fmt.Scan(&final_velocity)
	delta_t = final_time - initial_time
	delta_v = final_velocity - initial_velocity
	acceleration = delta_v / delta_t
	fmt.Println("acceleration =", acceleration)
	_ = err
	_ = err1
	_ = err2
	_ = err3
}
func linOfSym() {
	var a float64
	var b float64
	var los float64
	fmt.Print("What is a?: ")
	fmt.Scan(&a)
	fmt.Print("What is b?: ")
	fmt.Scan(&b)
	los = (-1 * b) / (2 * a)
	fmt.Println("Line of symmetry =", los)
}
func algeb() {
	var input string
	fmt.Print("Welcome to the algebra calculator available calculators include: quadratic equation(quad),Line of symmetry(los) what calc do you want?: ")
	_, err := fmt.Scan(&input)
	_ = err
	if input == "quad" {
		quad()
	} else if input == "los" {
		linOfSym()
	}
}
func heatEqu() {
	var solving_for string
	var q float64
	var m float64
	var c float64
	var dt float64
	var fi float64
	var ini float64
	fmt.Print("What are we solving for?(q,m,c,t): ")
	fmt.Scan(&solving_for)
	if solving_for == "q" {
		fmt.Print("What is m?: ")
		fmt.Scan(&m)
		fmt.Print("What is the final?: ")
		fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		fmt.Scan(&ini)
		fmt.Print("What is the specific heat?: ")
		fmt.Scan(&c)
		q = m * c * (fi - ini)
		fmt.Println("q =", q)
	} else if solving_for == "m" {
		fmt.Print("What is q?: ")
		fmt.Scan(&q)
		fmt.Print("What is the final?: ")
		fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		fmt.Scan(&ini)
		fmt.Print("What is the specific heat?: ")
		fmt.Scan(&c)
		m = q / (c * (fi - ini))
		fmt.Println("m =", m)
	} else if solving_for == "c" {
		fmt.Print("What is q?: ")
		fmt.Scan(&q)
		fmt.Print("What is the final?: ")
		fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		fmt.Scan(&ini)
		fmt.Print("What is m?: ")
		fmt.Scan(&m)
		c = q / ((fi - ini) * m)
		fmt.Println("c =", c)
	} else if solving_for == "t" {
		fmt.Print("What is q?: ")
		fmt.Scan(&q)
		fmt.Print("What is the final?: ")
		fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		fmt.Scan(&ini)
		fmt.Print("What is m?: ")
		fmt.Scan(&m)
		dt = q / (c * m)
		fmt.Println("ΔT =", dt)
	}
}
func areaCalc() {
	var input string
	fmt.Print("Welcome to the area finder the shapes that this can find the area for so far are trapazoid(trap),circle(circ)\nWhat do you want?: ")
	fmt.Scan(&input)
	if input == "trap" {
		AOTrap()
	} else if input == "circ" {
		AOCirc()
    }
}
func geom() {
	var input string
	fmt.Print("Welcome to the geometry calculator! available calculators include: pythagTheor(pythag),area calc(ac) \nwhat do you want?: ")
	fmt.Scan(&input)
	if input == "pythag" {
		pythagTheor()
	} else if input =="ac" {
		areaCalc()
	}
}
func chem() {
	var input string
	fmt.Print("Welcome to the chem calculator available calculators include: heat equation(heatEqu): ")
	_, err := fmt.Scan(&input)
	_ = err
	if input == "heatEqu" {
		heatEqu()
	}
}
func phys() {
	var input string
	fmt.Print("Welcome to the physics calculator available calculators include: velocity calculator(velocity),acceleration calculator(acc) what calc do you want?: ")
	_, err := fmt.Scan(&input)
	_ = err
	if input == "velocity" {
		velocity()
	} else if input == "acc" {
		acceleration()
	}
}
func misc() {
	fmt.Println("wip")
}
func main() {
	var input string
	fmt.Print("Welcome to Gocalc available calculators include algebra(algeb), geometry(geom), chemistry(chem), misc(misc), physics(phys)\nWhat calculator do you want?: ")
	_, err := fmt.Scan(&input)
	_ = err
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

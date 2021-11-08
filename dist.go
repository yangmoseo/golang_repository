package main

import (
	"distance/distalgs"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// create custom flag
type Coord []string

func (arr *Coord) String() string {
	return fmt.Sprintf("%v", *arr)
}

func (arr *Coord) Set(s string) error {
	*arr = strings.Split(s, ",")
	return nil

}

// paring string to float64
func selLatLon(inArray []string) ([]float64, error) {
	var outArray []float64
	var err error
	var val1 float64
	for _, value := range inArray {

		if len(value) != 0 {
			val1, err = strconv.ParseFloat(strings.Trim(value, " "), 64)
			if err != nil {
				fmt.Println(err)
				break
			}
			outArray = append(outArray, val1)
		}
	}
	return outArray, err
}

//calculate distance and Print
func printDis(origin []float64, dest []float64) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			flag.Usage()
		}
	}()

	//calc distance
	fmt.Printf("Distance : %v merter", distalgs.CalcDis(origin[1], origin[0], dest[1], dest[0]))
}

func main() {

	var coord1 Coord
	var coord2 Coord
	flag.Var(&coord1, "origin", "[required] input origin coord : longitude,latitude")
	flag.Var(&coord2, "dest", "[required] input destination coord : longitude,latitude")
	flag.Parse()

	fmt.Println("orogin:", coord1)
	fmt.Println("dest:", coord2)

	origin, err1 := selLatLon(coord1)
	dest, err2 := selLatLon(coord2)

	if err1 == nil && err2 == nil {
		printDis(origin, dest)
	} else {
		fmt.Println("Run time  error")
	}
}

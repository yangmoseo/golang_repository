package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//입력 : ./llmeter -origin 126.859706754,37.635076396 -dest 127.047839288,37.582547966
//출럭 : 200

func hsin(angle float64) float64 {
	return math.Pow(math.Sin(angle/2), 2)
}

func calcDis(originLat, originLon, destLat, destLon float64) float64 {
	r := 6378100.0
	originLatRad := originLat * math.Pi / 180
	originLonRad := originLon * math.Pi / 180
	destLatRad := destLat * math.Pi / 180
	destLonRad := destLon * math.Pi / 180
	h := hsin(destLatRad-originLatRad) + math.Cos(originLatRad)*math.Cos(destLatRad)*hsin(destLonRad-originLonRad)

	return 2 * r * math.Asin(math.Sqrt(h))
}

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

func main() {

	//get input args
	argsWithoutProg := os.Args[1:]

	// string join, replace -.,=
	argsString := strings.Join(argsWithoutProg, "_")
	customReplace := strings.NewReplacer("_.", ".", ",", "_", "=", "_")
	argsString = customReplace.Replace(argsString)

	//-origin, -dest split
	originIdx := strings.LastIndex(argsString, "-origin")
	destIdx := strings.LastIndex(argsString, "-dest")
	originString := (argsString[originIdx+len("-origin") : destIdx])[:]
	destString := argsString[destIdx+len("-dest"):]

	//lat, lon split
	originArray := strings.Split(originString, "_")
	destArray := strings.Split(destString, "_")

	//sel lat, lon
	origin, err1 := selLatLon(originArray)
	dest, err2 := selLatLon(destArray)

	if err1 == nil && err2 == nil {
		//calc distance
		fmt.Println(calcDis(origin[1], origin[0], dest[1], dest[0]))
	} else {
		fmt.Println("Run time  error")
	}
}

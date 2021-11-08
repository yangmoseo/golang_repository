package distalgs

import "math"

func hsin(angle float64) float64 {
	return math.Pow(math.Sin(angle/2), 2)
}

func CalcDis(originLat, originLon, destLat, destLon float64) float64 {
	r := 6378100.0
	originLatRad := originLat * math.Pi / 180
	originLonRad := originLon * math.Pi / 180
	destLatRad := destLat * math.Pi / 180
	destLonRad := destLon * math.Pi / 180
	h := hsin(destLatRad-originLatRad) + math.Cos(originLatRad)*math.Cos(destLatRad)*hsin(destLonRad-originLonRad)

	return 2 * r * math.Asin(math.Sqrt(h))

}

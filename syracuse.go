package main

import (
	"errors"
	"fmt"
)

func max(a, b int) int {
	_max := a
	if a < b {
		_max = b
	}
	return _max
}
func _syracuse(x, u0, _tempsVol, _tempsVolAltitude, _altitudeMax int, _tvaFound bool) (tempsVol, tempsVolAltitude, altitudeMax int) {
	fmt.Println(x, u0, _tempsVol, _tempsVolAltitude, _altitudeMax)
	if x == 1 {
		return _tempsVol, _tempsVolAltitude, _altitudeMax
	} else {
		if !_tvaFound && x <= u0 && _tempsVol > 0 {
			_tempsVolAltitude = max(0, _tempsVol-1)
			_tvaFound = true
		}
		_tempsVol += 1
		_altitudeMax = max(x, _altitudeMax)

		if x%2 == 0 {
			return _syracuse(x/2, u0, _tempsVol, _tempsVolAltitude, _altitudeMax, _tvaFound)
		} else {
			return _syracuse(3*x+1, u0, _tempsVol, _tempsVolAltitude, _altitudeMax, _tvaFound)
		}
	}
}

func syracuse(x int) (tempsVol, tempsVolAltitude, altitudeMax int, err error) {
	if x < 1 {
		return -1, -1, -1, errors.New("x doit etre strictement positif")
	}
	if x == 1 {
		return 0, 3, 4, nil
	}
	_tempsVol, _tempsVolAltitude, _altitudeMax := _syracuse(x, x, 0, 0, x, false)
	return _tempsVol, _tempsVolAltitude, _altitudeMax, nil
}

func main() {
	tempsVol, tempsVolAltitude, altitudeMax, err := syracuse(15)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tempsVol, tempsVolAltitude, altitudeMax)
	}
}

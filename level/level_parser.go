package level

import (
	"io/ioutil"
	"red_zenith/common"
	"red_zenith/entities"
	"strconv"
	"strings"
)

// GetLevel ...
func GetLevel(path string) (objects []entities.BaseEntity, spawnPoint common.Point) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return parseLevelString(string(data))
}

func parseLevelString(data string) (objects []entities.BaseEntity, spawnPoint common.Point) {
	objects = []entities.BaseEntity{}
	data = strings.ReplaceAll(data, "\r", "")
	lines := strings.Split(data, "\n")

	size := strings.Split(lines[1], ",")
	width, err := strconv.ParseFloat(size[0], 32)
	height, err := strconv.ParseFloat(size[1], 32)
	if err != nil {
		panic(err)
	}

	ceil := &entities.Ground{
		X:      0,
		Y:      float32(height),
		Width:  float32(width),
		Height: 50,
	}
	floor := &entities.Ground{
		X:      0,
		Y:      -50,
		Width:  float32(width),
		Height: 50,
	}
	left := &entities.Ground{
		X:      -50,
		Y:      -50,
		Width:  50,
		Height: float32(height) + 100,
	}
	right := &entities.Ground{
		X:      float32(width),
		Y:      -50,
		Width:  50,
		Height: float32(height) + 100,
	}
	objects = append(objects, ceil, floor, left, right)

	for i := 4; i < len(lines); i++ {
		e := parseEntity(lines[i])
		if e != nil {
			objects = append(objects, e)
		}
	}

	parts := strings.Split(lines[2], ",")
	x, err := strconv.ParseFloat(parts[0], 32)
	y, err := strconv.ParseFloat(parts[1], 32)
	handleError(err)
	spawnPoint = common.Point{
		X: float32(x),
		Y: float32(y),
	}

	return objects, spawnPoint
}

func parseEntity(str string) entities.BaseEntity {
	if str == "" {
		return nil
	}

	parts := strings.Split(str, ",")
	var result entities.BaseEntity = nil

	id, err := strconv.ParseInt(parts[0], 10, 32)
	x, err := strconv.ParseFloat(parts[1], 32)
	y, err := strconv.ParseFloat(parts[2], 32)
	handleError(err)

	switch id {
	case 0:
		width, err := strconv.ParseFloat(parts[3], 32)
		height, err := strconv.ParseFloat(parts[4], 32)
		handleError(err)
		result = &entities.Ground{
			X:      float32(x),
			Y:      float32(y),
			Width:  float32(width),
			Height: float32(height),
		}
	default:
		break
	}

	return result
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

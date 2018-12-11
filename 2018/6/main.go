package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/bwiggs/advent-of-code/2018/lib"
	"github.com/davecgh/go-spew/spew"
)

type Coord []int

func (c Coord) ID() string {
	return fmt.Sprintf("%d,%d", c[0], c[1])
}

func (c Coord) String() string {
	return c.ID()
}

func main() {
	coords, bounds := getCoords("input.txt")
	areas := make(map[string]int)
	infinite := make(map[string]bool)
	var closest Coord
	for y := bounds.yMin; y <= bounds.yMax; y++ {
		for x := bounds.xMin; x <= bounds.xMax; x++ {
			isInfinite := y == bounds.yMin || y == bounds.yMax || x == bounds.xMin || x == bounds.xMax

			closest = nil
			var minDist int
			var isEquallyBetween bool

			for _, c := range coords {
				dist := lib.RectlinearDist(x, y, c[0], c[1])
				if closest == nil || dist < minDist {
					closest = c
					minDist = dist
					isEquallyBetween = false
				} else if minDist == dist {
					isEquallyBetween = true
				}
			}

			if isInfinite {
				infinite[closest.ID()] = true
			} else if !isEquallyBetween {
				areas[closest.ID()]++
			}
		}
	}
	var largestAreaSize int
	for a, size := range areas {
		id := fmt.Sprintf("%d,%d", a[0], a[1])
		if _, infinite := infinite[id]; !infinite && size > largestAreaSize {
			largestAreaSize = size
		}
	}
	spew.Dump(largestAreaSize)
	spew.Dump(findRegionSize(coords, bounds, 10000))
}

func findRegionSize(coords []Coord, bounds Bounds, limit int) int {
	regionSize := 0
	m := image.NewRGBA(image.Rect(0, 0, bounds.xMax-bounds.xMin, bounds.yMax-bounds.yMin))
	for y := bounds.yMin; y <= bounds.yMax; y++ {
		for x := bounds.xMin; x <= bounds.xMax; x++ {
			totalDist := 0
			for _, c := range coords {
				totalDist += lib.RectlinearDist(x, y, c[0], c[1])
			}
			if totalDist < limit {
				regionSize++
				m.Set(x-bounds.xMin, y-bounds.yMin, color.RGBA{0, 0xff, 0, 0xff})
			} else {
				m.Set(x-bounds.xMin, y-bounds.yMin, color.RGBA{0, 0, 0, 0xff})
			}
		}
	}
	f, err := os.OpenFile("regions.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	png.Encode(f, m)
	return regionSize
}

type Bounds struct {
	xMin, xMax, yMin, yMax int
}

func getCoords(fn string) ([]Coord, Bounds) {
	coords := []Coord{}
	bounds := Bounds{}

	for i, l := range lib.ReadLines(fn) {
		xy := strings.Split(l, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		coords = append(coords, Coord{x, y})

		if i == 0 {
			bounds.xMin = x
			bounds.xMax = x
			bounds.yMin = y
			bounds.yMax = y
			continue
		}

		if x < bounds.xMin {
			bounds.xMin = x
		} else if x > bounds.xMax {
			bounds.xMax = x
		}

		if y < bounds.yMin {
			bounds.yMin = y
		} else if y > bounds.yMax {
			bounds.yMax = y
		}

	}
	return coords, bounds
}

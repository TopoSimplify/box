package box

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom"
)

func MBRToPolygon(o mbr.MBR) *geom.Polygon {
	var array = o.AsPolyArray()
	var coordinates = make([]geom.Point, 0, len(array))
	for i := range  array{
		coordinates = append(coordinates, geom.CreatePoint(array[i]))
	}
	return geom.NewPolygon(coordinates)
}

package box

import (
	"github.com/intdxdt/geom"
	"github.com/intdxdt/mbr"
)

func MBRToPolygon(o *mbr.MBR) *geom.Polygon {
	var coordinates = make([]*geom.Point, 0)
	for _, a := range o.AsPolyArray() {
		coordinates = append(coordinates, geom.NewPoint(a))
	}
	return geom.NewPolygon(coordinates)
}

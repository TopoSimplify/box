package box

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom"
)

func MBRToPolygon(o *mbr.MBR) *geom.Polygon {
	var coordinates = make([]geom.Point, 0)
	for _, a := range o.AsPolyArray() {
		coordinates = append(coordinates, geom.CreatePoint(a))
	}
	return geom.NewPolygon(coordinates)
}

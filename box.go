package box

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom"
)

func MBRToPolygon(o mbr.MBR) *geom.Polygon {
	return geom.NewPolygon(geom.Coordinates(
		[]geom.Point{{o[0], o[1]}, {o[0], o[3]}, {o[2], o[3]}, {o[2], o[1]}, {o[0], o[1]}},
	))
}

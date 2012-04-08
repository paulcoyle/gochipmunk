package cp

/*
#cgo CFLAGS: -DCP_USE_CGPOINTS=0
#cgo LDFLAGS: -lchipmunk
#include <chipmunk/chipmunk.h>
*/
import "C"

func MomentForCircle(mass float64, r1 float64, r2 float64, offset Vect) float64 {
  return float64(
    C.cpMomentForCircle(C.cpFloat(mass), C.cpFloat(r1), C.cpFloat(r2), offset.CPVect),
  )
}

func MomentForBox(mass float64, width float64, height float64) float64 {
  return float64(
    C.cpMomentForBox(C.cpFloat(mass), C.cpFloat(width), C.cpFloat(height)),
  )
}

func MomentForPoly(mass float64, verts []Vect, offset Vect) float64 {
  cpverts := make([]C.cpVect, 0)
  for _, vert := range verts {
    cpverts = append(cpverts, vert.CPVect)
  }
  return float64(C.cpMomentForPoly(C.cpFloat(mass), C.int(len(verts)), &cpverts[0], offset.CPVect))
}

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

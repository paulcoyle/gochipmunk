package cp

/*
#cgo CFLAGS: -DCP_USE_CGPOINTS=0
#cgo LDFLAGS: -lchipmunk
#include <chipmunk/chipmunk.h>
*/
import "C"

type BoundingBox struct {
  CPBB C.cpBB
}

func NewBoundingBox(l float64, b float64, r float64, t float64) *BoundingBox {
  cpbb := C.cpBBNew(C.cpFloat(l), C.cpFloat(b), C.cpFloat(r), C.cpFloat(t))
  boundingbox := BoundingBox{cpbb}
  return &boundingbox
}

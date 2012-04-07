package cp

/*
#cgo CFLAGS: -DCP_USE_CGPOINTS=0
#cgo LDFLAGS: -lchipmunk
#include <chipmunk/chipmunk.h>
*/
import "C"

type Segment struct {
  A Vect
  B Vect
}

func NewSegment(a Vect, b Vect) *Segment {
  seg := Segment{a, b}
  return &seg
}

func (s *Segment) GetMoment(mass float64) float64 {
  return float64(C.cpMomentForSegment(C.cpFloat(mass), s.A.CPVect, s.B.CPVect))
}

func (s *Segment) ToShape(body *Body, radius float64) *Shape {
  return NewSegmentShape(body, s.A, s.B, radius)
}

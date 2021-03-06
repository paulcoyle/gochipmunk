package cp

/*
#cgo CFLAGS: -DCP_USE_CGPOINTS=0
#cgo LDFLAGS: -lchipmunk
#include <chipmunk/chipmunk.h>
*/
import "C"

type Arbiter struct {
  CPArbiter *C.cpArbiter
}

type ContactPoint struct {
  Point    Vect
  Normal   Vect
  Distance float64
}

func NewArbiter(cparb *C.cpArbiter) *Arbiter {
  arbiter := Arbiter{cparb}
  return &arbiter
}

func (a *Arbiter) GetShapes() (*Shape, *Shape) {
  var shape_a, shape_b *C.cpShape
  C.cpArbiterGetShapes(a.CPArbiter, &shape_a, &shape_b)
  return LookupShape(shape_a), LookupShape(shape_b)
}

func (a *Arbiter) GetBodies() (*Body, *Body) {
  var body_a, body_b *C.cpBody
  C.cpArbiterGetBodies(a.CPArbiter, &body_a, &body_b)
  return LookupBody(body_a), LookupBody(body_b)
}

func (a *Arbiter) TotalImpulse() Vect {
  return Vect{C.cpArbiterTotalImpulse(a.CPArbiter)}
}

func (a *Arbiter) TotalImpulseWithFriction() Vect {
  return Vect{C.cpArbiterTotalImpulseWithFriction(a.CPArbiter)}
}

func (a *Arbiter) TotalKineticEnergy() float64 {
  return float64(C.cpArbiterTotalKE(a.CPArbiter))
}

func (a *Arbiter) GetContactPointSet() []*ContactPoint {
  points := make([]*ContactPoint, 0)
  rawSet := C.cpArbiterGetContactPointSet(a.CPArbiter)
  for i := 0; i < int(rawSet.count); i++ {
    point := ContactPoint{NewVectFromCPVect(rawSet.points[i].point),
                          NewVectFromCPVect(rawSet.points[i].normal),
                          float64(rawSet.points[i].dist)}
    points = append(points, &point)
  }
  return points
}

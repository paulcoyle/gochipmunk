package cp

/*
#cgo CFLAGS: -DCP_USE_CGPOINTS=0
#cgo LDFLAGS: -lchipmunk
#include <chipmunk/chipmunk.h>
*/
import "C"

var shapeLookup map[*C.cpShape]*Shape = make(map[*C.cpShape]*Shape)

type Shape struct {
  CPShape *C.cpShape
}

func NewSegmentShape(body *Body, a Vect, b Vect, radius float64) *Shape {
  var cpshape *C.cpShape = C.cpSegmentShapeNew(body.CPBody, a.CPVect, b.CPVect, C.cpFloat(radius))
  shape := Shape{cpshape} 
  shapeLookup[cpshape] = &shape
  return &shape
}

func NewCircleShape(body *Body, radius float64, offset Vect) *Shape {
  var cpshape *C.cpShape = C.cpCircleShapeNew(body.CPBody, C.cpFloat(radius), offset.CPVect)
  shape := Shape{cpshape} 
  shapeLookup[cpshape] = &shape
  return &shape
}

func LookupShape(cpshape *C.cpShape) *Shape {
  return shapeLookup[cpshape]
}

func (s *Shape) Free() {
  delete(shapeLookup, s.CPShape)
  C.cpShapeFree(s.CPShape)
  s.CPShape = nil
}

func (s *Shape) SetFriction(friction float64) {
  s.CPShape.u = C.cpFloat(friction)
}

func (s *Shape) SetCollisionType(ctype uint) {
  s.CPShape.collision_type = C.cpCollisionType(ctype)
}

func (s *Shape) GetCollisionType() uint {
  return uint(s.CPShape.collision_type)
}

func (s *Shape) GetBody() *Body {
  return LookupBody(s.CPShape.body)
}

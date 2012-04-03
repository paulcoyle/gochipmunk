package cp

/*
#cgo CFLAGS: -DCP_USE_CGPOINTS=0
#cgo LDFLAGS: -lchipmunk
#include <chipmunk/chipmunk.h>
*/
import "C"

type Body struct {
  CPBody *C.cpBody
}

var bodyLookup map[*C.cpBody]*Body = make(map[*C.cpBody]*Body)

func NewBody(mass float64, moment float64) *Body {
  var cpbody *C.cpBody = C.cpBodyNew(C.cpFloat(mass), C.cpFloat(moment))
  body := Body{cpbody}
  bodyLookup[cpbody] = &body
  return &body
}

func LookupBody(cpbody *C.cpBody) *Body {
  return bodyLookup[cpbody]
}

func (b *Body) Free() {
  delete(bodyLookup, b.CPBody)
  C.cpBodyFree(b.CPBody)
  b.CPBody = nil
}

func (b *Body) GetPosition() Vect {
  return Vect{C.cpBodyGetPos(b.CPBody)}
}

func (b *Body) SetPosition(pos Vect) {
  C.cpBodySetPos(b.CPBody, pos.CPVect)
}

func (b *Body) GetVelocity() Vect {
  return Vect{C.cpBodyGetVel(b.CPBody)}
}

package cp

/*
#cgo CFLAGS: -DCP_USE_CGPOINTS=0
#cgo LDFLAGS: -lchipmunk
#include <chipmunk/chipmunk.h>
*/
import "C"

type Vect struct {
  CPVect C.cpVect
}

func NewVect(x float64, y float64) Vect {
  return Vect{C.cpVect{C.cpFloat(x), C.cpFloat(y)}}
}

func NewZeroVect() Vect {
  return Vect{C.cpVect{0.0, 0.0}}
}

func (v *Vect) GetX() float64 {
  return float64(v.CPVect.x)
}

func (v *Vect) SetX(x float64) {
  v.CPVect.x = C.cpFloat(x)
}

func (v *Vect) GetY() float64 {
  return float64(v.CPVect.y)
}

func (v *Vect) SetY(y float64) {
  v.CPVect.y = C.cpFloat(y)
}

func (v *Vect) WrapToBounds(ax float64, ay float64, bx float64, by float64) {
  x, y := float64(v.CPVect.x), float64(v.CPVect.y)
  w, h := bx - ax, by - ay
  
  if x < ax {
    x += w
  } else if x >= bx {
    x -= w
  }
  
  if y < ay {
    y += h
  } else if y >= by {
    y -= h
  }
  
  v.CPVect.x = C.cpFloat(x)
  v.CPVect.y = C.cpFloat(y)
}

// TODO MAKE IMMUTABLE
var ZeroVect Vect = NewVect(0.0, 0.0)

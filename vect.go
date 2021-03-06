package cp

/*
#cgo CFLAGS: -DCP_USE_CGPOINTS=0
#cgo LDFLAGS: -lchipmunk
#include <chipmunk/chipmunk.h>
*/
import "C"
import "math"

type Vect struct {
  CPVect C.cpVect
}

func NewVect(x float64, y float64) Vect {
  return Vect{C.cpVect{C.cpFloat(x), C.cpFloat(y)}}
}

func NewVectFromCPVect(v C.cpVect) Vect {
  return Vect{v}
}

func NewZeroVect() Vect {
  return Vect{C.cpVect{0.0, 0.0}}
}

func (v *Vect) GetValues() (float64, float64) {
  return float64(v.CPVect.x), float64(v.CPVect.y)
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

func (v *Vect) Shift(other Vect) Vect {
  x, y   := v.GetValues()
  ox, oy := other.GetValues()
  return NewVect(x+ox, y+oy)
}

func (v *Vect) Multiply(factor float64) Vect {
  return NewVect(v.GetX() * factor, v.GetY() * factor)
}

func (v *Vect) Magnitude() float64 {
  x, y := v.GetValues()
  return math.Sqrt(x*x + y*y)
}

func (v *Vect) Normalize() Vect {
  mag := v.Magnitude()
  return NewVect(v.GetX()/mag, v.GetY()/mag)
}

func (v *Vect) Angle() float64 {
  return math.Atan2(v.GetY(), v.GetX())
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

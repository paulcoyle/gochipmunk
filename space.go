package cp

/*
#cgo CFLAGS: -DCP_USE_CGPOINTS=0
#cgo LDFLAGS: -lchipmunk
#include <chipmunk/chipmunk.h>

void cpgoSpaceAddPostStepCallback(cpSpace *space, void *object, void *data);
cpBool cpgoGenericBeginAndPreHandler(cpArbiter *arb, cpSpace *space, void *data);
void cpgoGenericPostAndSeparateHandler(cpArbiter *arb, cpSpace *space, void *data);
void cpgoSpaceAddBeginCollisionHandler(cpSpace *space, cpCollisionType a, cpCollisionType b, void *data);
void cpgoSpaceAddPreSolveCollisionHandler(cpSpace *space, cpCollisionType a, cpCollisionType b, void *data);
void cpgoSpaceAddPostSolveCollisionHandler(cpSpace *space, cpCollisionType a, cpCollisionType b, void *data);
void cpgoSpaceAddSeparateCollisionHandler(cpSpace *space, cpCollisionType a, cpCollisionType b, void *data);

*/
import "C"
import "unsafe"

//export c2goPostStepCallback
func c2goPostStepCallback(space *C.cpSpace, unused unsafe.Pointer, data unsafe.Pointer) {
  d := *(*PostStepCallbackData)(data)
  d.Callback(d.Space, d.Data)
}

//export c2goCollisionHandler
func c2goCollisionHandler(data unsafe.Pointer, arb *C.cpArbiter) int {
  arbiter := NewArbiter(arb)
  d := *(*CollisionHandlerData)(data)
  return d.Handler(d.Space, arbiter, d.Data)
}

var uglyassshit []*CollisionHandlerData = make([]*CollisionHandlerData, 0)

type Space struct {
  CPSpace    *C.cpSpace
  staticBody *Body
}

type CollisionHandlerData struct {
  Space   *Space
  Data    interface{}
  Handler CollisionHandler
}

type CollisionHandler func(space *Space, arbiter *Arbiter, data interface{}) int

type PostStepCallbackData struct {
  Space    *Space
  Data     interface{}
  Callback PostStepCallback
}

type PostStepCallback func(space *Space, object interface{})

func NewSpace() *Space {
  var cpspace *C.cpSpace = C.cpSpaceNew();
  return &Space{cpspace, nil}
}

func (s *Space) Free() {
  C.cpSpaceFree(s.CPSpace)
  s.CPSpace = nil
}

func (s *Space) StaticBody() *Body {
  if (s.staticBody == nil) {
    s.staticBody = &Body{s.CPSpace.staticBody}
  }
  return s.staticBody
}

func (s *Space) GetIterations() int {
  return int(s.CPSpace.iterations)
}

func (s *Space) SetIterations(iter int) {
  s.CPSpace.iterations = C.int(iter)
}

func (s *Space) SetGravity(v Vect) {
  s.CPSpace.gravity = v.CPVect
}

func (s *Space) AddBody(body *Body) {
  C.cpSpaceAddBody(s.CPSpace, body.CPBody)
}

func (s *Space) RemoveBody(body *Body) {
  C.cpSpaceRemoveBody(s.CPSpace, body.CPBody)
}

func (s *Space) AddShape(shape *Shape) {
  C.cpSpaceAddShape(s.CPSpace, shape.CPShape)
}

func (s *Space) RemoveShape(shape *Shape) {
  C.cpSpaceRemoveShape(s.CPSpace, shape.CPShape)
}

func (s *Space) AddStaticShape(shape *Shape) {
  C.cpSpaceAddStaticShape(s.CPSpace, shape.CPShape)
}

func (s *Space) RemoveStaticShape(shape *Shape) {
  C.cpSpaceRemoveStaticShape(s.CPSpace, shape.CPShape)
}

func (s *Space) AddPostStepCallback(data interface{}, callback PostStepCallback) {
  c := PostStepCallbackData{s, data, callback}
  C.cpgoSpaceAddPostStepCallback(s.CPSpace, unsafe.Pointer(&data), unsafe.Pointer(&c))
}

func (s *Space) AddBeginCollisionHandler(typeA uint, typeB uint, data interface{}, handler CollisionHandler) {
  h := CollisionHandlerData{s, data, handler}
  uglyassshit = append(uglyassshit, &h)
  C.cpgoSpaceAddBeginCollisionHandler(s.CPSpace,
    C.cpCollisionType(typeA), C.cpCollisionType(typeB), unsafe.Pointer(&h))
}

func (s *Space) AddPreSolveCollisionHandler(typeA uint, typeB uint, data interface{}, handler CollisionHandler) {
  h := CollisionHandlerData{s, data, handler}
  uglyassshit = append(uglyassshit, &h)
  C.cpgoSpaceAddPreSolveCollisionHandler(s.CPSpace,
    C.cpCollisionType(typeA), C.cpCollisionType(typeB), unsafe.Pointer(&h))
}

func (s *Space) AddPostSolveCollisionHandler(typeA uint, typeB uint, data interface{}, handler CollisionHandler) {
  h := CollisionHandlerData{s, data, handler}
  uglyassshit = append(uglyassshit, &h)
  C.cpgoSpaceAddPostSolveCollisionHandler(s.CPSpace,
    C.cpCollisionType(typeA), C.cpCollisionType(typeB), unsafe.Pointer(&h))
}

func (s *Space) AddSeparateCollisionHandler(typeA uint, typeB uint, data interface{}, handler CollisionHandler) {
  h := CollisionHandlerData{s, data, handler}
  uglyassshit = append(uglyassshit, &h)
  C.cpgoSpaceAddSeparateCollisionHandler(s.CPSpace,
    C.cpCollisionType(typeA), C.cpCollisionType(typeB), unsafe.Pointer(&h))
}

func (s *Space) Step(dt float64) {
  C.cpSpaceStep(s.CPSpace, C.cpFloat(dt))
}

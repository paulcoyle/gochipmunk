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
  d.Callback(d.Space, d.Object)
}

//export c2goCollisionHandler
func c2goCollisionHandler(data unsafe.Pointer, arb *C.cpArbiter) int {
  arbiter := NewArbiter(arb)
  d := *(*CollisionHandlerData)(data)
  return d.Handler(d.Space, arbiter)
}

var uglyassshit []*CollisionHandlerData = make([]*CollisionHandlerData, 0)

type Space struct {
  CPSpace    *C.cpSpace
  staticBody *Body
}

type CollisionHandlerData struct {
  Space   *Space
  Handler  CollisionHandler
}

type CollisionHandler func(space *Space, arbiter *Arbiter) int

type PostStepCallbackData struct {
  Space    *Space
  Object    interface{}
  Callback  PostStepCallback
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

func (s *Space) AddPostStepCallback(object interface{}, callback PostStepCallback) {
  data := PostStepCallbackData{s, object, callback}
  C.cpgoSpaceAddPostStepCallback(s.CPSpace, unsafe.Pointer(&object), unsafe.Pointer(&data))
}

func (s *Space) AddBeginCollisionHandler(typeA uint, typeB uint, begin CollisionHandler) {
  data := CollisionHandlerData{s, begin}
  uglyassshit = append(uglyassshit, &data)
  C.cpgoSpaceAddBeginCollisionHandler(s.CPSpace,
    C.cpCollisionType(typeA), C.cpCollisionType(typeB), unsafe.Pointer(&data))
}

func (s *Space) AddPreSolveCollisionHandler(typeA uint, typeB uint, begin CollisionHandler) {
  data := CollisionHandlerData{s, begin}
  uglyassshit = append(uglyassshit, &data)
  C.cpgoSpaceAddPreSolveCollisionHandler(s.CPSpace,
    C.cpCollisionType(typeA), C.cpCollisionType(typeB), unsafe.Pointer(&data))
}

func (s *Space) AddPostSolveCollisionHandler(typeA uint, typeB uint, begin CollisionHandler) {
  data := CollisionHandlerData{s, begin}
  uglyassshit = append(uglyassshit, &data)
  C.cpgoSpaceAddPostSolveCollisionHandler(s.CPSpace,
    C.cpCollisionType(typeA), C.cpCollisionType(typeB), unsafe.Pointer(&data))
}

func (s *Space) AddSeparateCollisionHandler(typeA uint, typeB uint, begin CollisionHandler) {
  data := CollisionHandlerData{s, begin}
  uglyassshit = append(uglyassshit, &data)
  C.cpgoSpaceAddSeparateCollisionHandler(s.CPSpace,
    C.cpCollisionType(typeA), C.cpCollisionType(typeB), unsafe.Pointer(&data))
}

func (s *Space) Step(dt float64) {
  C.cpSpaceStep(s.CPSpace, C.cpFloat(dt))
}

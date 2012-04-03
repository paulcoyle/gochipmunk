#include <chipmunk/chipmunk.h>
#include <stdio.h>
#include "_cgo_export.h"

void cpgoSpaceAddPostStepCallback(cpSpace *space, void *object, void *data) {
  cpSpaceAddPostStepCallback(space, &c2goPostStepCallback, object, data);
}

cpBool cpgoGenericBeginHandler(cpArbiter *arb, cpSpace *space, void *data) {
  cpShape *a, *b;
  cpArbiterGetShapes(arb, &a, &b);
  return c2goCollisionHandler(data, a, b);
}

void cpgoSpaceAddBeginCollisionHandler(cpSpace *space, cpCollisionType a, cpCollisionType b, void *data) {
  cpSpaceAddCollisionHandler(space, a, b, &cpgoGenericBeginHandler, NULL, NULL, NULL, data);
}


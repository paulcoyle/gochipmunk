#include <chipmunk/chipmunk.h>
#include <stdio.h>
#include "_cgo_export.h"

void cpgoSpaceAddPostStepCallback(cpSpace *space, void *object, void *data) {
  cpSpaceAddPostStepCallback(space, &c2goPostStepCallback, object, data);
}

cpBool cpgoGenericBeginAndPreHandler(cpArbiter *arb, cpSpace *space, void *data) {
  return c2goCollisionHandler(data, arb);
}

void cpgoGenericPostAndSeparateHandler(cpArbiter *arb, cpSpace *space, void *data) {
  c2goCollisionHandler(data, arb);
}

void cpgoSpaceAddBeginCollisionHandler(cpSpace *space, cpCollisionType a, cpCollisionType b, void *data) {
  cpSpaceAddCollisionHandler(space, a, b, &cpgoGenericBeginAndPreHandler, NULL, NULL, NULL, data);
}

void cpgoSpaceAddPreSolveCollisionHandler(cpSpace *space, cpCollisionType a, cpCollisionType b, void *data) {
  cpSpaceAddCollisionHandler(space, a, b, NULL, &cpgoGenericBeginAndPreHandler, NULL, NULL, data);
}

void cpgoSpaceAddPostSolveCollisionHandler(cpSpace *space, cpCollisionType a, cpCollisionType b, void *data) {
  cpSpaceAddCollisionHandler(space, a, b, NULL, NULL, &cpgoGenericPostAndSeparateHandler, NULL, data);
}

void cpgoSpaceAddSeparateCollisionHandler(cpSpace *space, cpCollisionType a, cpCollisionType b, void *data) {
  cpSpaceAddCollisionHandler(space, a, b, NULL, NULL, NULL, &cpgoGenericPostAndSeparateHandler, data);
}


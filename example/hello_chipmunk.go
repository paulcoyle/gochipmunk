// A gochipmunk implementation of the Hello Chipmunk example found at
//
// http://chipmunk-physics.net/release/ChipmunkLatest-Docs/
//
// This implementation also adds a basic collision handler and post-step
// callback.
package main

import (
  "github.com/paulcoyle/gochipmunk"
  "fmt"
  "time"
)

// Runs once at the end of a step when added to a space with
// space.AddPostStepCallback(object interface{}, callback PostStepCallback)
func postStep(space *cp.Space, object interface{}) {
  fmt.Println("POST-STEP CALLBACK")
  shape, ok := object.(*cp.Shape)
  if ok {
    space.RemoveBody(shape.GetBody())
    space.RemoveShape(shape)
  } else {
    fmt.Print("Couldn't resolve object into cp.Shape\n")
  }
}

// Handler for begin-collision events.  The chipmunk library only allows one
// of these per categorical pair of collision types.  That is:
//    (a,b) == (b,a)
// where a and b are collision types.
func beginCollide(space *cp.Space, arb *cp.Arbiter) int {
  fmt.Println("BEGIN COLLISION")
  
  // Uncomment to add a post step callback to remove the ball
  //space.AddPostStepCallback(b, postStep)
  
  // Change return value to 0 to skip the collision chain
  return 1
}

// Collision type constants
const SEGM uint = 1
const BALL uint = 2

func main() {
  space  := cp.NewSpace()
  space.SetGravity(cp.NewVect(0, -100))
  
  // static ground
  ground := cp.NewSegmentShape(space.StaticBody(), cp.NewVect(-20, 5), cp.NewVect(20, -5), 0)
  ground.SetFriction(1)
  ground.SetCollisionType(SEGM)
  space.AddShape(ground)
  
  var radius float64 = 5
  var mass   float64 = 1
  var moment float64 = cp.MomentForCircle(mass, 0, radius, cp.ZeroVect)
  
  ball := cp.NewBody(mass, moment)
  ball.SetPosition(cp.NewVect(0, 15))
  space.AddBody(ball)
  
  ballShape := cp.NewCircleShape(ball, radius, cp.ZeroVect)
  ballShape.SetFriction(0.7)
  ballShape.SetCollisionType(BALL)
  space.AddShape(ballShape)
  
  space.AddBeginCollisionHandler(SEGM, BALL, beginCollide)
  
  var start int64 = time.Now().UnixNano()
  var dt float64 = 1.0/60.0
  for time := 0.0; time < 2.0; time += dt {
    var pos cp.Vect = ball.GetPosition()
    var vel cp.Vect = ball.GetVelocity()
    fmt.Printf("Time is %5.2f. ball is at (%5.2f, %5.2f). It's velocity is (%5.2f, %5.2f)\n",
      time, pos.GetX(), pos.GetY(), vel.GetX(), vel.GetY())
    space.Step(dt)
  }
  fmt.Printf("TOTAL TIME: %dns\n", time.Now().UnixNano() - start)
  
  ballShape.Free()
  ball.Free()
  ground.Free()
  space.Free()
}

# gochipmunk!

This is a Go binding to the [Chipmunk physics library](http://chipmunk-physics.net/)
and is currently a work-in-progress.


## Installing

This library requires that the Chipmunk physics library is compiled and
installed (see the notes on Compatibility below).  This library targets Go1
so you will need the release version of Go installed (as of this writing).

Finally, use the `go get` command to fetch the library:

    go get github.com/paulcoyle/gochipmunk

Then, in your code:

    import "github.com/paulcoyle/gochipmunk"


## Integration

At this point, there is enough functionality to run the
[Hello Chipmunk (World)](http://chipmunk-physics.net/release/ChipmunkLatest-Docs/)
as well as begin-collision handlers and post-step callbacks.


## Compatibility

Currently, this library targets Go1 but there are plans to maintain a branch
which keeps up with the weekly Go releases.

This wrapper does not use the Apple-specfic CGPoint for vectors.  When
building Chipmunk be sure to use the following when running cmake:

    cmake -D CMAKE_C_FLAGS:STRING=-DCP_USE_CGPOINTS=0 .

## Documentation?

Forthcoming.

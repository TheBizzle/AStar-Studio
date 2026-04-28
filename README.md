AStar-Studio
===============

## What is it?

A web application that wraps around [my Golang A* implementation](https://github.com/TheBizzle/AStar-Golang).

Overkill?  Yeah, probably.

Pointless?  Yeah, probably.

But also educational!

## Why is it?

I wanted to actually give Go's concurrency constructs a test drive, so I contrived something that could build upon my other pathfinding work with Go. 🤷‍♂️

## How do you use it?

  * `go mod download` to set things up
  * `go run .` to launch the server
  * Navigate to `localhost:8080` in your web browser
  * Follow the instructions on the page

## What do the results mean?

Not much!  I just run A* with each of the heuristics 1000 times, and figure out which heuristic completed its 1000 runs the fastest.  This is benchmarking at its worst.  But it's a little interesting, I guess.

Ultimately, the "fastest" heuristic for a run is a bit random.  As a result, clicking the "Run Pathfinding" button repeatedly will give different results with each click (but not to the extent that some tries will fail while others will succeed).

Also note that the non-heuristic option (Dijkstra's algorithm) will check a bunch of nodes that the other options will not check, because the non-heuristic mode doesn't try to do anything clever.

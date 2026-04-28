// Package runner runs A* in multiple threads and reports about how long each run took.  This is *not* an
// effective way of benchmarking different algorithms/heuristics, so don't read anything too specific into
// the numbers that come out of it.
package runner

import (
	"fmt"
	"math"
	"sync"
	"time"

	astar "github.com/TheBizzle/AStar-Golang/astar"
	heur "github.com/TheBizzle/AStar-Golang/heuristic"
	core "github.com/TheBizzle/PathFindingCore-Golang/pathingmap"
)

type Result struct {
	Heuristic heur.Heuristic
	PMap      *core.PathingMap
	Duration  time.Duration
}

type candidate struct {
	Heur heur.Heuristic
	Time uint64
}

var heuristics = []heur.Heuristic{heur.Euclidean, heur.Manhattan, heur.Dijkstra}

func RunAStars(asciiMap string) (bool, string, []string) {
	pms := core.PathingMapString{Contents: asciiMap, Delim: "|\n"}

	numAttempts := 1000
	results := make([]Result, len(heuristics)*numAttempts)
	var wg sync.WaitGroup

	atLeastOneSucceeded := false
	count := 0

	for _, heuristic := range heuristics {
		for range numAttempts {
			wg.Add(1)
			go func(h heur.Heuristic, i int) {
				defer wg.Done()

				start := time.Now()
				runResult, stepData := astar.Run(pms, h)

				var duration time.Duration
				if runResult == core.SuccessfulRun {
					duration = time.Since(start)
					atLeastOneSucceeded = true
				} else {
					duration = -1
				}

				results[i] = Result{
					Heuristic: h,
					PMap:      &stepData.PathingMap,
					Duration:  duration,
				}
			}(heuristic, count)
			count++
		}
	}

	wg.Wait()

	if atLeastOneSucceeded {
		return gatherStats(results, numAttempts)
	} else {
		return false, fmt.Sprint(results[0].PMap), nil
	}
}

func gatherStats(results []Result, numAttempts int) (bool, string, []string) {
	var timingStrs []string

	times := map[heur.Heuristic]uint64{}

	exemplars := map[heur.Heuristic]string{}

	for _, result := range results {
		if times[result.Heuristic] != math.MaxUint64 {
			if result.Duration != -1 {
				times[result.Heuristic] += uint64(result.Duration)
			} else {
				times[result.Heuristic] = math.MaxUint64
			}
		}
		if exemplars[result.Heuristic] == "" {
			exemplars[result.Heuristic] = fmt.Sprint(result.PMap)
		}
	}

	fastest := candidate{Heur: heur.Euclidean, Time: math.MaxUint64}

	for _, h := range heuristics {
		sum := times[h]
		if sum != math.MaxUint64 {
			average := sum / uint64(numAttempts)
			if average < fastest.Time {
				fastest = candidate{Heur: h, Time: average}
			}
			timingStrs = append(timingStrs, fmt.Sprint(average))
		} else {
			timingStrs = append(timingStrs, "-1")
		}
	}

	return true, exemplars[fastest.Heur], timingStrs
}

package algo

import "math/rand"

type Interval struct {
	Start, End int
}

/**
** Problem: https://www.lintcode.com/problem/merge-intervals/leaderboard
** Approach:
**	- Quick sort the intervals by the Start
**	- Merge the duplicates
** Complexity: O(n*log n)
**	- Quick sort: O(n*log n) + Merge: O(n)
**/

// MergeIntervals merges the duplicated space of intervals.
func MergeIntervals(intervals []*Interval) []*Interval {
	if len(intervals) < 2 {
		return intervals
	}

	// Sort intervals with the Start
	intervals = quickSortIntervals(intervals)

	j := 0
	for _, in := range intervals {
		if in.Start < intervals[j].Start && in.End >= intervals[j].Start {
			intervals[j].Start = in.Start
		}
		if in.Start <= intervals[j].End && in.End > intervals[j].End {
			intervals[j].End = in.End
		}
		if in.Start > intervals[j].End || in.End < intervals[j].Start {
			j++
			intervals[j] = in
		}
	}

	return intervals[:j+1]
}

func quickSortIntervals(intervals []*Interval) []*Interval {
	if len(intervals) < 2 {
		return intervals
	}

	left, right := 0, len(intervals)-1

	pivot := rand.Int() % len(intervals)

	intervals[pivot], intervals[right] = intervals[right], intervals[pivot]

	for i := range intervals {
		if intervals[i].Start < intervals[right].Start {
			intervals[left], intervals[i] = intervals[i], intervals[left]
			left++
		}
	}

	intervals[left], intervals[right] = intervals[right], intervals[left]

	quickSortIntervals(intervals[:left])
	quickSortIntervals(intervals[left+1:])

	return intervals
}

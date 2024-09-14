package interviewcake

import (
	"101coding/common"
	"sort"
	"testing"
)

/*
Problem:
  - Given a list of unsorted, independent meetings, returns a list of a merged
    one.

Example:
  - Input: []meeting{{1, 2}, {2, 3}, {4, 5}}
    Output: []meeting{{1, 3}, {4, 5}}
  - Input: []meeting{{1, 5}, {2, 3}}
    Output: []meeting{{1, 5}}

Approach:
  - Sort the list in ascending order so that meetings that might need to be
    merged are next to each other.
  - Can merge two meetings together if the first one's end time is greater or
    or equal than the second one's start time.

Solution:
  - Sort the list in ascending order.
  - Create a new list of merged meetings and consider the first meeting in the
    original list to be the last merged one.
  - Iterate through the original list and verify if the last merged meeting's
    end time is greater or equal than the current meeting's start time.
  - If it is true, merge them using the last merged meeting's start time
    and the larger one's end time.

Cost:
  - O(nlogn) time, O(n) space.
  - Because we sort all meeting first, the runtime is O(nlogn). We create a new
    list of merged meeting times, so the space cost is O(n).
*/
func TestMergeMeetings(t *testing.T) {
	tests := []struct {
		in       []meeting
		expected []meeting
	}{
		{[]meeting{}, []meeting{}},
		{[]meeting{{1, 2}}, []meeting{{1, 2}}},
		{[]meeting{{1, 2}, {2, 3}}, []meeting{{1, 3}}},
		{[]meeting{{1, 5}, {2, 3}}, []meeting{{1, 5}}},
		{[]meeting{{1, 2}, {4, 5}}, []meeting{{1, 2}, {4, 5}}},
		{[]meeting{{1, 5}, {2, 3}, {4, 5}}, []meeting{{1, 5}}},
		{[]meeting{{1, 2}, {2, 3}, {4, 5}}, []meeting{{1, 3}, {4, 5}}},
		{[]meeting{{1, 6}, {2, 3}, {4, 5}}, []meeting{{1, 6}}},
		{[]meeting{{4, 5}, {2, 3}, {1, 6}}, []meeting{{1, 6}}},
	}
	for _, tt := range tests {
		//out := mergeMeetings(tt.in)
		out := mergeMeetingsV2(tt.in)
		common.Equal(t, tt.expected, out)
	}
}

type meeting struct {
	start int
	end   int
}

func mergeMeetings(meetings []meeting) []meeting {
	var res []meeting
	for i, m := range meetings {
		m2 := meeting{
			start: m.start,
			end:   m.end,
		}
		if i == 0 {
			res = append(res, m2)
			continue
		}
		preM := res[len(res)-1]
		if preM.end >= m2.start && preM.end < m2.end {
			preM.end = m2.end
			res[i-1] = preM
			continue
		}
		if preM.end >= m2.start && preM.end >= m2.end {
			continue
		}
		res = append(res, m2)
	}
	return res
}

func mergeMeetingsV2(meetings []meeting) []meeting {

	// sort with start
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].start < meetings[j].start
	})

	out := []meeting{}
	for i := range meetings {
		if i == 0 {
			out = append(out, meetings[i])
			continue
		}
		if out[len(out)-1].end >= meetings[i].start {
			out[len(out)-1].end = common.Max(meetings[i].end, out[len(out)-1].end)
		} else {
			out = append(out, meetings[i])
		}

	}
	return out
}

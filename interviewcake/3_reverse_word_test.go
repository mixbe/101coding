/*
Problem:
- Given a list of string that is made up of word but in reverse, return the
  correct order in-place.

Example:
- Input: []string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"}
  Output: []string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"}

Approach:
- Similar to reversing string, use the same idea to reverse all the characters
  in the list first so that we could have a list of words in the right order, not
  its characters.
- Iterate through the list again and reverse its characters.

Solution:
- Reverse all the characters to get a list of words in the right order using
  same idea as reversing a string.
- Iterate through the list again the reverse its characters by first keeping
  track of the start index for each word since they are separated by an empty
  string.
- Once we fine an empty string, update the start word index and use the same
  idea to reverse the characters order.

Cost:
- O(n) time, O(1) space.
*/

package interviewcake

import (
	"101coding/common"
	"testing"
)

func TestReverseWord(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{
			[]string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"},
			[]string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"},
		},
	}
	for _, tt := range tests {
		//result := reverseWord(tt.in)
		//common.Equal(t, tt.expected, result)
		res := reverseWordV2(tt.in)
		common.Equal(t, tt.expected, res)
	}
}

func reverseWord(list []string) []string {
	var res []string
	length := len(list)
	for i := range list {
		index := length - i - 1
		world := list[index]

		if world != "" && index != 0 {
			continue
		}
		i2 := index + 1
		if index == 0 {
			i2 = 0
		}
		for j := i2; j < length; j++ {
			if list[j] == "" {
				break
			}
			res = append(res, list[j])
		}
		if index != 0 {
			res = append(res, "")
		}
	}
	return res
}

func reverseWordV2(list []string) []string {
	// "w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"
	// y a s "" o l l e h "" d l r o w

	reverseChar(list, 0, len(list)-1)
	start := 0
	for i := range list {
		if i == len(list)-1 {
			reverseChar(list, start, i)
		}
		if list[i] == "" {
			reverseChar(list, start, i-1)
			start = i + 1
		}
	}
	return list
}

func reverseChar(list []string, start int, end int) {

	for start < end {
		common.Swap(list, start, end)
		start++
		end--
	}
}

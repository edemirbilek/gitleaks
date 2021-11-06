package detect

import (
	"testing"
)

// TestGetLocation tests the getLocation function.
func TestGetLocation(t *testing.T) {
	tests := []struct {
		linePairs    [][]int
		start        int
		end          int
		wantLocation Location
	}{
		{
			linePairs: [][]int{
				{0, 39},
				{40, 55},
				{56, 57},
			},
			start: 35,
			end:   38,
			wantLocation: Location{
				startLine:      1,
				startColumn:    36,
				endLine:        1,
				endColumn:      38,
				startLineIndex: 1,
				endLineIndex:   40,
			},
		},
		{
			linePairs: [][]int{
				{0, 39},
				{40, 55},
				{56, 57},
			},
			start: 40,
			end:   44,
			wantLocation: Location{
				startLine:      2,
				startColumn:    1,
				endLine:        2,
				endColumn:      4,
				startLineIndex: 41,
				endLineIndex:   56,
			},
		},
	}

	for _, test := range tests {
		loc := getLocation(test.linePairs, test.start, test.end)
		if loc != test.wantLocation {
			t.Errorf("\nstartLine %d\nstartColumn: %d\nendLine: %d\nendColumn: %d\nstartLineIndex: %d\nendlineIndex %d",
				loc.startLine, loc.startColumn, loc.endLine, loc.endColumn, loc.startLineIndex, loc.endLineIndex)

			t.Error("got", loc, "want", test.wantLocation)
		}
	}
}

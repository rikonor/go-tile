package smarttile

import (
	"reflect"
	"testing"
)

func TestTile(t *testing.T) {
	type testCase struct {
		in  *TilingInput
		out *TilingResult
	}

	testCases := []testCase{
		testCase{
			in: &TilingInput{Width: 1000, Height: 1000},
			out: &TilingResult{
				TileWidth: 100, TileHeight: 100,
				XAxisTiles: 10, YAxisTiles: 10,
				XAxisOffset: 0, YAxisOffset: 0,
			},
		},
	}

	for _, tc := range testCases {
		tilingResult, err := Tile(tc.in)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if !reflect.DeepEqual(tilingResult, tc.out) {
			t.Fatalf("unexpected result, expected %+v but got %+v\n", tc.out, tilingResult)
		}
	}
}

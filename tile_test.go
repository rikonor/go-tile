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
		// Perfectly divisible square
		testCase{
			in: &TilingInput{
				Width: 100, Height: 100,
				TileDensity: 100,
			},
			out: &TilingResult{
				TileWidth: 10, TileHeight: 10,
				XAxisTiles: 10, YAxisTiles: 10,
				XAxisOffset: 0, YAxisOffset: 0,
			},
		},

		// Perfectly divisible rectangle
		testCase{
			in: &TilingInput{
				Width: 200, Height: 100,
				TileDensity: 100,
			},
			out: &TilingResult{
				TileWidth: 10, TileHeight: 10,
				XAxisTiles: 20, YAxisTiles: 10,
				XAxisOffset: 0, YAxisOffset: 0,
			},
		},

		// Non-divisible square
		testCase{
			in: &TilingInput{
				Width: 105, Height: 115,
				TileDensity: 100,
			},
			out: &TilingResult{
				TileWidth: 10, TileHeight: 10,
				XAxisTiles: 10, YAxisTiles: 11,
				XAxisOffset: 2, YAxisOffset: 2,
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

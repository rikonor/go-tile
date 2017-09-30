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
		// Perfectly divisible square - #1
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

		// Perfectly divisible square - #2
		testCase{
			in: &TilingInput{
				Width: 1000, Height: 1000,
				TileDensity: 1,
			},
			out: &TilingResult{
				TileWidth: 100, TileHeight: 100,
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

		// Non-divisible rectangle - #1
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

		// Non-divisible rectangle - #2
		testCase{
			in: &TilingInput{
				Width: 1003, Height: 807,
				TileDensity: 20,
			},
			out: &TilingResult{
				TileWidth: 22, TileHeight: 22,
				XAxisTiles: 45, YAxisTiles: 36,
				XAxisOffset: 6, YAxisOffset: 7,
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

func TestDensitySize(t *testing.T) {
	type testCase struct {
		// input
		tileSize int

		// output
		tileDensity    float64
		actualTileSize int
	}

	testCases := []testCase{
		// Exact sizes
		testCase{
			tileSize:       100,
			tileDensity:    1,
			actualTileSize: 100,
		},
		testCase{
			tileSize:       10,
			tileDensity:    100,
			actualTileSize: 10,
		},
		testCase{
			tileSize:       20,
			tileDensity:    25,
			actualTileSize: 20,
		},
		testCase{
			tileSize:       1,
			tileDensity:    10000,
			actualTileSize: 1,
		},
		testCase{
			tileSize:       16,
			tileDensity:    39.0625,
			actualTileSize: 16,
		},

		// Non-exact sizes
		testCase{
			tileSize:       17,
			tileDensity:    34.602076,
			actualTileSize: 17,
		},
		testCase{
			tileSize:       48,
			tileDensity:    4.340278,
			actualTileSize: 48,
		},
	}

	for _, tc := range testCases {
		tileDensity, actualSize := TileDensity(tc.tileSize)
		if !equalFloat64(tileDensity, tc.tileDensity) || actualSize != tc.actualTileSize {
			t.Fatalf("wrong results, expected (%f, %d) but got (%f, %d)", tc.tileDensity, tc.actualTileSize, tileDensity, actualSize)
		}
	}
}

// equalFloat64 checks if two float64 values are equal up to a certain threshold
// it does this by first trimming to a certain number of digits before comparing them
func equalFloat64(a, b float64) bool {
	// Compare up to 6 decimal digits
	eps := 1e-6

	// Check if difference between numbers falls within band of width 2 x epsilon
	return (a-b) < eps && (b-a) < eps
}

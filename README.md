go-tile
---

This package is used for block tiling calculations.

Things to consider:

1. Tiles have to be square, meaning a tile dimensions have to be `n x m` where `n` is equal to `m`.
1. Tile size cannot be fractional, meaning `n` and `m` have to be integer numbers.
1. Because of the above requirements, in some cases an image has to be cropped in order to be tiled.

### Installation

```
go get -u github.com/rikonor/go-tile
```

### Usage

```go
package main

import (
	"fmt"
	"log"

	tile "github.com/rikonor/go-tile"
)

func main() {
	// Lets say we want to tile a block of size 1012x2036 with square tiles
	result, err := tile.Tile(&tile.TilingInput{
		// Input the dimensions of our block
		Width:  1012,
		Height: 2036,

		// The size of the tiles is driven by a tile density parameter
		// The higher the density, the more tiles we expect to have in a given area
		TileDensity: tile.DensityMedium,
	})

	// Handle error...

	// The results of this operation are:

	// The size of the tile
	fmt.Printf("dx=%d, dy=%d\n", result.TileWidth, result.TileHeight)

	// The number of tiles in each direction
	fmt.Printf("Nx=%d, Ny=%d\n", result.XAxisTiles, result.YAxisTiles)

	// The required offset on each axis
	// NOTE: In some cases, an image can't be perfectly tiled and needs to be cropped
	// the offset defines the cropping
	fmt.Printf("OffsetX=%d, OffsetY=%d\n", result.XAxisOffset, result.YAxisOffset)

}
```

### Helpers

*_TileDensity_*

```go
// Let's say that we'd like to aim for a given tile size
// Use the TileDensity helper to calculate the necessary density
desiredTileSize := 35

tileDensity, actualSize := tile.TileDensity(desiredTileSize)

// Not every tile size can be reached, so the actual achieved tile size is also returned
fmt.Printf("To achieve tile size %d use tile density %f\n", actualSize, tileDensity)
```

### LICENSE

MIT

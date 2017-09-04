package smarttile

import "math"

// Common density values
const (
	DensityLow    = 1.0
	DensityMedium = 10.0
	DensityHigh   = 100.0

	// Default density when none specified
	DefaultDensity = DensityHigh
)

const (
	densityReferenceBlockSize = 100
)

// TilingInput is the input given to the Tile operation
type TilingInput struct {
	// Width and Height of what we want to tile in pixels
	Width, Height int

	// TileDensity defines the density of our tiling
	// as the number of tiles expected in a 100x100 block
	// The higher the density the more tiles we will have
	TileDensity float64
}

// TilingResult is the result of the Tiling operation
type TilingResult struct {
	// Tile dimensions in pixels
	TileWidth, TileHeight int

	// Count of tiles in each direction
	XAxisTiles, YAxisTiles int

	// Offset information due to cropping (given in pixels)
	XAxisOffset, YAxisOffset int
}

// Tile is used to tile a rectangle
func Tile(tileInput *TilingInput) (*TilingResult, error) {
	// If no density provided, fallback to default value
	tileDensity := tileInput.TileDensity
	if tileDensity == 0 {
		tileDensity = DefaultDensity
	}

	// Change names for easy handling
	Dx, Dy := tileInput.Width, tileInput.Height

	// Number of tiles in a subblock
	dsSqrt := math.Sqrt(tileDensity)
	nx := densityReferenceBlockSize / int(dsSqrt)
	ny := densityReferenceBlockSize / int(dsSqrt)

	// Size of the tiles
	dx := densityReferenceBlockSize / nx
	dy := densityReferenceBlockSize / ny

	// Number of tiles
	Nx := Dx / dx
	Ny := Dy / dy

	// Offsets
	XLeftSpace := Dx - (dx * Nx)
	YLeftSpace := Dy - (dy * Ny)

	XAxisOffset := XLeftSpace / 2
	YAxisOffset := YLeftSpace / 2

	return &TilingResult{
		TileWidth: dx, TileHeight: dy,
		XAxisTiles: Nx, YAxisTiles: Ny,
		XAxisOffset: XAxisOffset, YAxisOffset: YAxisOffset,
	}, nil
}

package smarttile

// TilingInput is the input given to the Tile operation
type TilingInput struct {
	// Width and Height of what we want to tile in pixels
	Width, Height int
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
	return nil, nil
}

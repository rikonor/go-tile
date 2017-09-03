go-tile
---

This package is used for image tiling calculations.

### Motivation

Tile an `N x M` image, where `N` is the number of X-axis pixels and `M` is the number of Y-axis pixels.

Things to consider:

1. Tiles have to be square, meaning a tile dimensions have to be `n x m` where `n` is equal to `m`.
1. Tile size cannot be fractional, meaning `n` and `m` have to be integer numbers.
1. Because of the above requirements, in some cases an image has to be cropped in order to be tiled.

The output of the tiling operation should be:

1. Tile dimension - Width and Height in pixels.
1. Number of X-axis tiles.
1. Number of Y-axis tiles.
1. Cropping information - X-axis offset and Y-axis offset in pixels.

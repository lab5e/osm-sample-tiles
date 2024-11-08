package osmsampletiles

import (
	"embed"
	"io/fs"
)

//go:embed map/tiles
var tiles embed.FS

// SampleOSMTiles returns an embeddded file system with Open Street Map tiles containting sample data. The
// area is roughly the city center of Trondheim.
func SampleOSMTiles() fs.FS {
	return tiles
}

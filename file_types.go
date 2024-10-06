package media_vector

import (
	"github.com/smartmediafiles/media/media/maps"
	"github.com/smartmediafiles/media/media/types"
)

// List of supported media.Vector file types.
const (
	VectorSvg types.FileType = "svg" // Scalable Vector Graphics
	VectorAi  types.FileType = "ai"  // Adobe Illustrator
	VectorPs  types.FileType = "ps"  // Adobe PostScript
	VectorEps types.FileType = "eps" // Encapsulated PostScript
)

// VectorFileTypesExtensions is a map of supported media.Vector file types and their extensions.
var VectorFileTypesExtensions = maps.MapFileTypeExtensions{
	VectorSvg: {".svg"},
	VectorAi:  {".ai"},
	VectorPs:  {".ps", ".ps2", ".ps3"},
	VectorEps: {".eps", ".eps2", ".eps3", ".epi", ".ept", ".epsf", ".epsi"},
}

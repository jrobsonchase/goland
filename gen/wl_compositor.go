package gen

// A compositor.  This object is a singleton global.  The
// compositor is in charge of combining the contents of multiple
// surfaces into one displayable output.
type WlCompositor interface {
	// Ask the compositor to create a new surface.
	CreateSurface(Id WlNewId)
	// Ask the compositor to create a new region.
	CreateRegion(Id WlNewId)
}

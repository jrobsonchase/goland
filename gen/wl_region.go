package gen

// A region object describes an area.
// Region objects are used to describe the opaque and input
// regions of a surface.
type WlRegion interface {
	// Destroy the region.  This will invalidate the object ID.
	Destroy()
	// Add the specified rectangle to the region.
	Add(X WlInt, Y WlInt, Width WlInt, Height WlInt)
	// Subtract the specified rectangle from the region.
	Subtract(X WlInt, Y WlInt, Width WlInt, Height WlInt)
}

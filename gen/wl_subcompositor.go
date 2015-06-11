package gen

// The global interface exposing sub-surface compositing capabilities.
// A wl_surface, that has sub-surfaces associated, is called the
// parent surface. Sub-surfaces can be arbitrarily nested and create
// a tree of sub-surfaces.
// The root surface in a tree of sub-surfaces is the main
// surface. The main surface cannot be a sub-surface, because
// sub-surfaces must always have a parent.
// A main surface with its sub-surfaces forms a (compound) window.
// For window management purposes, this set of wl_surface objects is
// to be considered as a single window, and it should also behave as
// such.
// The aim of sub-surfaces is to offload some of the compositing work
// within a window from clients to the compositor. A prime example is
// a video player with decorations and video in separate wl_surface
// objects. This should allow the compositor to pass YUV video buffer
// processing to dedicated overlay hardware when possible.
type WlSubcompositor interface {
	// Informs the server that the client will not be using this
	// protocol object anymore. This does not affect any other
	// objects, wl_subsurface objects included.
	Destroy()
	// Create a sub-surface interface for the given surface, and
	// associate it with the given parent surface. This turns a
	// plain wl_surface into a sub-surface.
	// The to-be sub-surface must not already have another role, and it
	// must not have an existing wl_subsurface object. Otherwise a protocol
	// error is raised.
	GetSubsurface(Id WlNewId, Surface WlObject, Parent WlObject)
}
type WlSubcompositorError uint32

const (
	WlSubcompositorBadSurface WlSubcompositorError = 0
)

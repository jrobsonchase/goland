package gen

// This interface is implemented by servers that provide
// desktop-style user interfaces.
// It allows clients to associate a wl_shell_surface with
// a basic surface.
type WlShell interface {
	// Create a shell surface for an existing surface. This gives
	// the wl_surface the role of a shell surface. If the wl_surface
	// already has another role, it raises a protocol error.
	// Only one shell surface can be associated with a given surface.
	GetShellSurface(Id WlNewId, Surface WlObject)
}
type WlShellError uint32

const (
	WlShellRole WlShellError = 0
)

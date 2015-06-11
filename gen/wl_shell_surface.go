package gen

// An interface that may be implemented by a wl_surface, for
// implementations that provide a desktop-style user interface.
// It provides requests to treat surfaces like toplevel, fullscreen
// or popup windows, move, resize or maximize them, associate
// metadata like title and class, etc.
// On the server side the object is automatically destroyed when
// the related wl_surface is destroyed.  On client side,
// wl_shell_surface_destroy() must be called before destroying
// the wl_surface object.
type WlShellSurface interface {
	// Ping a client to check if it is receiving events and sending
	// requests. A client is expected to reply with a pong request.
	Ping(Serial WlUint)
	// The configure event asks the client to resize its surface.
	// The size is a hint, in the sense that the client is free to
	// ignore it if it doesn't resize, pick a smaller size (to
	// satisfy aspect ratio or resize in steps of NxM pixels).
	// The edges parameter provides a hint about how the surface
	// was resized. The client may use this information to decide
	// how to adjust its content to the new size (e.g. a scrolling
	// area might adjust its content position to leave the viewable
	// content unmoved).
	// The client is free to dismiss all but the last configure
	// event it received.
	// The width and height arguments specify the size of the window
	// in surface local coordinates.
	Configure(Edges WlUint, Width WlInt, Height WlInt)
	// The popup_done event is sent out when a popup grab is broken,
	// that is, when the user clicks a surface that doesn't belong
	// to the client owning the popup surface.
	PopupDone()
	// A client must respond to a ping event with a pong request or
	// the client may be deemed unresponsive.
	Pong(Serial WlUint)
	// Start a pointer-driven move of the surface.
	// This request must be used in response to a button press event.
	// The server may ignore move requests depending on the state of
	// the surface (e.g. fullscreen or maximized).
	Move(Seat WlObject, Serial WlUint)
	// Start a pointer-driven resizing of the surface.
	// This request must be used in response to a button press event.
	// The server may ignore resize requests depending on the state of
	// the surface (e.g. fullscreen or maximized).
	Resize(Seat WlObject, Serial WlUint, Edges WlUint)
	// Map the surface as a toplevel surface.
	// A toplevel surface is not fullscreen, maximized or transient.
	SetToplevel()
	// Map the surface relative to an existing surface.
	// The x and y arguments specify the locations of the upper left
	// corner of the surface relative to the upper left corner of the
	// parent surface, in surface local coordinates.
	// The flags argument controls details of the transient behaviour.
	SetTransient(Parent WlObject, X WlInt, Y WlInt, Flags WlUint)
	// Map the surface as a fullscreen surface.
	// If an output parameter is given then the surface will be made
	// fullscreen on that output. If the client does not specify the
	// output then the compositor will apply its policy - usually
	// choosing the output on which the surface has the biggest surface
	// area.
	// The client may specify a method to resolve a size conflict
	// between the output size and the surface size - this is provided
	// through the method parameter.
	// The framerate parameter is used only when the method is set
	// to "driver", to indicate the preferred framerate. A value of 0
	// indicates that the app does not care about framerate.  The
	// framerate is specified in mHz, that is framerate of 60000 is 60Hz.
	// A method of "scale" or "driver" implies a scaling operation of
	// the surface, either via a direct scaling operation or a change of
	// the output mode. This will override any kind of output scaling, so
	// that mapping a surface with a buffer size equal to the mode can
	// fill the screen independent of buffer_scale.
	// A method of "fill" means we don't scale up the buffer, however
	// any output scale is applied. This means that you may run into
	// an edge case where the application maps a buffer with the same
	// size of the output mode but buffer_scale 1 (thus making a
	// surface larger than the output). In this case it is allowed to
	// downscale the results to fit the screen.
	// The compositor must reply to this request with a configure event
	// with the dimensions for the output on which the surface will
	// be made fullscreen.
	SetFullscreen(Method WlUint, Framerate WlUint, Output WlObject)
	// Map the surface as a popup.
	// A popup surface is a transient surface with an added pointer
	// grab.
	// An existing implicit grab will be changed to owner-events mode,
	// and the popup grab will continue after the implicit grab ends
	// (i.e. releasing the mouse button does not cause the popup to
	// be unmapped).
	// The popup grab continues until the window is destroyed or a
	// mouse button is pressed in any other clients window. A click
	// in any of the clients surfaces is reported as normal, however,
	// clicks in other clients surfaces will be discarded and trigger
	// the callback.
	// The x and y arguments specify the locations of the upper left
	// corner of the surface relative to the upper left corner of the
	// parent surface, in surface local coordinates.
	SetPopup(Seat WlObject, Serial WlUint, Parent WlObject, X WlInt, Y WlInt, Flags WlUint)
	// Map the surface as a maximized surface.
	// If an output parameter is given then the surface will be
	// maximized on that output. If the client does not specify the
	// output then the compositor will apply its policy - usually
	// choosing the output on which the surface has the biggest surface
	// area.
	// The compositor will reply with a configure event telling
	// the expected new surface size. The operation is completed
	// on the next buffer attach to this surface.
	// A maximized surface typically fills the entire output it is
	// bound to, except for desktop element such as panels. This is
	// the main difference between a maximized shell surface and a
	// fullscreen shell surface.
	// The details depend on the compositor implementation.
	SetMaximized(Output WlObject)
	// Set a short title for the surface.
	// This string may be used to identify the surface in a task bar,
	// window list, or other user interface elements provided by the
	// compositor.
	// The string must be encoded in UTF-8.
	SetTitle(Title WlString)
	// Set a class for the surface.
	// The surface class identifies the general class of applications
	// to which the surface belongs. A common convention is to use the
	// file name (or the full path if it is a non-standard location) of
	// the application's .desktop file as the class.
	SetClass(Class WlString)
}

// These values are used to indicate which edge of a surface
// is being dragged in a resize operation. The server may
// use this information to adapt its behavior, e.g. choose
// an appropriate cursor image.
type WlShellSurfaceResize uint32

const (
	WlShellSurfaceNone        WlShellSurfaceResize = 0
	WlShellSurfaceTop         WlShellSurfaceResize = 1
	WlShellSurfaceBottom      WlShellSurfaceResize = 2
	WlShellSurfaceLeft        WlShellSurfaceResize = 4
	WlShellSurfaceTopLeft     WlShellSurfaceResize = 5
	WlShellSurfaceBottomLeft  WlShellSurfaceResize = 6
	WlShellSurfaceRight       WlShellSurfaceResize = 8
	WlShellSurfaceTopRight    WlShellSurfaceResize = 9
	WlShellSurfaceBottomRight WlShellSurfaceResize = 10
)

// These flags specify details of the expected behaviour
// of transient surfaces. Used in the set_transient request.
type WlShellSurfaceTransient uint32

const (
	WlShellSurfaceInactive WlShellSurfaceTransient = 0x1
)

// Hints to indicate to the compositor how to deal with a conflict
// between the dimensions of the surface and the dimensions of the
// output. The compositor is free to ignore this parameter.
type WlShellSurfaceFullscreenMethod uint32

const (
	WlShellSurfaceDefault WlShellSurfaceFullscreenMethod = 0
	WlShellSurfaceScale   WlShellSurfaceFullscreenMethod = 1
	WlShellSurfaceDriver  WlShellSurfaceFullscreenMethod = 2
	WlShellSurfaceFill    WlShellSurfaceFullscreenMethod = 3
)

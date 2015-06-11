package gen

// The wl_pointer interface represents one or more input devices,
// such as mice, which control the pointer location and pointer_focus
// of a seat.
// The wl_pointer interface generates motion, enter and leave
// events for the surfaces that the pointer is located over,
// and button and axis events for button presses, button releases
// and scrolling.
type WlPointer interface {
	// Notification that this seat's pointer is focused on a certain
	// surface.
	// When an seat's focus enters a surface, the pointer image
	// is undefined and a client should respond to this event by setting
	// an appropriate pointer image with the set_cursor request.
	Enter(Serial WlUint, Surface WlObject, SurfaceX WlFixed, SurfaceY WlFixed)
	// Notification that this seat's pointer is no longer focused on
	// a certain surface.
	// The leave notification is sent before the enter notification
	// for the new focus.
	Leave(Serial WlUint, Surface WlObject)
	// Notification of pointer location change. The arguments
	// surface_x and surface_y are the location relative to the
	// focused surface.
	Motion(Time WlUint, SurfaceX WlFixed, SurfaceY WlFixed)
	// Mouse button click and release notifications.
	// The location of the click is given by the last motion or
	// enter event.
	// The time argument is a timestamp with millisecond
	// granularity, with an undefined base.
	Button(Serial WlUint, Time WlUint, Button WlUint, State WlUint)
	// Scroll and other axis notifications.
	// For scroll events (vertical and horizontal scroll axes), the
	// value parameter is the length of a vector along the specified
	// axis in a coordinate space identical to those of motion events,
	// representing a relative movement along the specified axis.
	// For devices that support movements non-parallel to axes multiple
	// axis events will be emitted.
	// When applicable, for example for touch pads, the server can
	// choose to emit scroll events where the motion vector is
	// equivalent to a motion event vector.
	// When applicable, clients can transform its view relative to the
	// scroll distance.
	Axis(Time WlUint, Axis WlUint, Value WlFixed)
	// Set the pointer surface, i.e., the surface that contains the
	// pointer image (cursor). This request gives the surface the role
	// of a cursor. If the surface already has another role, it raises
	// a protocol error.
	// The cursor actually changes only if the pointer
	// focus for this device is one of the requesting client's surfaces
	// or the surface parameter is the current pointer surface. If
	// there was a previous surface set with this request it is
	// replaced. If surface is NULL, the pointer image is hidden.
	// The parameters hotspot_x and hotspot_y define the position of
	// the pointer surface relative to the pointer location. Its
	// top-left corner is always at (x, y) - (hotspot_x, hotspot_y),
	// where (x, y) are the coordinates of the pointer location, in surface
	// local coordinates.
	// On surface.attach requests to the pointer surface, hotspot_x
	// and hotspot_y are decremented by the x and y parameters
	// passed to the request. Attach must be confirmed by
	// wl_surface.commit as usual.
	// The hotspot can also be updated by passing the currently set
	// pointer surface to this request with new values for hotspot_x
	// and hotspot_y.
	// The current and pending input regions of the wl_surface are
	// cleared, and wl_surface.set_input_region is ignored until the
	// wl_surface is no longer used as the cursor. When the use as a
	// cursor ends, the current and pending input regions become
	// undefined, and the wl_surface is unmapped.
	SetCursor(Serial WlUint, Surface WlObject, HotspotX WlInt, HotspotY WlInt)
	// Using this request client can tell the server that it is not going to
	// use the pointer object anymore.
	// This request destroys the pointer proxy object, so user must not call
	// wl_pointer_destroy() after using this request.
	Release()
}
type WlPointerError uint32

const (
	WlPointerRole WlPointerError = 0
)

// Describes the physical state of a button which provoked the button
// event.
type WlPointerButtonState uint32

const (
	WlPointerReleased WlPointerButtonState = 0
	WlPointerPressed  WlPointerButtonState = 1
)

// Describes the axis types of scroll events.
type WlPointerAxis uint32

const (
	WlPointerVerticalScroll   WlPointerAxis = 0
	WlPointerHorizontalScroll WlPointerAxis = 1
)

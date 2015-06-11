package gen

// An output describes part of the compositor geometry.  The
// compositor works in the 'compositor coordinate system' and an
// output corresponds to rectangular area in that space that is
// actually visible.  This typically corresponds to a monitor that
// displays part of the compositor space.  This object is published
// as global during start up, or when a monitor is hotplugged.
type WlOutput interface {
	// The geometry event describes geometric properties of the output.
	// The event is sent when binding to the output object and whenever
	// any of the properties change.
	Geometry(X WlInt, Y WlInt, PhysicalWidth WlInt, PhysicalHeight WlInt, Subpixel WlInt, Make WlString, Model WlString, Transform WlInt)
	// The mode event describes an available mode for the output.
	// The event is sent when binding to the output object and there
	// will always be one mode, the current mode.  The event is sent
	// again if an output changes mode, for the mode that is now
	// current.  In other words, the current mode is always the last
	// mode that was received with the current flag set.
	// The size of a mode is given in physical hardware units of
	// the output device. This is not necessarily the same as
	// the output size in the global compositor space. For instance,
	// the output may be scaled, as described in wl_output.scale,
	// or transformed , as described in wl_output.transform.
	Mode(Flags WlUint, Width WlInt, Height WlInt, Refresh WlInt)
	// This event is sent after all other properties has been
	// sent after binding to the output object and after any
	// other property changes done after that. This allows
	// changes to the output properties to be seen as
	// atomic, even if they happen via multiple events.
	Done()
	// This event contains scaling geometry information
	// that is not in the geometry event. It may be sent after
	// binding the output object or if the output scale changes
	// later. If it is not sent, the client should assume a
	// scale of 1.
	// A scale larger than 1 means that the compositor will
	// automatically scale surface buffers by this amount
	// when rendering. This is used for very high resolution
	// displays where applications rendering at the native
	// resolution would be too small to be legible.
	// It is intended that scaling aware clients track the
	// current output of a surface, and if it is on a scaled
	// output it should use wl_surface.set_buffer_scale with
	// the scale of the output. That way the compositor can
	// avoid scaling the surface, and the client can supply
	// a higher detail image.
	Scale(Factor WlInt)
}

// This enumeration describes how the physical
// pixels on an output are laid out.
type WlOutputSubpixel uint32

const (
	WlOutputUnknown       WlOutputSubpixel = 0
	WlOutputNone          WlOutputSubpixel = 1
	WlOutputHorizontalRgb WlOutputSubpixel = 2
	WlOutputHorizontalBgr WlOutputSubpixel = 3
	WlOutputVerticalRgb   WlOutputSubpixel = 4
	WlOutputVerticalBgr   WlOutputSubpixel = 5
)

// This describes the transform that a compositor will apply to a
// surface to compensate for the rotation or mirroring of an
// output device.
// The flipped values correspond to an initial flip around a
// vertical axis followed by rotation.
// The purpose is mainly to allow clients render accordingly and
// tell the compositor, so that for fullscreen surfaces, the
// compositor will still be able to scan out directly from client
// surfaces.
type WlOutputTransform uint32

const (
	WlOutputNormal     WlOutputTransform = 0
	WlOutput90         WlOutputTransform = 1
	WlOutput180        WlOutputTransform = 2
	WlOutput270        WlOutputTransform = 3
	WlOutputFlipped    WlOutputTransform = 4
	WlOutputFlipped90  WlOutputTransform = 5
	WlOutputFlipped180 WlOutputTransform = 6
	WlOutputFlipped270 WlOutputTransform = 7
)

// These flags describe properties of an output mode.
// They are used in the flags bitfield of the mode event.
type WlOutputMode uint32

const (
	WlOutputCurrent   WlOutputMode = 0x1
	WlOutputPreferred WlOutputMode = 0x2
)

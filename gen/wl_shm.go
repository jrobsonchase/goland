package gen

// A global singleton object that provides support for shared
// memory.
// Clients can create wl_shm_pool objects using the create_pool
// request.
// At connection setup time, the wl_shm object emits one or more
// format events to inform clients about the valid pixel formats
// that can be used for buffers.
type WlShm interface {
	// Informs the client about a valid pixel format that
	// can be used for buffers. Known formats include
	// argb8888 and xrgb8888.
	Format(Format WlUint)
	// Create a new wl_shm_pool object.
	// The pool can be used to create shared memory based buffer
	// objects.  The server will mmap size bytes of the passed file
	// descriptor, to use as backing memory for the pool.
	CreatePool(Id WlNewId, Fd WlFd, Size WlInt)
}

// These errors can be emitted in response to wl_shm requests.
type WlShmError uint32

const (
	WlShmInvalidFormat WlShmError = 0
	WlShmInvalidStride WlShmError = 1
	WlShmInvalidFd     WlShmError = 2
)

// This describes the memory layout of an individual pixel.
// All renderers should support argb8888 and xrgb8888 but any other
// formats are optional and may not be supported by the particular
// renderer in use.
type WlShmFormat uint32

const (
	WlShmArgb8888    WlShmFormat = 0
	WlShmXrgb8888    WlShmFormat = 1
	WlShmC8          WlShmFormat = 0x20203843
	WlShmRgb332      WlShmFormat = 0x38424752
	WlShmBgr233      WlShmFormat = 0x38524742
	WlShmXrgb4444    WlShmFormat = 0x32315258
	WlShmXbgr4444    WlShmFormat = 0x32314258
	WlShmRgbx4444    WlShmFormat = 0x32315852
	WlShmBgrx4444    WlShmFormat = 0x32315842
	WlShmArgb4444    WlShmFormat = 0x32315241
	WlShmAbgr4444    WlShmFormat = 0x32314241
	WlShmRgba4444    WlShmFormat = 0x32314152
	WlShmBgra4444    WlShmFormat = 0x32314142
	WlShmXrgb1555    WlShmFormat = 0x35315258
	WlShmXbgr1555    WlShmFormat = 0x35314258
	WlShmRgbx5551    WlShmFormat = 0x35315852
	WlShmBgrx5551    WlShmFormat = 0x35315842
	WlShmArgb1555    WlShmFormat = 0x35315241
	WlShmAbgr1555    WlShmFormat = 0x35314241
	WlShmRgba5551    WlShmFormat = 0x35314152
	WlShmBgra5551    WlShmFormat = 0x35314142
	WlShmRgb565      WlShmFormat = 0x36314752
	WlShmBgr565      WlShmFormat = 0x36314742
	WlShmRgb888      WlShmFormat = 0x34324752
	WlShmBgr888      WlShmFormat = 0x34324742
	WlShmXbgr8888    WlShmFormat = 0x34324258
	WlShmRgbx8888    WlShmFormat = 0x34325852
	WlShmBgrx8888    WlShmFormat = 0x34325842
	WlShmAbgr8888    WlShmFormat = 0x34324241
	WlShmRgba8888    WlShmFormat = 0x34324152
	WlShmBgra8888    WlShmFormat = 0x34324142
	WlShmXrgb2101010 WlShmFormat = 0x30335258
	WlShmXbgr2101010 WlShmFormat = 0x30334258
	WlShmRgbx1010102 WlShmFormat = 0x30335852
	WlShmBgrx1010102 WlShmFormat = 0x30335842
	WlShmArgb2101010 WlShmFormat = 0x30335241
	WlShmAbgr2101010 WlShmFormat = 0x30334241
	WlShmRgba1010102 WlShmFormat = 0x30334152
	WlShmBgra1010102 WlShmFormat = 0x30334142
	WlShmYuyv        WlShmFormat = 0x56595559
	WlShmYvyu        WlShmFormat = 0x55595659
	WlShmUyvy        WlShmFormat = 0x59565955
	WlShmVyuy        WlShmFormat = 0x59555956
	WlShmAyuv        WlShmFormat = 0x56555941
	WlShmNv12        WlShmFormat = 0x3231564e
	WlShmNv21        WlShmFormat = 0x3132564e
	WlShmNv16        WlShmFormat = 0x3631564e
	WlShmNv61        WlShmFormat = 0x3136564e
	WlShmYuv410      WlShmFormat = 0x39565559
	WlShmYvu410      WlShmFormat = 0x39555659
	WlShmYuv411      WlShmFormat = 0x31315559
	WlShmYvu411      WlShmFormat = 0x31315659
	WlShmYuv420      WlShmFormat = 0x32315559
	WlShmYvu420      WlShmFormat = 0x32315659
	WlShmYuv422      WlShmFormat = 0x36315559
	WlShmYvu422      WlShmFormat = 0x36315659
	WlShmYuv444      WlShmFormat = 0x34325559
	WlShmYvu444      WlShmFormat = 0x34325659
)

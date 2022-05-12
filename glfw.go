package glfw

/*
#include <GLFW/glfw3.h>
*/
import "C"

// Version
const (
	VersionMajor    int = C.GLFW_VERSION_MAJOR
	VersionMinor    int = C.GLFW_VERSION_MINOR
	VersionRevision int = C.GLFW_VERSION_REVISION
)

const (
	True  = C.GLFW_TRUE
	False = C.GLFW_FALSE
)

// Key and button actions
const (
	Release = C.GLFW_RELEASE
	Press   = C.GLFW_PRESS
	Repeat  = C.GLFW_REPEAT
)

// Joystick hat states
const (
	HatCentered  = C.GLFW_HAT_CENTERED
	HatUp        = C.GLFW_HAT_UP
	HatRight     = C.GLFW_HAT_RIGHT
	HatDown      = C.GLFW_HAT_DOWN
	HatLeft      = C.GLFW_HAT_LEFT
	HatRightUp   = C.GLFW_HAT_RIGHT_UP
	HatRightDown = C.GLFW_HAT_RIGHT_DOWN
	HatLeftUp    = C.GLFW_HAT_LEFT_UP
	HatLeftDown  = C.GLFW_HAT_LEFT_DOWN
)

// The unknown key

const KeyUnknown = C.GLFW_KEY_UNKNOWN

// Printable keys
const (
	KeySpace        = C.GLFW_KEY_SPACE
	KeyApostrophe   = C.GLFW_KEY_APOSTROPHE
	KeyComma        = C.GLFW_KEY_COMMA
	KeyMinus        = C.GLFW_KEY_MINUS
	KeyPeriod       = C.GLFW_KEY_PERIOD
	KeySlash        = C.GLFW_KEY_SLASH
	Key0            = C.GLFW_KEY_0
	Key1            = C.GLFW_KEY_1
	Key2            = C.GLFW_KEY_2
	Key3            = C.GLFW_KEY_3
	Key4            = C.GLFW_KEY_4
	Key5            = C.GLFW_KEY_5
	Key6            = C.GLFW_KEY_6
	Key7            = C.GLFW_KEY_7
	Key8            = C.GLFW_KEY_8
	Key9            = C.GLFW_KEY_9
	KeySemicolon    = C.GLFW_KEY_SEMICOLON
	KeyEqual        = C.GLFW_KEY_EQUAL
	KeyA            = C.GLFW_KEY_A
	KeyB            = C.GLFW_KEY_B
	KeyC            = C.GLFW_KEY_C
	KeyD            = C.GLFW_KEY_D
	KeyE            = C.GLFW_KEY_E
	KeyF            = C.GLFW_KEY_F
	KeyG            = C.GLFW_KEY_G
	KeyH            = C.GLFW_KEY_H
	KeyI            = C.GLFW_KEY_I
	KeyJ            = C.GLFW_KEY_J
	KeyK            = C.GLFW_KEY_K
	KeyL            = C.GLFW_KEY_L
	KeyM            = C.GLFW_KEY_M
	KeyN            = C.GLFW_KEY_N
	KeyO            = C.GLFW_KEY_O
	KeyP            = C.GLFW_KEY_P
	KeyQ            = C.GLFW_KEY_Q
	KeyR            = C.GLFW_KEY_R
	KeyS            = C.GLFW_KEY_S
	KeyT            = C.GLFW_KEY_T
	KeyU            = C.GLFW_KEY_U
	KeyV            = C.GLFW_KEY_V
	KeyW            = C.GLFW_KEY_W
	KeyX            = C.GLFW_KEY_X
	KeyY            = C.GLFW_KEY_Y
	KeyZ            = C.GLFW_KEY_Z
	KeyLeftBracket  = C.GLFW_KEY_LEFT_BRACKET
	KeyBackslash    = C.GLFW_KEY_BACKSLASH
	KeyRightBracket = C.GLFW_KEY_RIGHT_BRACKET
	KeyGraveAccent  = C.GLFW_KEY_GRAVE_ACCENT
	KeyWorld1       = C.GLFW_KEY_WORLD_1
	KeyWorld2       = C.GLFW_KEY_WORLD_2
)

// Function keys
const (
	KeyEscape       = C.GLFW_KEY_ESCAPE
	KeyEnter        = C.GLFW_KEY_ENTER
	KeyTab          = C.GLFW_KEY_TAB
	KeyBackspace    = C.GLFW_KEY_BACKSPACE
	KeyInsert       = C.GLFW_KEY_INSERT
	KeyDelete       = C.GLFW_KEY_DELETE
	KeyRight        = C.GLFW_KEY_RIGHT
	KeyLeft         = C.GLFW_KEY_LEFT
	KeyDown         = C.GLFW_KEY_DOWN
	KeyUp           = C.GLFW_KEY_UP
	KeyPageUp       = C.GLFW_KEY_PAGE_UP
	KeyPageDown     = C.GLFW_KEY_PAGE_DOWN
	KeyHome         = C.GLFW_KEY_HOME
	KeyEnd          = C.GLFW_KEY_END
	KeyCapsLock     = C.GLFW_KEY_CAPS_LOCK
	KeyScrollLock   = C.GLFW_KEY_SCROLL_LOCK
	KeyNumLock      = C.GLFW_KEY_NUM_LOCK
	KeyPrintScreen  = C.GLFW_KEY_PRINT_SCREEN
	KeyPause        = C.GLFW_KEY_PAUSE
	KeyF1           = C.GLFW_KEY_F1
	KeyF2           = C.GLFW_KEY_F2
	KeyF3           = C.GLFW_KEY_F3
	KeyF4           = C.GLFW_KEY_F4
	KeyF5           = C.GLFW_KEY_F5
	KeyF6           = C.GLFW_KEY_F6
	KeyF7           = C.GLFW_KEY_F7
	KeyF8           = C.GLFW_KEY_F8
	KeyF9           = C.GLFW_KEY_F9
	KeyF10          = C.GLFW_KEY_F10
	KeyF11          = C.GLFW_KEY_F11
	KeyF12          = C.GLFW_KEY_F12
	KeyF13          = C.GLFW_KEY_F13
	KeyF14          = C.GLFW_KEY_F14
	KeyF15          = C.GLFW_KEY_F15
	KeyF16          = C.GLFW_KEY_F16
	KeyF17          = C.GLFW_KEY_F17
	KeyF18          = C.GLFW_KEY_F18
	KeyF19          = C.GLFW_KEY_F19
	KeyF20          = C.GLFW_KEY_F20
	KeyF21          = C.GLFW_KEY_F21
	KeyF22          = C.GLFW_KEY_F22
	KeyF23          = C.GLFW_KEY_F23
	KeyF24          = C.GLFW_KEY_F24
	KeyF25          = C.GLFW_KEY_F25
	KeyKp0          = C.GLFW_KEY_KP_0
	KeyKp1          = C.GLFW_KEY_KP_1
	KeyKp2          = C.GLFW_KEY_KP_2
	KeyKp3          = C.GLFW_KEY_KP_3
	KeyKp4          = C.GLFW_KEY_KP_4
	KeyKp5          = C.GLFW_KEY_KP_5
	KeyKp6          = C.GLFW_KEY_KP_6
	KeyKp7          = C.GLFW_KEY_KP_7
	KeyKp8          = C.GLFW_KEY_KP_8
	KeyKp9          = C.GLFW_KEY_KP_9
	KeyKpDecimal    = C.GLFW_KEY_KP_DECIMAL
	KeyKpDivide     = C.GLFW_KEY_KP_DIVIDE
	KeyKpMultiply   = C.GLFW_KEY_KP_MULTIPLY
	KeyKpSubtract   = C.GLFW_KEY_KP_SUBTRACT
	KeyKpAdd        = C.GLFW_KEY_KP_ADD
	KeyKpEnter      = C.GLFW_KEY_KP_ENTER
	KeyKpEqual      = C.GLFW_KEY_KP_EQUAL
	KeyLeftShift    = C.GLFW_KEY_LEFT_SHIFT
	KeyLeftControl  = C.GLFW_KEY_LEFT_CONTROL
	KeyLeftAlt      = C.GLFW_KEY_LEFT_ALT
	KeyLeftSuper    = C.GLFW_KEY_LEFT_SUPER
	KeyRightShift   = C.GLFW_KEY_RIGHT_SHIFT
	KeyRightControl = C.GLFW_KEY_RIGHT_CONTROL
	KeyRightAlt     = C.GLFW_KEY_RIGHT_ALT
	KeyRightSuper   = C.GLFW_KEY_RIGHT_SUPER
	KeyMenu         = C.GLFW_KEY_MENU
	KeyLast         = C.GLFW_KEY_LAST
)

// Modifier key flags
const (
	ModShift   = C.GLFW_MOD_SHIFT
	ModControl = C.GLFW_MOD_CONTROL
	ModAlt     = C.GLFW_MOD_ALT
	ModSuper   = C.GLFW_MOD_SUPER
	ModCapsLoc = C.GLFW_MOD_CAPS_LOCK
	ModNumLock = C.GLFW_MOD_NUM_LOCK
)

// Mouse buttons
const (
	MouseButton1      = C.GLFW_MOUSE_BUTTON_1
	MouseButton2      = C.GLFW_MOUSE_BUTTON_2
	MouseButton3      = C.GLFW_MOUSE_BUTTON_3
	MouseButton4      = C.GLFW_MOUSE_BUTTON_4
	MouseButton5      = C.GLFW_MOUSE_BUTTON_5
	MouseButton6      = C.GLFW_MOUSE_BUTTON_6
	MouseButton7      = C.GLFW_MOUSE_BUTTON_7
	MouseButton8      = C.GLFW_MOUSE_BUTTON_8
	MouseButtonLast   = C.GLFW_MOUSE_BUTTON_LAST
	MouseButtonLeft   = C.GLFW_MOUSE_BUTTON_LEFT
	MouseButtonRight  = C.GLFW_MOUSE_BUTTON_RIGHT
	MouseButtonMiddle = C.GLFW_MOUSE_BUTTON_MIDDLE
)

// Joysticks
const (
	Joystick1    = C.GLFW_JOYSTICK_1
	Joystick2    = C.GLFW_JOYSTICK_2
	Joystick3    = C.GLFW_JOYSTICK_3
	Joystick4    = C.GLFW_JOYSTICK_4
	Joystick5    = C.GLFW_JOYSTICK_5
	Joystick6    = C.GLFW_JOYSTICK_6
	Joystick7    = C.GLFW_JOYSTICK_7
	Joystick8    = C.GLFW_JOYSTICK_8
	Joystick9    = C.GLFW_JOYSTICK_9
	Joystick10   = C.GLFW_JOYSTICK_10
	Joystick11   = C.GLFW_JOYSTICK_11
	Joystick12   = C.GLFW_JOYSTICK_12
	Joystick13   = C.GLFW_JOYSTICK_13
	Joystick14   = C.GLFW_JOYSTICK_14
	Joystick15   = C.GLFW_JOYSTICK_15
	Joystick16   = C.GLFW_JOYSTICK_16
	JoystickLast = C.GLFW_JOYSTICK_LAST
)

// Gamepad buttons
const (
	GamepadButtonA           = C.GLFW_GAMEPAD_BUTTON_A
	GamepadButtonB           = C.GLFW_GAMEPAD_BUTTON_B
	GamepadButtonX           = C.GLFW_GAMEPAD_BUTTON_X
	GamepadButtonY           = C.GLFW_GAMEPAD_BUTTON_Y
	GamepadButtonLeftBumper  = C.GLFW_GAMEPAD_BUTTON_LEFT_BUMPER
	GamepadButtonRightBumper = C.GLFW_GAMEPAD_BUTTON_RIGHT_BUMPER
	GamepadButtonBack        = C.GLFW_GAMEPAD_BUTTON_BACK
	GamepadButtonStart       = C.GLFW_GAMEPAD_BUTTON_START
	GamepadButtonGuide       = C.GLFW_GAMEPAD_BUTTON_GUIDE
	GamepadButtonLeftThumb   = C.GLFW_GAMEPAD_BUTTON_LEFT_THUMB
	GamepadButtonRightThumb  = C.GLFW_GAMEPAD_BUTTON_RIGHT_THUMB
	GamepadButtonDPadUp      = C.GLFW_GAMEPAD_BUTTON_DPAD_UP
	GamepadButtonDPadRight   = C.GLFW_GAMEPAD_BUTTON_DPAD_RIGHT
	GamepadButtonDPadDown    = C.GLFW_GAMEPAD_BUTTON_DPAD_DOWN
	GamepadButtonDPadLeft    = C.GLFW_GAMEPAD_BUTTON_DPAD_LEFT
	GamepadButtonLast        = C.GLFW_GAMEPAD_BUTTON_LAST
	GamepadButtonCross       = C.GLFW_GAMEPAD_BUTTON_CROSS
	GamepadButtonCircle      = C.GLFW_GAMEPAD_BUTTON_CIRCLE
	GamepadButtonSquare      = C.GLFW_GAMEPAD_BUTTON_SQUARE
	GamepadButtonTriangle    = C.GLFW_GAMEPAD_BUTTON_TRIANGLE
)

// Gamepad axes
const (
	GamepadAxisLeftX        = C.GLFW_GAMEPAD_AXIS_LEFT_X
	GamepadAxisLeftY        = C.GLFW_GAMEPAD_AXIS_LEFT_Y
	GamepadAxisRightX       = C.GLFW_GAMEPAD_AXIS_RIGHT_X
	GamepadAxisRightY       = C.GLFW_GAMEPAD_AXIS_RIGHT_Y
	GamepadAxisLeftTrigger  = C.GLFW_GAMEPAD_AXIS_LEFT_TRIGGER
	GamepadAxisRightTrigger = C.GLFW_GAMEPAD_AXIS_RIGHT_TRIGGER
	GamepadAxisLast         = C.GLFW_GAMEPAD_AXIS_LAST
)

// Error codes
const (
	NoError            = C.GLFW_NO_ERROR
	NotInitialized     = C.GLFW_NOT_INITIALIZED
	NoCurrentContext   = C.GLFW_NO_CURRENT_CONTEXT
	InvalidEnum        = C.GLFW_INVALID_ENUM
	InvalidValue       = C.GLFW_INVALID_VALUE
	OutOfMemory        = C.GLFW_OUT_OF_MEMORY
	ApiUnavailable     = C.GLFW_API_UNAVAILABLE
	VersionUnavailable = C.GLFW_VERSION_UNAVAILABLE
	PlatformError      = C.GLFW_PLATFORM_ERROR
	FormatUnavailable  = C.GLFW_FORMAT_UNAVAILABLE
	NoWindowContext    = C.GLFW_NO_WINDOW_CONTEXT
)

// Input focus window hint and attribute
const (
	Focused                = C.GLFW_FOCUSED
	Iconified              = C.GLFW_ICONIFIED
	Resizable              = C.GLFW_RESIZABLE
	Visible                = C.GLFW_VISIBLE
	Decorated              = C.GLFW_DECORATED
	AutoIconify            = C.GLFW_AUTO_ICONIFY
	Floating               = C.GLFW_FLOATING
	Maximized              = C.GLFW_MAXIMIZED
	CenterCursor           = C.GLFW_CENTER_CURSOR
	TransparentFramebuffer = C.GLFW_TRANSPARENT_FRAMEBUFFER
	Hovered                = C.GLFW_HOVERED
	FocusOnShow            = C.GLFW_FOCUS_ON_SHOW
	RedBits                = C.GLFW_RED_BITS
	GreenBits              = C.GLFW_GREEN_BITS
	BlueBits               = C.GLFW_BLUE_BITS
	AlphaBits              = C.GLFW_ALPHA_BITS
	DepthBits              = C.GLFW_DEPTH_BITS
	StencilBits            = C.GLFW_STENCIL_BITS
	AccumRedBits           = C.GLFW_ACCUM_RED_BITS
	AccumGreenBits         = C.GLFW_ACCUM_GREEN_BITS
	AccumBlueBits          = C.GLFW_ACCUM_BLUE_BITS
	AccumAlphaBits         = C.GLFW_ACCUM_ALPHA_BITS
	AuxBuffers             = C.GLFW_AUX_BUFFERS
	Stereo                 = C.GLFW_STEREO
	Samples                = C.GLFW_SAMPLES
	SrgbCapable            = C.GLFW_SRGB_CAPABLE
	RefreshRate            = C.GLFW_REFRESH_RATE
	DoubleBuffer           = C.GLFW_DOUBLEBUFFER
	ClientApi              = C.GLFW_CLIENT_API
	ContextVersionMajor    = C.GLFW_CONTEXT_VERSION_MAJOR
	ContextVersionMinor    = C.GLFW_CONTEXT_VERSION_MINOR
	ContextRevision        = C.GLFW_CONTEXT_REVISION
	ContextRobustness      = C.GLFW_CONTEXT_ROBUSTNESS
	OpenglForwardCompat    = C.GLFW_OPENGL_FORWARD_COMPAT
	OpenglDebugContext     = C.GLFW_OPENGL_DEBUG_CONTEXT
	OpenglProfile          = C.GLFW_OPENGL_PROFILE
	ContextReleaseBehavior = C.GLFW_CONTEXT_RELEASE_BEHAVIOR
	ContextNoError         = C.GLFW_CONTEXT_NO_ERROR
	ContextCreationApi     = C.GLFW_CONTEXT_CREATION_API
	ScaleToMonitor         = C.GLFW_SCALE_TO_MONITOR
	CocoaRetinaFramebuffer = C.GLFW_COCOA_RETINA_FRAMEBUFFER
	CocoaFrameName         = C.GLFW_COCOA_FRAME_NAME
	CocoaGraphicsSwitching = C.GLFW_COCOA_GRAPHICS_SWITCHING
	X11ClassName           = C.GLFW_X11_CLASS_NAME
	X11InstanceName        = C.GLFW_X11_INSTANCE_NAME
)

// Api type
const (
	NoApi       = C.GLFW_NO_API
	OpenglApi   = C.GLFW_OPENGL_API
	OpenglEsApi = C.GLFW_OPENGL_ES_API
)

// Robustness
const (
	NoRobustness        = C.GLFW_NO_ROBUSTNESS
	NoResetNotification = C.GLFW_NO_RESET_NOTIFICATION
	LoseContextOnReset  = C.GLFW_LOSE_CONTEXT_ON_RESET
)

// OpenGL Profile
const (
	OpenglAnyProfile    = C.GLFW_OPENGL_ANY_PROFILE
	OpenglCoreProfile   = C.GLFW_OPENGL_CORE_PROFILE
	OpenglCompatProfile = C.GLFW_OPENGL_COMPAT_PROFILE
)

// Inputs
const (
	Cursor             = C.GLFW_CURSOR
	StickyKeys         = C.GLFW_STICKY_KEYS
	StickyMouseButtons = C.GLFW_STICKY_MOUSE_BUTTONS
	LockKeyMods        = C.GLFW_LOCK_KEY_MODS
	RawMouseMotion     = C.GLFW_RAW_MOUSE_MOTION
)

// Cursor state
const (
	CursorNormal   = C.GLFW_CURSOR_NORMAL
	CursorHidden   = C.GLFW_CURSOR_HIDDEN
	CursorDisabled = C.GLFW_CURSOR_DISABLED
)

// Release behavior
const (
	AnyReleaseBehavior   = C.GLFW_ANY_RELEASE_BEHAVIOR
	ReleaseBehaviorFlush = C.GLFW_RELEASE_BEHAVIOR_FLUSH
	ReleaseBehaviorNone  = C.GLFW_RELEASE_BEHAVIOR_NONE
)

// Context api
const (
	NativeContextApi = C.GLFW_NATIVE_CONTEXT_API
	EglContextApi    = C.GLFW_EGL_CONTEXT_API
	OsmesaContextApi = C.GLFW_OSMESA_CONTEXT_API
)

// Standard cursor shapes
const (
	ArrowCursor     = C.GLFW_ARROW_CURSOR
	IBeamCursor     = C.GLFW_IBEAM_CURSOR
	CrossHairCursor = C.GLFW_CROSSHAIR_CURSOR
	HandCursor      = C.GLFW_HAND_CURSOR
	HResizeCursor   = C.GLFW_HRESIZE_CURSOR
	VResizeCursor   = C.GLFW_VRESIZE_CURSOR
)

// Connectivity
const (
	Connected    = C.GLFW_CONNECTED
	Disconnected = C.GLFW_DISCONNECTED
)

// Joystick hat buttons init hint
const (
	JoystickHatButtons  = C.GLFW_JOYSTICK_HAT_BUTTONS
	CocoaChdirResources = C.GLFW_COCOA_CHDIR_RESOURCES
	CocoaMenubar        = C.GLFW_COCOA_MENUBAR
)

const (
	DontCare = C.GLFW_DONT_CARE
)

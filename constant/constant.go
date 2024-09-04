package constant

var (
	StableVar        = 0.2
	WindowSize       = 8
	StableRollingVar = StableVar * float64(WindowSize)
	MaxThreshold     = 0.4
)
var (
	Sharp         = 7.0
	SlightlySharp = 4.0
	SlightlySlow  = 1.0

	DiscSharp         = "剧烈"
	DiscSlightlySharp = "较剧烈"
	DiscSlow          = "较缓慢"
	DiscSlightlySlow  = "缓慢"
	DiscStable        = "平稳"
)
var (
	Risk0 = 0
	Risk1 = 1
	Risk2 = 2
	Risk3 = 3
	Risk4 = 4

	C0 = 0
	C1 = 1
	C2 = 2
	C3 = 3
)
var (
	JsonName  = "data.json"
	Sheetname = "sheet2"
)

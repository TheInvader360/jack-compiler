package writer

type Segment string

const (
	SegmentConstant Segment = "constant"
	SegmentArgument Segment = "argument"
	SegmentLocal    Segment = "local"
	SegmentStatic   Segment = "static"
	SegmentThis     Segment = "this"
	SegmentThat     Segment = "that"
	SegmentPointer  Segment = "pointer"
	SegmentTemp     Segment = "temp"
	SegmentUnknown  Segment = ""
)

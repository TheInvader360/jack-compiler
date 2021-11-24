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

func GetSegment(segment string) Segment {
	switch segment {
	case "constant":
		return SegmentConstant
	case "argument":
		return SegmentArgument
	case "local":
		return SegmentLocal
	case "static":
		return SegmentStatic
	case "this":
		return SegmentThis
	case "that":
		return SegmentThat
	case "pointer":
		return SegmentPointer
	case "temp":
		return SegmentTemp
	}
	return SegmentUnknown
}

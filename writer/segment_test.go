package writer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSegment(t *testing.T) {
	segment := GetSegment("constant")
	assert.Equal(t, SegmentConstant, segment)
	segment = GetSegment("argument")
	assert.Equal(t, SegmentArgument, segment)
	segment = GetSegment("local")
	assert.Equal(t, SegmentLocal, segment)
	segment = GetSegment("static")
	assert.Equal(t, SegmentStatic, segment)
	segment = GetSegment("this")
	assert.Equal(t, SegmentThis, segment)
	segment = GetSegment("that")
	assert.Equal(t, SegmentThat, segment)
	segment = GetSegment("pointer")
	assert.Equal(t, SegmentPointer, segment)
	segment = GetSegment("temp")
	assert.Equal(t, SegmentTemp, segment)
	segment = GetSegment("invalid")
	assert.Equal(t, SegmentUnknown, segment)
}

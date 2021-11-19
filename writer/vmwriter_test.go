package writer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVMWriter(t *testing.T) {
	vmw := NewVMWriter()
	assert.Equal(t, "", vmw.Code)
}

func TestWritePush(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WritePush(SegmentConstant, 0)
	assert.Equal(t, "push constant 0\n", vmw.Code)
}

func TestWritePop(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WritePop(SegmentTemp, 0)
	assert.Equal(t, "pop temp 0\n", vmw.Code)
}

func TestWriteArithmetic(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WriteArithmetic("+")
	vmw.WriteArithmetic("-")
	vmw.WriteArithmetic("*")
	vmw.WriteArithmetic("/")
	vmw.WriteArithmetic("&")
	vmw.WriteArithmetic("|")
	vmw.WriteArithmetic("<")
	vmw.WriteArithmetic("=")
	vmw.WriteArithmetic(">")
	vmw.WriteArithmetic("neg")
	vmw.WriteArithmetic("not")
	assert.Equal(t, "add\nsub\ncall Math.multiply 2\ncall Math.divide 2\nand\nor\nlt\neq\ngt\nneg\nnot\n", vmw.Code)
}

func TestWriteLabel(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WriteLabel("WHILE_EXP0")
	assert.Equal(t, "label WHILE_EXP0\n", vmw.Code)
}

func TestWriteGoto(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WriteGoto("IF_FALSE0")
	assert.Equal(t, "goto IF_FALSE0\n", vmw.Code)
}

func TestWriteIf(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WriteIf("WHILE_END0")
	assert.Equal(t, "if-goto WHILE_END0\n", vmw.Code)
}

func TestWriteCall(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WriteCall("Math.multiply", 2)
	assert.Equal(t, "call Math.multiply 2\n", vmw.Code)
}

func TestWriteFunction(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WriteFunction("Main.main", 0)
	assert.Equal(t, "function Main.main 0\n", vmw.Code)
}

func TestWriteReturn(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WriteReturn()
	assert.Equal(t, "return\n", vmw.Code)
}

func TestWriteComment(t *testing.T) {
	vmw := NewVMWriter()
	vmw.WriteComment("Comment")
	assert.Equal(t, "// Comment\n", vmw.Code)
}

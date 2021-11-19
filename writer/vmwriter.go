package writer

import "fmt"

type VMWriter struct {
	Code string
}

// NewVMWriter - creates a vm writer
func NewVMWriter() *VMWriter {
	return &VMWriter{}
}

// WritePush - writes a vm push command
func (vmw *VMWriter) WritePush(segment Segment, index int) {
	vmw.Code += fmt.Sprintf("push %s %d\n", segment, index)
}

// WritePop - writes a vm pop command
func (vmw *VMWriter) WritePop(segment Segment, index int) {
	vmw.Code += fmt.Sprintf("pop %s %d\n", segment, index)
}

// WriteArithmetic - writes a vm arithmetic command
func (vmw *VMWriter) WriteArithmetic(command string) {
	switch command {
	case "+":
		command = "add"
	case "-":
		command = "sub"
	case "*":
		command = "call Math.multiply 2"
	case "/":
		command = "call Math.divide 2"
	case "&":
		command = "and"
	case "|":
		command = "or"
	case "<":
		command = "lt"
	case "=":
		command = "eq"
	case ">":
		command = "gt"
	}
	// "neg" and "not" commands remain unaltered
	vmw.Code += fmt.Sprintf("%s\n", command)
}

// WriteLabel - writes a vm label command
func (vmw *VMWriter) WriteLabel(label string) {
	vmw.Code += fmt.Sprintf("label %s\n", label)
}

// WriteGoto - writes a vm goto command
func (vmw *VMWriter) WriteGoto(label string) {
	vmw.Code += fmt.Sprintf("goto %s\n", label)
}

// WriteIf - writes a vm if-goto command
func (vmw *VMWriter) WriteIf(label string) {
	vmw.Code += fmt.Sprintf("if-goto %s\n", label)
}

// WriteCall - writes a vm call command
func (vmw *VMWriter) WriteCall(name string, nArgs int) {
	vmw.Code += fmt.Sprintf("call %s %d\n", name, nArgs)
}

// WriteFunction - writes a vm function command
func (vmw *VMWriter) WriteFunction(name string, nLocals int) {
	vmw.Code += fmt.Sprintf("function %s %d\n", name, nLocals)
}

// WriteReturn - writes a vm return command
func (vmw *VMWriter) WriteReturn() {
	vmw.Code += "return\n"
}

// WriteComment - writes a vm comment
func (vmw *VMWriter) WriteComment(comment string) {
	vmw.Code += fmt.Sprintf("// %s\n", comment)
}

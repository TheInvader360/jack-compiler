package writer

import "fmt"

type VMWriter struct {
	debug bool
	Code  string
}

// NewVMWriter - creates a vm writer
func NewVMWriter(debug bool) *VMWriter {
	return &VMWriter{debug: debug}
}

// WritePush - writes a vm push command
func (vmw *VMWriter) WritePush(segment string, index int) {
	line := fmt.Sprintf("push %s %d", segment, index)
	vmw.Code += line + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WritePush()", line)
	}
}

// WritePop - writes a vm pop command
func (vmw *VMWriter) WritePop(segment string, index int) {
	line := fmt.Sprintf("pop %s %d", segment, index)
	vmw.Code += line + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WritePop()", line)
	}
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
	case "&amp;":
		command = "and"
	case "|":
		command = "or"
	case "&lt;":
		command = "lt"
	case "=":
		command = "eq"
	case "&gt;":
		command = "gt"
	}
	// "neg" and "not" commands remain unaltered
	vmw.Code += command + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WriteArithmetic()", command)
	}
}

// WriteLabel - writes a vm label command
func (vmw *VMWriter) WriteLabel(label string) {
	line := fmt.Sprintf("label %s", label)
	vmw.Code += line + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WriteLabel()", line)
	}
}

// WriteGoto - writes a vm goto command
func (vmw *VMWriter) WriteGoto(label string) {
	line := fmt.Sprintf("goto %s", label)
	vmw.Code += line + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WriteGoto()", line)
	}
}

// WriteIf - writes a vm if-goto command
func (vmw *VMWriter) WriteIf(label string) {
	line := fmt.Sprintf("if-goto %s", label)
	vmw.Code += line + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WriteIf()", line)
	}
}

// WriteCall - writes a vm call command
func (vmw *VMWriter) WriteCall(name string, nArgs int) {
	line := fmt.Sprintf("call %s %d", name, nArgs)
	vmw.Code += line + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WriteCall()", line)
	}
}

// WriteFunction - writes a vm function command
func (vmw *VMWriter) WriteFunction(name string, nLocals int) {
	line := fmt.Sprintf("function %s %d", name, nLocals)
	vmw.Code += line + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WriteFunction()", line)
	}
}

// WriteReturn - writes a vm return command
func (vmw *VMWriter) WriteReturn() {
	line := "return"
	vmw.Code += line + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WriteReturn()", line)
	}
}

// WriteComment - writes a vm comment
func (vmw *VMWriter) WriteComment(comment string) {
	line := fmt.Sprintf("// %s", comment)
	vmw.Code += line + "\n"
	if vmw.debug {
		fmt.Println("VMWriter.WriteComment()", line)
	}
}

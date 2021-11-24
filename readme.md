# Jack Compiler

A compiler that translates Jack programs into executable VM language programs (nand2tetris)

## Usage

```bash
git clone https://github.com/TheInvader360/jack-compiler
cd jack-compiler
```

Then:

```bash
go run main.go -path=examples/Square/Main.jack
```

The translated program is exported to the source directory with the same base filename but a .vm extension (so the given example would generate examples/HelloWorld/Main.vm)

Or:

```bash
go run main.go -path=examples/Square
```

Translates all .jack files in the specified directory and exports them all to the same directory with the same base filename but a .vm extension

# Jack Compiler

A compiler that translates Jack programs into executable VM language programs (nand2tetris)

## Usage

```bash
git clone https://github.com/TheInvader360/jack-compiler
cd jack-compiler
```

Then:

```bash
go run main.go -path=examples/HelloWorld/Main.jack
```

The translated program is exported to the source directory with the same base filename but a .vm extension (so the given example would generate examples/HelloWorld/Main.vm)

Or:

```bash
go run main.go -path=examples/Sokoban
```

Translates all .jack files in the specified directory and exports them all to the same directory with the same base filename but a .vm extension

## Options

You can use these runtime flags to modify the compiler's behaviour

```bash
-comments
 enable comments in vm output

-debug
 enable debug terminal output

-path=[string]
 jack source path (default "examples/HelloWorld")

-xml
 include intermediate xml files in output
```

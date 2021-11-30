echo "Running tests..."
go test ./... --cover
echo

echo "Compiling examples..."
go run main.go -path=examples/Array -xml
go run main.go -path=examples/Average -xml
go run main.go -path=examples/ComplexArrays -xml
go run main.go -path=examples/ConvertToBin -xml
go run main.go -path=examples/HelloWorld -xml
go run main.go -path=examples/MultiConditionIf -xml
go run main.go -path=examples/Pong -xml
go run main.go -path=examples/Seven -xml
go run main.go -path=examples/SimpleIf -xml
go run main.go -path=examples/SimpleWhile -xml
go run main.go -path=examples/Sokoban -xml
go run main.go -path=examples/Square -xml
go run main.go -path=examples/SquareExpressionless -xml
echo

cd ../nand2tetris/tools/

echo "Comparing t-xml files..."
./TextComparer.sh ../projects/10/materials/ArrayTest/MainT.xml ../../jack-compiler/examples/Array/MainT.xml 
./TextComparer.sh ../projects/10/materials/Square/MainT.xml ../../jack-compiler/examples/Square/MainT.xml 
./TextComparer.sh ../projects/10/materials/Square/SquareT.xml ../../jack-compiler/examples/Square/SquareT.xml 
./TextComparer.sh ../projects/10/materials/Square/SquareGameT.xml ../../jack-compiler/examples/Square/SquareGameT.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/MainT.xml ../../jack-compiler/examples/SquareExpressionless/MainT.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/SquareT.xml ../../jack-compiler/examples/SquareExpressionless/SquareT.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/SquareGameT.xml ../../jack-compiler/examples/SquareExpressionless/SquareGameT.xml 
echo

echo "Comparing xml files..."
./TextComparer.sh ../projects/10/materials/ArrayTest/Main.xml ../../jack-compiler/examples/Array/Main.xml 
./TextComparer.sh ../projects/10/materials/Square/Main.xml ../../jack-compiler/examples/Square/Main.xml 
./TextComparer.sh ../projects/10/materials/Square/Square.xml ../../jack-compiler/examples/Square/Square.xml 
./TextComparer.sh ../projects/10/materials/Square/SquareGame.xml ../../jack-compiler/examples/Square/SquareGame.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/Main.xml ../../jack-compiler/examples/SquareExpressionless/Main.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/Square.xml ../../jack-compiler/examples/SquareExpressionless/Square.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/SquareGame.xml ../../jack-compiler/examples/SquareExpressionless/SquareGame.xml 
echo

echo "Comparing vm files..."
./TextComparer.sh ../../jack-compiler/examples/Average/Main.vm-golden ../../jack-compiler/examples/Average/Main.vm
./TextComparer.sh ../../jack-compiler/examples/ComplexArrays/Main.vm-golden ../../jack-compiler/examples/ComplexArrays/Main.vm
./TextComparer.sh ../../jack-compiler/examples/ConvertToBin/Main.vm-golden ../../jack-compiler/examples/ConvertToBin/Main.vm
./TextComparer.sh ../../jack-compiler/examples/MultiConditionIf/Main.vm-golden ../../jack-compiler/examples/MultiConditionIf/Main.vm
./TextComparer.sh ../../jack-compiler/examples/Pong/Ball.vm-golden ../../jack-compiler/examples/Pong/Ball.vm
./TextComparer.sh ../../jack-compiler/examples/Pong/Bat.vm-golden ../../jack-compiler/examples/Pong/Bat.vm
./TextComparer.sh ../../jack-compiler/examples/Pong/Main.vm-golden ../../jack-compiler/examples/Pong/Main.vm
./TextComparer.sh ../../jack-compiler/examples/Pong/PongGame.vm-golden ../../jack-compiler/examples/Pong/PongGame.vm
./TextComparer.sh ../../jack-compiler/examples/Seven/Main.vm-golden ../../jack-compiler/examples/Seven/Main.vm
./TextComparer.sh ../../jack-compiler/examples/SimpleIf/Main.vm-golden ../../jack-compiler/examples/SimpleIf/Main.vm
./TextComparer.sh ../../jack-compiler/examples/SimpleWhile/Main.vm-golden ../../jack-compiler/examples/SimpleWhile/Main.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Board.vm-golden ../../jack-compiler/examples/Sokoban/Board.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Cell.vm-golden ../../jack-compiler/examples/Sokoban/Cell.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/CellType.vm-golden ../../jack-compiler/examples/Sokoban/CellType.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Controller.vm-golden ../../jack-compiler/examples/Sokoban/Controller.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Direction.vm-golden ../../jack-compiler/examples/Sokoban/Direction.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Level.vm-golden ../../jack-compiler/examples/Sokoban/Level.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/LevelManager.vm-golden ../../jack-compiler/examples/Sokoban/LevelManager.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Main.vm-golden ../../jack-compiler/examples/Sokoban/Main.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Model.vm-golden ../../jack-compiler/examples/Sokoban/Model.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Player.vm-golden ../../jack-compiler/examples/Sokoban/Player.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Sprites.vm-golden ../../jack-compiler/examples/Sokoban/Sprites.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/State.vm-golden ../../jack-compiler/examples/Sokoban/State.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/Utils.vm-golden ../../jack-compiler/examples/Sokoban/Utils.vm
./TextComparer.sh ../../jack-compiler/examples/Sokoban/View.vm-golden ../../jack-compiler/examples/Sokoban/View.vm
./TextComparer.sh ../../jack-compiler/examples/Square/Main.vm-golden ../../jack-compiler/examples/Square/Main.vm
./TextComparer.sh ../../jack-compiler/examples/Square/Square.vm-golden ../../jack-compiler/examples/Square/Square.vm
./TextComparer.sh ../../jack-compiler/examples/Square/SquareGame.vm-golden ../../jack-compiler/examples/Square/SquareGame.vm
echo

echo "Launching emulator..."
./VMEmulator.sh
echo

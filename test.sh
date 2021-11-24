go test ./... --cover

go run main.go -path=examples/Array
go run main.go -path=examples/Average
go run main.go -path=examples/ComplexArrays
go run main.go -path=examples/ConvertToBin
go run main.go -path=examples/HelloWorld
go run main.go -path=examples/Pong
go run main.go -path=examples/Seven
go run main.go -path=examples/SimpleIf
go run main.go -path=examples/SimpleWhile
go run main.go -path=examples/Square
go run main.go -path=examples/SquareExpressionless

cd ../nand2tetris/tools/
./TextComparer.sh ../projects/10/materials/ArrayTest/MainT.xml ../../jack-compiler/examples/Array/MainT.xml 
./TextComparer.sh ../projects/10/materials/Square/MainT.xml ../../jack-compiler/examples/Square/MainT.xml 
./TextComparer.sh ../projects/10/materials/Square/SquareT.xml ../../jack-compiler/examples/Square/SquareT.xml 
./TextComparer.sh ../projects/10/materials/Square/SquareGameT.xml ../../jack-compiler/examples/Square/SquareGameT.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/MainT.xml ../../jack-compiler/examples/SquareExpressionless/MainT.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/SquareT.xml ../../jack-compiler/examples/SquareExpressionless/SquareT.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/SquareGameT.xml ../../jack-compiler/examples/SquareExpressionless/SquareGameT.xml 

./TextComparer.sh ../projects/10/materials/ArrayTest/Main.xml ../../jack-compiler/examples/Array/Main.xml 
./TextComparer.sh ../projects/10/materials/Square/Main.xml ../../jack-compiler/examples/Square/Main.xml 
./TextComparer.sh ../projects/10/materials/Square/Square.xml ../../jack-compiler/examples/Square/Square.xml 
./TextComparer.sh ../projects/10/materials/Square/SquareGame.xml ../../jack-compiler/examples/Square/SquareGame.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/Main.xml ../../jack-compiler/examples/SquareExpressionless/Main.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/Square.xml ../../jack-compiler/examples/SquareExpressionless/Square.xml 
./TextComparer.sh ../projects/10/materials/ExpressionLessSquare/SquareGame.xml ../../jack-compiler/examples/SquareExpressionless/SquareGame.xml 

./VMEmulator.sh


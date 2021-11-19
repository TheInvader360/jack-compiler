go test ./... --cover

go run main.go examples/Array
go run main.go examples/Average
go run main.go examples/ComplexArrays
go run main.go examples/ConvertToBin
go run main.go examples/HelloWorld
go run main.go examples/Pong
go run main.go examples/Seven
go run main.go examples/Square
go run main.go examples/SquareExpressionless

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


go test ./... --cover
go run main.go examples/Array
go run main.go examples/HelloWorld
go run main.go examples/Square
go run main.go examples/SquareExpressionless
cd ../nand2tetris/tools/
./TextComparer.sh ../projects/10/ArrayTest/MainT.xml ../../jack-compiler/examples/Array/MainT.xml 
./TextComparer.sh ../projects/10/Square/MainT.xml ../../jack-compiler/examples/Square/MainT.xml 
./TextComparer.sh ../projects/10/Square/SquareT.xml ../../jack-compiler/examples/Square/SquareT.xml 
./TextComparer.sh ../projects/10/Square/SquareGameT.xml ../../jack-compiler/examples/Square/SquareGameT.xml 
./TextComparer.sh ../projects/10/ExpressionLessSquare/MainT.xml ../../jack-compiler/examples/SquareExpressionless/MainT.xml 
./TextComparer.sh ../projects/10/ExpressionLessSquare/SquareT.xml ../../jack-compiler/examples/SquareExpressionless/SquareT.xml 
./TextComparer.sh ../projects/10/ExpressionLessSquare/SquareGameT.xml ../../jack-compiler/examples/SquareExpressionless/SquareGameT.xml 
./TextComparer.sh ../projects/10/ArrayTest/Main.xml ../../jack-compiler/examples/Array/Main.xml 
./TextComparer.sh ../projects/10/Square/Main.xml ../../jack-compiler/examples/Square/Main.xml 
./TextComparer.sh ../projects/10/Square/Square.xml ../../jack-compiler/examples/Square/Square.xml 
./TextComparer.sh ../projects/10/Square/SquareGame.xml ../../jack-compiler/examples/Square/SquareGame.xml 
./TextComparer.sh ../projects/10/ExpressionLessSquare/Main.xml ../../jack-compiler/examples/SquareExpressionless/Main.xml 
./TextComparer.sh ../projects/10/ExpressionLessSquare/Square.xml ../../jack-compiler/examples/SquareExpressionless/Square.xml 
./TextComparer.sh ../projects/10/ExpressionLessSquare/SquareGame.xml ../../jack-compiler/examples/SquareExpressionless/SquareGame.xml 


# Test file for the quad program 

(https://github.com/01-edu/public/tree/master/subjects/quad)


\
\
Will work if you program is structured as shown (exporting via the piscine package):

![Package structure](https://i.imgur.com/sig1jwD.png)


Place the files in your repo. (same folder that your **main** project file is located)

or copy and paste this in your terminal

	curl -o sudoku_test.go https://raw.githubusercontent.com/nik-don/01-founders/main/sudoku/sudoku_test.go && curl -o test-cases.txt https://raw.githubusercontent.com/nik-don/01-founders/main/sudoku/test-cases.txt




\
\
Usage:



	go test

or


	go test quad_test.go 

----


	go clean -testcache && go test
to reset cache *and* test


\
\
You can edit the test-cases file and add your own.
\
#<Quad>,<x>,<y>#
\
Expected Outcome



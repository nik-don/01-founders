# Test file for the sudoku solver program  [![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/J3J36ZB3M)

(https://github.com/01-edu/public/tree/master/subjects/sudoku)


\
\
Place the files in your repo. (same folder that your **main** project file is located)

## simply copy and paste this in your terminal

	curl -o sudoku_test.go https://raw.githubusercontent.com/nik-don/01-founders/main/sudoku/sudoku_test.go && curl -o test-cases.txt https://raw.githubusercontent.com/nik-don/01-founders/main/sudoku/test-cases.txt
\
\
Make sure you change the program name to main.go or change it inside the test file:

![Screenshot 2021-11-16 222523](https://user-images.githubusercontent.com/93073558/142076129-c7825560-ab72-4b9d-8780-f10f108ca312.jpg)

\
\
Usage:



	go test 

or


	go test sudoku_test.go 

----


	go clean -testcache && go test . 
to reset cache *and* test


![Peek 2021-11-14 09-27](https://user-images.githubusercontent.com/93073558/141675489-c26a910b-f0d9-43e5-971a-6ce39cdc1760.gif)



\
\
You can edit the test-cases file and add your own.
\
#Input#
\
Expected Outcome


## FAQ
**Question:** Why does the test fail when the expected solution looks identical to the program output?

**Answer:** Try running a valid sudoku input with the `cat -e` and make sure there are **NO spaces** after the last digit. 
\
\
example `go run . "1.58.2..." ".9..764.5" "2..4..819" ".19..73.6" "762.83.9." "....61.5." "..76...3." "43..2.5.1" "6..3.89.." | cat -e`




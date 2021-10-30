# Test file for the Ascii-art program 

which consists in receiving a string as an argument 
and outputting the string in a graphic representation using ASCII

(https://github.com/01-edu/public/tree/master/subjects/ascii-art)


\
\
Place the files in your repo. (same folder that your **main** project file is located)

or copy and paste this in your terminal

	curl -o ascii-art_test.go https://raw.githubusercontent.com/nik-don/01-founders/main/ascii-art/ascii-art_test.go && curl -o test-cases.txt https://raw.githubusercontent.com/nik-don/01-founders/main/ascii-art/test-cases.txt


![Peek 2021-10-30 21-51](https://user-images.githubusercontent.com/93073558/139558058-09dae194-9e19-464b-a4f0-adfd50c62985.gif)


\
\
Usage:



	go test . 

or


	go clean -testcache && go test . -v
to reset cache *and* verbose output of what is happening


\
\
You can edit the test-cases file and add your own.
#Argument#
Expected Outcome



\
\
To Do:
- Make the test file run parallel tests


\
\
Inspired by https://github.com/N1X0N3/ascii-art-testfile

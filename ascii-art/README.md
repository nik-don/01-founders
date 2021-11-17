# Test file for the Ascii-art program [![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/J3J36ZB3M)

which consists in receiving a string as an argument 
and outputting the string in a graphic representation using ASCII

(https://github.com/01-edu/public/tree/master/subjects/ascii-art)


\
\
Place the files in your repo. (same folder that your **main** project file is located)

## simply copy and paste this in your terminal

	curl -o ascii-art_test.go https://raw.githubusercontent.com/nik-don/01-founders/main/ascii-art/ascii-art_test.go && curl -o test-cases.txt https://raw.githubusercontent.com/nik-don/01-founders/main/ascii-art/test-cases.txt
\
\
Make sure you change the program name to main.go or change it inside the test file:

![Screenshot 2021-11-16 222523](https://user-images.githubusercontent.com/93073558/142076129-c7825560-ab72-4b9d-8780-f10f108ca312.jpg)

\
\

![Peek 2021-10-30 21-51](https://user-images.githubusercontent.com/93073558/139558058-09dae194-9e19-464b-a4f0-adfd50c62985.gif)


\
\
Usage:



	go test	

use -v for verbose

or


	go test ascii-art_test.go -v

----


	go clean -testcache && go test -v 
to reset cache *and* test


\
\
You can edit the test-cases file and add your own.
\
#Argument#
\
Expected Outcome



----
If you want to:
\
You can comment/uncomment these lines:
\
![Peek 2021-11-03 07-47](https://user-images.githubusercontent.com/93073558/140024727-67521f66-ae98-4ff0-8b7c-5212d0078cfa.gif)
\

\
To test for/do not test for these cases:

![image](https://user-images.githubusercontent.com/93073558/140023337-99ba2081-56bc-492e-bb2b-9ef16071b59c.png)


----




\
\
To Do:
- Make the test file run parallel tests


\
\
Inspired by https://github.com/N1X0N3/ascii-art-testfile


----

## FAQ
**Question:** Is the code same as N1X0N3's test file?

**Answer:** Although it is inspired by it, the code is 99% different and the way the test-cases are stored in the file is different.

##
**Q:** What resources did you look to make this test?

**A:** https://gist.github.com/nik-don/e88dd0e519ff6d693bd8212d148df152
##

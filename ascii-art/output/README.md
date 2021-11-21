[<img height='36' style='border:0px;height:36px;' src='https://cdn.ko-fi.com/cdn/kofi2.png?v=3' border='0' alt='Buy Me a Coffee at ko-fi.com' />](https://ko-fi.com/J3J36ZB3M)

# Test file for the Ascii-art-reverse program 



which consists of saving the graphic representation of text in ASCII characters into a file.

(https://github.com/01-edu/public/tree/master/subjects/ascii-art/output)


\
\
Place the files in your repo. (same folder that your **main** project file is located)

## simply copy and paste this in your terminal

	curl -o ascii-art_output_test.go https://raw.githubusercontent.com/nik-don/01-founders/main/ascii-art/output/ascii_art_output_test.go && curl -o test-cases.txt https://raw.githubusercontent.com/nik-don/01-founders/main/ascii-art/output/test-cases.txt
\
\
Make sure you change the program name to main.go or change it inside the test file:

![Screenshot 2021-11-16 222523](https://user-images.githubusercontent.com/93073558/142076129-c7825560-ab72-4b9d-8780-f10f108ca312.jpg)




\
\
Usage:



	go test	

use -v for verbose

or


	go test ascii-art_output_test.go -v

----


	go clean -testcache && go test -v 
to reset cache *and* test


\
\
You can edit the test-cases file and add your own.
\
#"Argument" --reverse=fileName#
\
Text to be deciphered




\
\
To Do:
- Make the test file run parallel tests


\
\
\
[<img src="https://img.buymeacoffee.com/button-api/?text=Buy me a beer&emoji=🍺&slug=nikdon&button_colour=FFDD00&font_colour=000000&font_family=Cookie&outline_colour=000000&coffee_colour=ffffff">](https://www.buymeacoffee.com/nikdon)

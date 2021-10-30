# Test file for the Ascii-art program 

which consists in receiving a string as an argument 
and outputting the string in a graphic representation using ASCII

(https://github.com/01-edu/public/tree/master/subjects/ascii-art)

\
\
Usage:

Place the files in your repo. (same folder that your main project file is located)

	go test . 

or


	go clean -testcache && go test . -v
to reset cache *and* verbose output of what is happening


\
\
You can edit the test-cases file and add your own.
#Argument#
Expected Outcome

Example:
#Hi#
 _    _   _  $
| |  | | (_) $
| |__| |  _  $
|  __  | | | $
| |  | | | | $
|_|  |_| |_| $
             $
             $


\
\
To Do:
- Make the test file run parallel tests



Inspired by https://github.com/N1X0N3/ascii-art-testfile
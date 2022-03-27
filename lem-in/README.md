              "=.
             "=. \
                \ \
             _,-=\/=._        _.-,_
            /         \      /=-._ "-.
           |=-./~\___/~\    /     `-._\
           |   \o/   \o/   /         /
            \_   `~~~;/    | LEM-IN  |
              `~,._,-'    /          /
                 | |      =-._      /
             _,-=/ \=-._     /|`-._/
           //           \\   )\
          /|             |)_.'/
         //|             |\_."   _.-\
        (|  \           /    _.`=    \
        ||   ":_    _.;"_.-;"   _.-=.:
     _-."/    / `-."\_."        =-_.;\
    `-_./   /             _.-=.    / \\
           |              =-_.;\ ."   \\
           \                   \\/     \\
           /\_                .'\\      \\
          //  `=_         _.-"   \\      \\
         //      `~-.=`"`'       ||      ||
         ||    _.-_/|            ||      |\_.-_
     _.-_/|   /_.-._/            |\_.-_  \_.-._\
    /_.-._/                      \_.-._\
 
(art from https://ascii.co.uk)


Problems:

- Decipher the problem..
- Find the paths an ant can take
- Print out the path each ant takes, keeping account of the contraints.
    



----
## Lem-in Solution

Level of Satisfication for this project : 
★★★☆☆ (needs improving)
I would like to re-visit and try to improve it.

Includes a very basic Visualizer.


# You can try it out:

(ubuntu/linux)  https://nextcloud.nikolo.dev/s/ES9i6Hq7i9NgYAM

to directly download to the terminal, paste this:
```
curl -o lem-in-v0-5-nik https://nextcloud.nikolo.dev/s/ES9i6Hq7i9NgYAM/download/lem-in-v0-5-nik
```

if you want an example to try:
```
curl -o example00.txt https://raw.githubusercontent.com/nik-don/01-founders/main/lem-in/example00.txt
```


# How to use:

Provide execution permission:
```
chmod +x lem-in-v0-5-nik
```

./lem-in-v0-5-nik <FILENAME>:
```
./lem-in-v0-5-nik example00.txt
```

https://user-images.githubusercontent.com/93073558/160241679-9a322ba7-dea5-4aaf-a0aa-527ca86afbda.mp4


# Flags you can use:
|    |                                          | flag syntax |
|----|------------------------------------------|--------------------------|
| -c | show counters                            | -c=true, -c=t or just -c |
| -r | show raw solution without data from file |                          |
| -v | show visualizer                          |                          |
| -h | help                                     |                          |


# Notes:

Current version is a bit buggy for example01.

And the visualizer does not look good if the shortest path is at the top.

Please hit ctrl+c and run again.


# Benchmark

Printing only the solution for [example04](https://github.com/01-edu/public/tree/master/subjects/lem-in/examples#example04)

(Please note there has been no pre-optimization. It is quite slow.)
- No Linked lists, graph representation used...
- Mainly maps and structs

![image](https://user-images.githubusercontent.com/93073558/160244265-e16bf480-988c-4d2f-953a-f3618cddfb60.png)
 
Processor (was asked about this)
- win 10 virtual (amd ryzen - shared cores)
- ubuntu (intel celeron)
  
![image](https://user-images.githubusercontent.com/93073558/160244300-65844e4f-ea57-4865-b1d1-1f65b2641679.png)

X - No of Ants.
Y - Time in seconds

# My Write up

[Read..](https://github.com/nik-don/01-founders/blob/main/lem-in/log.md)

# Audit Solutions

https://user-images.githubusercontent.com/93073558/160259628-6ea3597f-c262-4ac9-86f1-704c685dec5a.mp4


https://user-images.githubusercontent.com/93073558/160259636-727b37be-8e3a-4a67-ba4b-7a977f4010c7.mp4


https://user-images.githubusercontent.com/93073558/160259644-0e4c3fc5-8474-4513-a3ce-4c2371ed6e7a.mp4

https://user-images.githubusercontent.com/93073558/160259651-2bda75d7-5398-4320-bb9c-1f75c45a14cd.mp4

https://user-images.githubusercontent.com/93073558/160259656-998b3d18-1caa-414e-b2f8-eaa2270103e3.mp4

https://user-images.githubusercontent.com/93073558/160259661-389f54f0-c556-4cd9-ae66-7a9547f4fefd.mp4

https://user-images.githubusercontent.com/93073558/160259671-09648f9c-93e7-4df8-b97a-13c691d00485.mp4








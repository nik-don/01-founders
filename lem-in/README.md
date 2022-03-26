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

Includes a very basic 2D Visualizer.


# You can try it out:

(ubuntu/linux)  https://nextcloud.nikolo.dev/s/ES9i6Hq7i9NgYAM

or paste this:
```
curl -o lem-in-v0-5-nik https://nextcloud.nikolo.dev/s/ES9i6Hq7i9NgYAM/download/lem-in-v0-5-nik
```

# How to use:

https://user-images.githubusercontent.com/93073558/160241679-9a322ba7-dea5-4aaf-a0aa-527ca86afbda.mp4

```
./lem-in-v0-5-nik <FILENAME>
```

# Flags you can use:
```
-c
-r
-v
-h
```

# Notes:

Current version is a bit buggy for example01.

And the visualizer does not look good if the shortest path is at the top.

Please hit ctrl+c and run again.


# Benchmark

Printing only the solution.

(Please note there has been no pre-optimization)
- No Linked lists, graph representation
- Mainly maps and structs

![image](https://user-images.githubusercontent.com/93073558/160244265-e16bf480-988c-4d2f-953a-f3618cddfb60.png)

![image](https://user-images.githubusercontent.com/93073558/160244300-65844e4f-ea57-4865-b1d1-1f65b2641679.png)

X - No of Ants.
Y - Time in seconds

# My Write up

[Read..](https://github.com/nik-don/01-founders/blob/main/lem-in/log.md)

# Audit Solutions


https://user-images.githubusercontent.com/93073558/160242525-b3dcd659-ed42-4e89-9eed-fd6821d99c42.mp4

https://user-images.githubusercontent.com/93073558/160242529-a9a94072-c2f7-4299-832e-892872424e8a.mp4


https://user-images.githubusercontent.com/93073558/160242539-0531bf54-ba80-46c3-9f59-7296292dc712.mp4


https://user-images.githubusercontent.com/93073558/160242544-5481364a-ddac-44e5-8d40-0e10b66e8266.mp4




https://user-images.githubusercontent.com/93073558/160242552-057aa4b2-bf0f-455b-91e9-bc1a1f62bcd7.mp4



https://user-images.githubusercontent.com/93073558/160242556-6709092e-a83d-4cb1-a41f-db3f5451564d.mp4




https://user-images.githubusercontent.com/93073558/160242558-d5e9c8b7-2479-437b-93bd-375cf30aed3d.mp4









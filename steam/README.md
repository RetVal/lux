# steam
steam API for Go, http://store.steampowered.com/


using steamworks v1.34
first go in `wrapper/` and `make && make copy` which will build the c wrapper around the actual steamworks library and copy it in the location that the go wrapper can use.
then `go install github.com/luxengine/steam` and you should be good. 

The install process is still very young so I don't know how it can work on windows and linux (linux should be pretty easy to extend to)

make an issue if you actually want to use this it'll go up the priotity list.

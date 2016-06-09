osdetected=$(shell uname -s)

ifeq ($(osdetected), Linux)
	libext=so
	libdir=linux$(shell getconf LONG_BIT)
	flags=-shared -fPIC
endif

ifeq ($(osdetected), Darwin)
	libext=dylib
	libdir=osx32
	flags=-dynamiclib
endif

all:
	g++ $(shell echo $(flags)) -o libsteam_capi.$(shell echo $(libext)) -Lsdk/redistributable_bin/$(shell echo $(libdir)) -lsteam_api -Isdk/public/steam *.cpp

clean:
	rm *.o

copy:
	mkdir ../lib || true
	mkdir ../h || true
	cp libsteam_capi.$(shell echo $(libext)) ../lib/
	cp libsteam_capi.$(shell echo $(libext)) ../
	cp sdk/redistributable_bin/$(shell echo $(libdir))/libsteam_api.$(shell echo $(libext)) ../lib
	cp sdk/redistributable_bin/$(shell echo $(libdir))/libsteam_api.$(shell echo $(libext)) ../
	cp *.h ../h/
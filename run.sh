#!/usr/bin/fish

go build
./rtiow >out.ppm
eog out.ppm

@echo off
REM cd build-stage && docker build -t gocv/gocv-alpine:3.4.2-buildstage .
cd gobuild && docker build -t gocv/gocv-alpine:3.4.2-gobuild .
REM cd ../runtime && docker build -t vin_test/gocv .
cd ..
docker run -it --rm -v /shared:/shared -p 1111:1111 gocv/gocv-alpine:3.4.2-gobuild
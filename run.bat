@echo off
cd build-stage && docker build -t gocv/gocv-alpine:3.4.2-buildstage .
cd ../runtime && docker build -t vin_test/gocv .
cd ..
docker run -it --rm -v /shared:/shared -p 1111:1111 vin_test/gocv
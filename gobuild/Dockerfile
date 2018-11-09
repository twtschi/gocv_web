FROM gocv/gocv-alpine:3.4.2-buildstage as build-stage
FROM denismakogon/ffmpeg-alpine:4.0-golang

# OpenCV 3.4.1 shared objects from build-stage
COPY --from=build-stage /usr/local/lib64 /usr/local/lib64
COPY --from=build-stage /usr/local/lib64/pkgconfig/opencv.pc /usr/local/lib64/pkgconfig/opencv.pc
COPY --from=build-stage /usr/local/include/opencv /usr/local/include/opencv
COPY --from=build-stage /usr/local/include/opencv2 /usr/local/include/opencv2

ENV PKG_CONFIG_PATH /usr/local/lib64/pkgconfig
ENV LD_LIBRARY_PATH /usr/local/lib64
ENV CGO_CPPFLAGS -I/usr/local/include
ENV CGO_CXXFLAGS "--std=c++1z"
ENV CGO_LDFLAGS "-L/usr/local/lib -lopencv_core -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d -lopencv_plot -lopencv_tracking"

ENV BUILD="ca-certificates \
         git \
         build-base \
         musl-dev \
         alpine-sdk \
         make \
         gcc \
         g++ \
         libc-dev \
         linux-headers \
         libjpeg-turbo \
         libpng \
         libwebp \
         libwebp-dev \
         tiff \
         libavc1394 \
         jasper-libs \
         openblas \
         libgphoto2 \
         gstreamer \
         gst-plugins-base"

ENV DEV="clang clang-dev cmake pkgconf \
         openblas-dev gstreamer-dev gst-plugins-base-dev \
         libgphoto2-dev libjpeg-turbo-dev libpng-dev \
         tiff-dev jasper-dev libavc1394-dev"


RUN apk update && \
    apk add --no-cache ${BUILD} ${DEV}

# copy workdir
ENV GOPATH /work_dir
COPY . /work_dir
RUN cd /work_dir
# RUN go get -u -d gocv.io/x/gocv
RUN go build -o /work_dir/bin/main main
WORKDIR /work_dir/bin
# CMD ["./main"]

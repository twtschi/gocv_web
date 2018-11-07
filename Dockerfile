FROM denismakogon/gocv-alpine:3.4.2-buildstage as build-stage

COPY . /work_dir
RUN cd /work_dir

RUN go get -u -d gocv.io/x/gocv
RUN go build -o bin/main main
# RUN cd $GOPATH/src/gocv.io/x/gocv && go build -o $GOPATH/bin/gocv-version ./cmd/version/main.go

# FROM denismakogon/gocv-alpine:3.4.2-runtime

# COPY --from=build-stage /go/bin/gocv-version /gocv-version
# ENTRYPOINT ["/gocv-version"]

# Copy golang server

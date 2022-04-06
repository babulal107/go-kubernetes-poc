ARG APP_DIR=/go/src/go-kubernetes-poc
FROM golang:latest AS build-stage
ARG APP_DIR
RUN mkdir -p ${APP_DIR}
COPY . ${APP_DIR}
WORKDIR ${APP_DIR}
RUN export CGO_ENABLED=0 && export GOOS=linux && export GOMOD111=on &&\
 go mod tidy &&\
 cd cmd/service &&\
 go build -tags netgo -a -v -o go_k8s_app .

#Final Build
FROM alpine:3.12
ARG APP_DIR
WORKDIR  /go/bin
COPY --from=build-stage ${APP_DIR}/cmd/service/go_k8s_app ./
COPY --from=build-stage ${APP_DIR}/configs/* ./configs/
EXPOSE 8080/tcp
ENTRYPOINT ["/go/bin/go_k8s_app"]
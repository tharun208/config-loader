FROM golang:1.16-alpine AS build-env
WORKDIR /home/config-loader
COPY . /home/config-loader/
RUN go build -o config-loader ./app/cmd

FROM golang:1.16-alpine
COPY --from=build-env /home/config-loader/config-loader /home/config-loader-src/config-loader
WORKDIR /home/config-loader-src
CMD ["/home/config-loader-src/config-loader"]





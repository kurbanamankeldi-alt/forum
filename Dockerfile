#  build stage 
FROM golang:1.27-alpine AS build

# go-sqlite3 driver needs a C toolchain.
RUN apk add --no-cache gcc musl-dev

WORKDIR /src

# cache the dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy the rest and build the binary
COPY . .
RUN CGO_ENABLED=1 go build -o /forum .

# runtime stage, just a lighweight Linux img
FROM alpine:latest

# add unpriviledge user w no password, no home dir, security practice
RUN adduser -D -H hiveGuys 
RUN mkdir /data
# grant user ownership of database so it can create and write to db
RUN chown hiveGuys:hiveGuys /data
COPY --from=build /forum /usr/local/bin/forum

USER hiveGuys

# declare that our server need a volume
VOLUME /data
EXPOSE 8080

ENTRYPOINT ["forum"]
CMD [ "--seed" ]
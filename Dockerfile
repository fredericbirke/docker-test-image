FROM golang:1.13.4 AS build-env
WORKDIR /app

# Copy project into app folder
COPY . .

# RUN: Get Packages
RUN go get -v -d
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=auto GOARCH=amd64 go build -o main


# Build runtime image
FROM sickp/alpine-sshd:latest
COPY --from=build-env /app/main .
EXPOSE 22 29170 29171
ENTRYPOINT ["./main"]

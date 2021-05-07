FROM golang:1.16.3-alpine AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 go build -o mailer ./main/app.go

FROM busybox:1.32.1

WORKDIR /app
COPY --from=build /src/mailer .

CMD ["./mailer"]
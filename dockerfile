FROM golang:latest AS build-backend

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux 
RUN go mod tidy
RUN go build -o calendar .

FROM alpine:latest AS production
COPY --from=build-backend /app .
EXPOSE 8080
CMD [".calendar", "serve"]

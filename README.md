# CalendarGO

## About:

App for students and teachers where:

### Students can

- sign up for class or request additional classes
- chat with teachers
- pick date
- cancel the classes
- download ics file
- check tasks and upload solutions

### Teachers can

- manage their students groups
- accept or reject class
- check students answers

## Technologies:

### Backend:

- Go
- Mysql
- Gin

### Frontend:

- React Next.js
- Tailwind

### Continuous integration

- Github actions

### Development Enviremnt

- Docker

  ```yaml
  "image": mcr.microsoft.com/devcontainers/go:1-1.21-bookworm
  ```

### How to Start

```bash
chmod +x setupsdk.sh
./setupsdk.sh
pushd Calendar
go get .
popd
go work init
go work use Calendar
go work use .
```

### Unit Tests

```bash
pushd Calendar && go test -v && popd
# or
pushd Calendar && go test -json && popd
```

### Build

```bash
go run Main.go serve
```

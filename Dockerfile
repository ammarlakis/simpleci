FROM golang:1.16-alpine as BUILD

WORKDIR /src/
COPY . .
RUN go build -o /build/simpleci

FROM alpine
WORKDIR /app
COPY --from=BUILD /build/simpleci .

EXPOSE 8000
CMD [ "/app/simpleci" ]

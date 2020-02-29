FROM golang:1.13.4 as builder

COPY . /src
WORKDIR /src
RUN go build -o bin/scraping_moneyforward *.go

# chromedriverが入ったubuntu
FROM selenium/standalone-chrome:latest

WORKDIR /app
COPY --from=builder /src/bin/scraping_moneyforward /app/
COPY --from=builder /src/.env /app/
ENTRYPOINT ["./scraping_moneyforward"]

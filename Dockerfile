##
# BUILDER
##
FROM golang:1.17.6-alpine as BUILDER

WORKDIR /srv/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
	go build -a -installsuffix \
	cgo -ldflags '-extldflags "-static"' -o op-hookscatcher-bot .

##
# Main image
##
FROM scratch

WORKDIR /srv/app

COPY --from=BUILDER /srv/app/op-hookscatcher-bot .

EXPOSE 80

CMD ["/srv/app/op-hookscatcher-bot"]

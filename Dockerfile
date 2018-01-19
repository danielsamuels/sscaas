FROM golang:latest 

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app

RUN go get github.com/danielsamuels/sscaas/plugins/define
RUN go get github.com/danielsamuels/sscaas/plugins/dellarism
RUN go get github.com/danielsamuels/sscaas/plugins/excuse
RUN go get github.com/danielsamuels/sscaas/plugins/reddit
RUN go get github.com/danielsamuels/sscaas/plugins/sing
RUN go get github.com/danielsamuels/sscaas/plugins/soundcloud
RUN go get github.com/danielsamuels/sscaas/plugins/troutslap
RUN go get github.com/danielsamuels/sscaas/plugins/uptime
RUN go get github.com/danielsamuels/sscaas/plugins/urbandictionary
RUN go get github.com/danielsamuels/sscaas/plugins/nsfw

RUN go build -o main . 

CMD ["/app/main"]

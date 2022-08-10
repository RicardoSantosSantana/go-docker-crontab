FROM golang:1.17.5-alpine3.15 as stage

WORKDIR /usr/app

COPY . .

RUN go mod init meli 
RUN go mod tidy  
RUN go build -o meli

FROM alpine
RUN apk update && apk add bash
RUN apk add --no-cache libc6-compat 

WORKDIR /usr/app 
COPY --from=stage /usr/app/meli .
COPY --from=stage /usr/app/get_date.sh .

RUN chmod +x get_date.sh
RUN chmod +x meli

RUN crontab -l | { cat; echo "* * * * *  bash /usr/app/get_date.sh"; } | crontab -
CMD ["crond","-f"]

FROM golang:latest

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/whatbugsme
WORKDIR /go/src/whatbugsme

RUN touch server.log

RUN go get ./...
RUN go install ./...
RUN go build -o main

CMD if [ ${APP_ENV} = production ]; \
	then \
	main; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080
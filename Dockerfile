FROM golang:onbuild
WORKDIR /Users/iwanbudihalim/Documents/Project/Exam/go3
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

EXPOSE 8080

CMD ["app"]
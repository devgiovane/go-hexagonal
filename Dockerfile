FROM golang:1.20
WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
RUN go mod init hexagonal
RUN go get -u github.com/spf13/cobra@latest && \
    go install github.com/spf13/cobra-cli@latest
RUN apt update && apt install sqlite3 -y
CMD ["tail", "-f", "/dev/null"]
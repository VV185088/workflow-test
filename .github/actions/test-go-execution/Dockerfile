FROM golang:1.14

COPY src/get_sum_of_two_values.go /get_sum_of_two_values.go
# COPY entrypoint.sh /entrypoint.sh

# install dependencies
RUN go get github.com/go-yaml/yaml
# build & run golang script to push data to Service Now
RUN go build -o /get_sum_of_two_values /get_sum_of_two_values.go
# ENTRYPOINT [ "/entrypoint.sh" ]

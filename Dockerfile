FROM alpine:3.2
ADD sort-srv /sort-srv
ENTRYPOINT [ "/sort-srv" ]

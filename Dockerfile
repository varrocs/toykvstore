# Strange. It doesn't work with the latest alpine image
FROM debian:latest
COPY ./toykvstore /
CMD ["/toykvstore"]

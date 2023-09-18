# Docker para desarrollo
FROM golang:1.14.3-buster

WORKDIR /go/src/github.com/nmarsollier/resourcesgo

ENV MONGO_URL mongodb://host.docker.internal:27017

# Puerto de Auth Service y debug
EXPOSE 3000

# Just a terminal, manual mode
# CMD ["bash"]

# To run in debug mode
CMD ["go" , "run" , "/go/src/github.com/nmarsollier/resourcesgo"]

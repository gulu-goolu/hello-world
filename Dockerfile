FROM ubuntu
COPY main /bin/.
EXPOSE 8080
entrypoint /bin/main

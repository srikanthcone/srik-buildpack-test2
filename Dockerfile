FROM scratch
EXPOSE 8080
ENTRYPOINT ["/srik-buildpack-test2"]
COPY ./bin/ /
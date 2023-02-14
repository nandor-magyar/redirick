FROM gcr.io/distroless/static
# maybe
#FROM scratch
LABEL maintainer="nandor-magyar"

LABEL org.dyrectorio.proxy.port=8080
LABEL org.dyrectorio.env.REDIRECT=required,url

COPY --chown=nonroot:nonroot ./redirick /redirick
EXPOSE 8080
ENTRYPOINT ["/redirick"]

# static-healthcheck
[![Go Report Card](https://goreportcard.com/badge/wtfcoderz/static-healthcheck)](http://goreportcard.com/report/wtfcoderz/static-healthcheck)
[![](https://images.microbadger.com/badges/image/wtfcoderz/static-healthcheck.svg)](https://microbadger.com/images/wtfcoderz/static-healthcheck)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/wtfcoderz/static-healthcheck/blob/master/LICENSE.md)
[![Docker Automated buil](https://img.shields.io/badge/docker--hub-automatic--build-blue.svg)](https://hub.docker.com/r/wtfcoderz/static-healthcheck/)

This project was created to add healthcheck in docker images without installing curl or other tools  
This is useful for statically compiled images `FROM scratch` for example  

## Standalone Usage

```
Usage of /healthcheck:
  -http value
        httpcheck: ex: my.domain.com:80
  -tcp value
        tcpcheck: ex: my.domain.com:80
```

These flags accept multiple values :  
`/healthcheck -tcp my.domain.com:80 -http my2.domain.com:80 -http 127.0.0.1:8080`

-http do a GET of the provided URI : return OK if response code < 400  
-tcp verify that the privided host:port respond  

Return 0 if all checks are OK  
Return 1 otherwise  

## Dockerfile Usage

In fact, you only need two lines in your existing Dockerfile to add this  
```
COPY        --from=wtfcoderz/static-healthcheck /healthcheck /
HEALTHCHECK --interval=5s --timeout=2s --start-period=1s --retries=2 CMD ["/healthcheck", "-tcp", "127.0.0.1:80"]
```

Dockerfile example for Caddy server with a Healthcheck  
```
FROM  alpine:latest as build
RUN   apk add --no-cache curl \
      && curl --insecure --silent --show-error --fail --location \
        --header "Accept: application/tar+gzip, application/x-gzip, application/octet-stream" -o - \
        "https://caddyserver.com/download/linux/amd64?plugins=http.git,http.cgi" \
        | tar --no-same-owner -xz caddy \
      && chmod 0755 /caddy \
      && /caddy -version

FROM  scratch
COPY  --from=build /caddy /caddy
COPY        --from=wtfcoderz/static-healthcheck /healthcheck /
HEALTHCHECK --interval=5s --timeout=2s --start-period=1s --retries=2 CMD ["/healthcheck", "-tcp", "127.0.0.1:2015"]
COPY  Caddyfile /etc/Caddyfile
CMD   ["/caddy", "-conf","/etc/Caddyfile","--log","stdout"]
```



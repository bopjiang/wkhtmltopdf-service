FROM golang:1.10.0 AS build
COPY wkhtmltopdf-service.go wkhtmltopdf-service.go
RUN go build -o /bin/wkhtmltopdf-service wkhtmltopdf-service.go

FROM debian:9.3-slim
RUN apt-get update && apt-get install -y \
        wget \
        libfontconfig1 \
        libssl1.0.2 \
        libxext6 \
        libxrender1 \
        xz-utils
RUN wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz -O /tmp/wkhtmltox.tar.xz && \
        tar xvJf /tmp/wkhtmltox.tar.xz -C /tmp && \ 
        mv /tmp/wkhtmltox/bin/wkhtmltopdf /bin/wkhtmltopdf && \
        rm -fr /tmp/wkhtmltox
COPY --from=build /bin/wkhtmltopdf-service /bin/wkhtmltopdf-service 
EXPOSE 80
ENTRYPOINT ["/bin/wkhtmltopdf-service"]

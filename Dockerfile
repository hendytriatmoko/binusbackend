FROM ubuntu:18.04
COPY ./sellpump.crt /etc/ssl/certs/sellpump.crt
RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates
COPY user user
COPY config config
COPY logs logs
CMD ["./user"]
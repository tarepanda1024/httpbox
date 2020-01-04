FROM alpine:3.11.2
LABEL maintainer=liuxiaodong2017@gmail.com
ADD httpbin /httpbin
EXPOSE 8080
CMD ["/httpbin","start"]
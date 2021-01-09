FROM golang:1.15 as builder
COPY . /blog/
WORKDIR /blog
RUN go build -ldflags="-s -w"

FROM gcr.io/distroless/base
COPY assets /blog/assets
COPY pages /blog/pages
COPY posts /blog/posts
COPY templates /blog/templates
COPY ./posts.json /blog/posts.json
COPY --from=builder /blog/aflm.dev /blog
WORKDIR /blog
CMD ["./aflm.dev"]

FROM golang:1.15 as builder
COPY . /blog/
WORKDIR /blog
RUN go build

FROM gcr.io/distroless/base
COPY assets /blog/assets
COPY pages /blog/pages
COPY posts /blog/posts
COPY templates /blog/templates
COPY ./posts.yaml /blog/posts.yaml
COPY --from=builder /blog/aflm.dev /blog
WORKDIR /blog
CMD ["./aflm.dev"]
FROM golang:1.17 AS builder

WORKDIR /root/tektonctl

COPY . .
RUN export GOPROXY=https://goproxy.cn \
    && export GOINSECURE=gitlab.sftcwl.com \
    && export GOPRIVATE="gitlab.sftcwl.com" \
    && go build -v


FROM hub.sftcwl.com/op/git-init:v2.0.0

ADD git-credentials /root/.git-credentials
ADD gitconfig /root/.gitconfig

COPY --from=builder /root/tektonctl/tektonctl /bin/tektonctl


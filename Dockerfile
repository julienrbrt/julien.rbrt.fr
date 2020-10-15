FROM golang

COPY . /julien.rbrt.fr/
WORKDIR /julien.rbrt.fr

ARG PORT=8080

RUN go get

RUN cd frontend && GOOS=js GOARCH=wasm go build -o ../public/main.wasm && cd ..
RUN go build -o server

CMD ./server

EXPOSE ${PORT}

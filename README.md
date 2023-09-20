## A small go server to serve images/gifs with image manipulation and caching

Work in progress

### Install

This project use [bimg](https://github.com/h2non/bimg) to manipulate images so you need to install [libvips](https://libvips.github.io/libvips/install.html)

You also need [swag]() to generate the documentation

```
go install github.com/swaggo/swag/cmd/swag@latest
```

and run

```
swag init -g cmd/main.g
```

### Start

Copy .env.example to .env and populate it with your envs and run the server

```
go run cmd/main.go
```

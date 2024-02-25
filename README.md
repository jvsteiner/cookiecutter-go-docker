# Go microservice template for docker

This template was created with `cookiecutter`.

It is intended as a minimal project starter for a Go http service intended to run on Docker. I've use `alpine:golang` as the base image, the `scratch` image in the second stage, and have included a `Makefile` to make it easy to build and run the service.

## Usage

if you have `cookiecutter` installed, you can use this template with the following command:

```bash
cookiecutter gh:jvsteiner/cookiecutter-go-docker
```

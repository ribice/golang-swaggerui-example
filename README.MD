# golang-swaggerui-example

[![Build Status](https://travis-ci.org/ribice/golang-swagger-ui-example.svg?branch=master)](https://travis-ci.org/ribice/golang-swaggerui-example)
[![Coverage Status](https://coveralls.io/repos/github/ribice/golang-swaggerui-example/badge.svg?branch=master)](https://coveralls.io/github/ribice/golang-swaggerui-example?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ribice/golang-swaggerui-example)](https://goreportcard.com/report/github.com/ribice/golang-swaggerui-example)

Golang-swaggerui-example is an example repository for setting up API documentations using SwaggerUI in your Golang project. Detailed post is available on [my blog](https://ribice.ba/swagger-golang).

## API Interaction

The project is deployed on free app engine tier, and accessible via [LINK](https://ribice.ba/goswagg/). Swagger.json file is served on /swagger route - [LINK](https://ribice.ba/goswagg/v1/swagger).

For now, in order to get SwaggerUI working, head over to [Swagger Petstore](http://petstore.swagger.io/#/), and paste the following URL at the top:

```
https://ribice.ba/goswagg/v1/swagger
```

In the future, I might deploy my own SwaggerUI in order to prevent this step. Also, swagger.json should not be in your git repository, instead it should be produced by a CI tool.

## NOTE

The code provided here doesn't follow any standards. It is intended only to demonstrate go-swagger spec generation in a simple go project.
Do not use this project structure/implementation as a reference for your Golang REST projects.

## License

golang-swaggerui-example is licensed under the MIT license. Check the [LICENSE](LICENSE.md) file for details.

## Author

[Emir Ribic](https://ribice.ba)
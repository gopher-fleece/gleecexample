# 🎯 Example Golang REST API using Gleece

[![Build Gleecexample](https://github.com/gopher-fleece/gleecexample/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/gopher-fleece/gleecexample/actions/workflows/build.yml)

A minimal example demonstrating how to build a REST API service using the Gleece library in Go.

## ✨ Features

- 🛠️ Fully featured REST API [Controller](./controllers/gleecexample.ctrl.go) implementation
- 📚 Auto-generated [OpenAPI v3](./openapi/swagger.json) specification 
- 🔄 Auto-generated [routes](./routes/generated-gleece.go) for [gin](https://github.com/gin-gonic/gin) using [Go-playground](https://github.com/go-playground/validator) validation
- ✅ Custom validation with [ValidateStartsWithLetter](./validators/custom.validators.go)
- 🔐 Authentication implementation in [security](./security/authentication.go) for authorization
- 🔄 [GitHub Actions](./.github/workflows/build.yml) pipeline for building service
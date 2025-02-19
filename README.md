# 🎯 Example Golang REST API using Gleece

[![Build Gleecexample](https://github.com/gopher-fleece/gleecexample/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/gopher-fleece/gleecexample/actions/workflows/build.yml) <a href="https://docs.gleece.dev">
    <img src="https://img.shields.io/badge/docs-gleece.dev-blue" alt="Documentation">
</a>

A minimal example demonstrating how to build a REST API service using the Gleece library in Go.

## ✨ Features

- 🎮 Fully-featured REST API [Controller](./controllers/gleecexample.ctrl.go) implementation
- 📝 Auto-generated [OpenAPI v3](./openapi/swagger.json) specification 
- ⚡ Auto-generated [routes](./routes/generated-gleece.go) for [gin](https://github.com/gin-gonic/gin) with [Go-playground](https://github.com/go-playground/validator) validation
- 🛡 Authentication implementation in [security](./security/authentication.go) for authorization
- ✔️ Custom validation with [ValidateStartsWithLetter](./validators/custom.validators.go)
- 🔄 Middlewares with [Log Middleware](./middlewares/log.middlewares.go)
- 🧩 Route's `AfterOperationRoutesExtension` [template extension](./assets/after.operation.extension.hbs) 
- ⚙️ Definition of [gleece.config.json](./gleece.config.json) configuration 
- 🚀 Application [boot process](./main.go)
- 📦 [GitHub Actions](./.github/workflows/build.yml) pipeline for building service
# go-clang/clang-v14

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-clang/clang-v14)](https://pkg.go.dev/github.com/go-clang/clang-v14)
[![GitHub Workflow](https://img.shields.io/github/workflow/status/go-clang/clang-v14/Test/main?label=test&logo=github&style=flat-square)](https://github.com/go-clang/clang-v14/actions)

Native Go bindings for Clang v14 C API.

Currently, generated from [llvm@14.0.0-rc1](https://github.com/llvm/llvm-project/tree/llvmorg-14.0.0-rc1).

## Install/Update

```bash
CGO_LDFLAGS="-L`llvm-config --libdir`" \
  go install github.com/go-clang/clang-v14/...
```

## Usage

An example on how to use the AST visitor of the Clang API can be found in [cmd/go-clang-dump/main.go](cmd/go-clang-dump/main.go)

## I need bindings for a different Clang version

The Go bindings are placed in their own repositories to provide the correct bindings for the corresponding Clang version. A list of supported versions can be found in [go-clang/gen's README](https://github.com/go-clang/gen#where-are-the-bindings).

## I found a bug/missing a feature in go-clang

We are using the issue tracker of the `go-clang/gen` repository. Please go through the [open issues](https://github.com/go-clang/gen/issues) in the tracker first. If you cannot find your request just open up a [new issue](https://github.com/go-clang/gen/issues/new).

## How is this binding generated?

The [go-clang/gen](https://github.com/go-clang/gen) repository is used to automatically generate this binding.

# License

This project, like all go-clang projects, is licensed under a BSD-3 license which can be found in the [LICENSE file](https://github.com/go-clang/license/blob/master/LICENSE) in [go-clang's license repository](https://github.com/go-clang/license)
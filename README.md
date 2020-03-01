[![Go Report Card](https://goreportcard.com/badge/gomodules.xyz/jsonpatch-gen)](https://goreportcard.com/report/gomodules.xyz/jsonpatch-gen)
[![Build Status](https://github.com/gomodules/jsonpatch-gen/workflows/CI/badge.svg)](https://github.com/gomodules/jsonpatch-gen/actions?workflow=CI)
[![Github All Releases](https://img.shields.io/github/downloads/gomodules/jsonpatch-gen/total.svg)](https://github.com/gomodules/jsonpatch-gen/releases)

# jsonpatch-gen
`jsonpatch-gen` generates RFC 6902 json patch from  a pair of json files. Generated patches can be used with `kubectl` or `kustomize`.

### Usage

```console
jsonpatch-gen src.yaml dst.yaml [-type json|strategic] [-o yaml|json]
```

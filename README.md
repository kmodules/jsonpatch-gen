[![Go Report Card](https://goreportcard.com/badge/kmodules.xyz/jsonpatch-gen)](https://goreportcard.com/report/kmodules.xyz/jsonpatch-gen)
[![Build Status](https://github.com/kmodules/jsonpatch-gen/workflows/CI/badge.svg)](https://github.com/kmodules/jsonpatch-gen/actions?workflow=CI)
[![Github All Releases](https://img.shields.io/github/downloads/kmodules/jsonpatch-gen/total.svg)](https://github.com/kmodules/jsonpatch-gen/releases)

# jsonpatch-gen
`jsonpatch-gen` generates RFC 6902 json patch or strategic merge patch from a pair of json/yaml files. Generated patches can be used with `kubectl` or `kustomize`.

### Usage

```console
jsonpatch-gen src.yaml dst.yaml [-t json|strategic] [-o yaml|json]
```

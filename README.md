# Osuny CLI

Application en ligne de commande aidant à coder avec Osuny.

## Installation

```bash
go install github.com/osunyorg/osuny@latest
```

## Usage

```bash
osuny
```

## Contribution

### Local install

```bash 
go install
```

### Versioning

Set version in `cmd/root.go`, like that 

```go
var version = "v0.0.7"
```

Then do

```bash
go mod tidy
git add .
git commit -am "v0.0.7"
git tag v0.0.7
git push origin v0.0.7
git push origin main
```
(not sure about the last 2 commands)


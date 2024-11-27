_default:
  @just --list

_check-build-dir:
  @if [ ! -d "./build/bin" ]; then \
    mkdir -p ./build/bin; \
  fi

compile: _check-build-dir listener publisher mitm

listener: _check-build-dir
  @go build -o ./build/bin/listener ./cmd/listener/main.go

publisher: _check-build-dir
  @go build -o ./build/bin/publisher ./cmd/publisher/main.go

mitm: _check-build-dir
  @go build -o ./build/bin/mitm ./cmd/mitm/main.go

run target *args:
  @if [ -f "./build/bin/{{target}}" ]; then \
    ./build/bin/{{target}} {{args}}; \
  else \
    echo "Erro: File {{target}} not found, please compile it before trying to run it."; \
    exit 1; \
  fi

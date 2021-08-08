# notes

## protocol buffers

Update internal/grpc/player.proto

run protoc like:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/player.proto
```

## testing

- Using ginkgo

- Put everything in it's own pacakge.
  So like config/. In package run `ginkgo bootstrap` then `gingko generate`.
  
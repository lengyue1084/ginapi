.PHONY: wire
# generate wire
wire:
	wire ./cmd/wire.go

.PHONY:run
run:
	go run ./cmd/main.go ./cmd/wire_gen.go


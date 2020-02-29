run: build start

build:
	@echo ">> Building GRPC..."
	@go build -o stockbit-grpc ./cmd/grpc
	@echo ">> Finished"

start-grpc: server-grpc.PID

server-grpc.PID:
	@./stockbit-grpc & echo $$! > $@;

stop-grpc: server-grpc.PID
	kill `cat $<` && rm $<

start: start-grpc
stop: stop-grpc

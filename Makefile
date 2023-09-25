build-server:
	go build -o servidor/servidor servidor/servidor.go

build-client:
	go build -o cliente/cliente cliente/cliente.go

build:
	make build-server
	make build-client

run-server:
	make build-server
	./servidor/servidor

run-client:
	make build-server
	./cliente/cliente
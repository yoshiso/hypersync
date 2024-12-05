build:
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o bin/hypersync main.go

generate:
	go generate ./ent
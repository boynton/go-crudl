API=crudl.sadl
#API=crudl.smithy

TARGET=generated
CMD=crudld

all: generate

test: generate $(TARGET)/bin/test
	$(TARGET)/bin/test

generate: $(TARGET)/test/test.go $(TARGET)/$(CMD)/main.go $(TARGET)/$(CMD)/controller.go $(TARGET)/crudl_server.go $(TARGET)/crudl_client.go $(TARGET)/go.mod

build: generate $(TARGET)/bin/$(CMD)

run:: build
	$(TARGET)/bin/$(CMD)

$(TARGET)/bin/$(CMD) $(TARGET)/bin/test:
	(cd $(TARGET) && mkdir -p bin && go build -o bin/ example/...)

$(TARGET)/go.mod:
	(cd $(TARGET) && go mod init example && go mod tidy)

$(TARGET)/crudl_server.go: $(API)
	mkdir -p $(TARGET)
	sadl -g go-server -o $(TARGET) $(API)

$(TARGET)/crudl_client.go: $(API)
	mkdir -p $(TARGET)
	../sadl/bin/sadl -g go-client -o $(TARGET) $(API)

$(TARGET)/$(CMD)/main.go: main.go
	mkdir -p $(TARGET)/$(CMD)
	cp -p main.go $(TARGET)/$(CMD)/main.go

$(TARGET)/$(CMD)/controller.go: controller.go
	mkdir -p $(TARGET)/$(CMD)
	cp -p controller.go $(TARGET)/$(CMD)/controller.go

$(TARGET)/test/test.go: test.go
	mkdir -p $(TARGET)/test
	cp -p test.go $(TARGET)/test/test.go


clean::
	rm -rf $(TARGET)


swagger: $(TARGET)/crudl.json
	@echo browse to http://localhost:8080/index.html
	local-swagger $(TARGET)/crudl.json

$(TARGET)/crudl.json: crudl.sadl
	sadl -g openapi crudl.sadl > $(TARGET)/crudl.json

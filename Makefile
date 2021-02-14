API=crudl.sadl
#API=crudl.smithy

TARGET=generated
CMD=crudld

all: generate

generate: $(TARGET)/$(CMD)/main.go $(TARGET)/$(CMD)/controller.go $(TARGET)/crudl_server.go $(TARGET)/go.mod

build: generate $(TARGET)/bin/$(CMD)

run:: build
	$(TARGET)/bin/$(CMD)

$(TARGET)/bin/$(CMD):
	(cd $(TARGET) && go build -o bin/ example/...)

$(TARGET)/go.mod:
	(cd $(TARGET) && go mod init example && go mod tidy)

$(TARGET)/crudl_server.go: $(API)
	mkdir -p $(TARGET)
	sadl -g go-server -o $(TARGET) $(API)

$(TARGET)/$(CMD)/main.go: main.go
	mkdir -p $(TARGET)/$(CMD)
	cp -p main.go $(TARGET)/$(CMD)/main.go

$(TARGET)/$(CMD)/controller.go: controller.go
	mkdir -p $(TARGET)/$(CMD)
	cp -p controller.go $(TARGET)/$(CMD)/controller.go

clean::
	rm -rf $(TARGET)



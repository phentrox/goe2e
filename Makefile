install:
	go build -o goseq
	sudo mv goseq /bin

build:
	go build -o goseq
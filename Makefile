install:
	go build -o goSeqE2E ./cmd
	sudo mv goSeqE2E /bin

build:
	go build -o goSeqE2E ./cmd
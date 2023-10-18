.PHONY: tests
tests:
	protoc --protoc_path=./testpb \
		--go_out=paths=source_relative:./testpb \
		--go-xorm_out=paths=source_relative:./testpb \
		./testpb/test.proto
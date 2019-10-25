protoc --proto_path=. --twirp_out=. --go_out=. rpc/checkoperator/checkoperator.proto

retool do build/bin/protoc -Ibuild/protoc/include -I. rpc/platform/platform.proto --go_out=paths=source_relative:. --twirp_out=.

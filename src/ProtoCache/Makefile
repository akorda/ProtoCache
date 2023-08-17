build:
	dotnet build /property:GenerateFullPaths=true /consoleloggerparameters:NoSummary

proto:
	protoc \
	-I=../ \
	--csharp_out=ProtoCache \
	--csharp_opt=base_namespace=ProtoCache,file_extension=.g.cs,internal_access \
	ProtoCacheServer/proto/protoCache.proto

.PHONY: proto
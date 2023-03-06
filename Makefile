# PROTOC变量
PROTOC_VAR = -I. -I../rockim-apis  -I${GOPATH}/src -I../rockim-apis/third_party

protoc-shared:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis ../rockim-apis/rockim/shared/enums/**.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis ../rockim-apis/rockim/shared/reasons/**.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis ../rockim-apis/rockim/shared/**.proto

protoc-admin:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-http_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/admin/manager/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/admin/manager/v1/types/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-http_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/admin/tenant/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/admin/tenant/v1/types/*.proto

protoc-client:
	protoc ${PROTOC_VAR} --go-extend_out=plugins=setters:./apis  --go-extend_opt=paths=source_relative --go-http_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis"  ../rockim-apis/rockim/api/client/http/v1/**/*.proto
	protoc ${PROTOC_VAR} --go-extend_out=plugins=setters:./apis  --go-extend_opt=paths=source_relative --go-http_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis"  ../rockim-apis/rockim/api/client/http/v1/**/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/client/http/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/client/http/v1/**/types/*.proto
	protoc ${PROTOC_VAR} --go-extend_out=plugins=setters:./apis  --go-extend_opt=paths=source_relative --go-grpc_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/client/socket/v1/*.proto


protoc-openapi:
	protoc ${PROTOC_VAR} --go-extend_out=plugins=setters:./apis  --go-extend_opt=paths=source_relative --go-http_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/openapi/v1/*.proto
	#protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/openapi/v1/types/*.proto

protoc-service:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-grpc_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/service/*.proto

protoc-comet:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-grpc_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/service/comet/v1/*.proto


protoc-platform:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-grpc_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/service/platform/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis  ../rockim-apis/rockim/service/platform/v1/types/*.proto

protoc-user:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-grpc_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/service/user/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis  ../rockim-apis/rockim/service/user/v1/types/*.proto

protoc-group:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-grpc_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/service/group/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis  ../rockim-apis/rockim/service/group/v1/types/*.proto


wire:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && wire'
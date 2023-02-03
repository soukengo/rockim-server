# PROTOC变量
PROTOC_VAR = -I. -I../rockim-apis  -I${GOPATH}/src -Ithird_party

protoc-shared:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis ../rockim-apis/rockim/shared/enums/**.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis ../rockim-apis/rockim/shared/reasons/**.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis ../rockim-apis/rockim/shared/**.proto

protoc-java:
	protoc ${PROTOC_VAR} --java_out=./dist/java ../rockim-apis/rockim/shared/enums/**.proto
	protoc ${PROTOC_VAR} --java_out=./dist/java ../rockim-apis/rockim/shared/reasons/**.proto
	protoc ${PROTOC_VAR} --java_out=./dist/java ../rockim-apis/rockim/shared/**.proto
protoc-cpp:
	protoc ${PROTOC_VAR} --cpp_out=./dist/cpp ../rockim-apis/rockim/shared/enums/**.proto
	protoc ${PROTOC_VAR} --cpp_out=./dist/cpp ../rockim-apis/rockim/shared/reasons/**.proto
	protoc ${PROTOC_VAR} --cpp_out=./dist/cpp ../rockim-apis/rockim/shared/**.proto

protoc-admin:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-http_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/admin/manager/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/admin/manager/v1/types/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-http_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/admin/tenant/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/admin/tenant/v1/types/*.proto

protoc-platform:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-grpc_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/service/platform/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis  ../rockim-apis/rockim/service/platform/v1/types/*.proto

protoc-user:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-grpc_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/service/user/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis  ../rockim-apis/rockim/service/user/v1/types/*.proto

protoc-client:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --go-http_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/client/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:./apis --validate_out="paths=source_relative,lang=go:./apis" ../rockim-apis/rockim/api/client/v1/types/*.proto


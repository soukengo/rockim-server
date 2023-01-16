# PROTOC变量
PROTOC_VAR = -I. -I../rockim-apis -I../rockim-apis/api  -I${GOPATH}/src -Ithird_party

protoc-shard:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:. ../rockim-apis/api/rockim/shared/enums/**.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:. ../rockim-apis/api/rockim/shared/reasons/**.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:. ../rockim-apis/api/rockim/shared/**.proto

protoc-admin:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. --validate_out="paths=source_relative,lang=go:." ../rockim-apis/api/rockim/admin/manager/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. --validate_out="paths=source_relative,lang=go:." ../rockim-apis/api/rockim/admin/manager/v1/types/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. --validate_out="paths=source_relative,lang=go:." ../rockim-apis/api/rockim/admin/tenant/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. --validate_out="paths=source_relative,lang=go:." ../rockim-apis/api/rockim/admin/tenant/v1/types/*.proto

protoc-platform:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --validate_out="paths=source_relative,lang=go:." ../rockim-apis/api/rockim/service/platform/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:.  ../rockim-apis/api/rockim/service/platform/v1/types/*.proto

protoc-user:
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --validate_out="paths=source_relative,lang=go:." ../rockim-apis/api/rockim/service/user/v1/*.proto
	protoc ${PROTOC_VAR} --go_out=paths=source_relative:.  ../rockim-apis/api/rockim/service/user/v1/types/*.proto


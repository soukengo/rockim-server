package service

const (
	ProjectName     = "rockim"
	AppGateway      = "rockim.gateway"
	AppAdmin        = "rockim.admin"
	AppPlatform     = "rockim.platform"
	AppUser         = "rockim.user"
	AppGroup        = "rockim.group"
	AppRelationship = "rockim.relationship"
	AppMessage      = "rockim.message"
)

func GenServiceRequest(productId string) *ServiceRequest {
	return &ServiceRequest{ProductId: productId}
}

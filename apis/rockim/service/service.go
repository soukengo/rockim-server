package service

const (
	ProjectName     = "rockim"
	AppAdmin        = "rockim.admin"
	AppGateway      = "rockim.gateway"
	AppComet        = "rockim.comet"
	AppPlatform     = "rockim.platform"
	AppUser         = "rockim.user"
	AppGroup        = "rockim.group"
	AppRelationship = "rockim.relationship"
	AppMessage      = "rockim.message"
)

func GenServiceRequest(productId string) *ServiceRequest {
	return &ServiceRequest{ProductId: productId}
}

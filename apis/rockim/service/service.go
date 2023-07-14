package service

const (
	ProjectName     = "rockim"
	AppAdmin        = "rockim.admin"
	AppGateway      = "rockim.gateway"
	AppComet        = "rockim.comet"
	AppPlatform     = "rockim.platform"
	AppUser         = "rockim.user"
	AppSession      = "rockim.session"
	AppGroup        = "rockim.group"
	AppRelationship = "rockim.relationship"
	AppMessage      = "rockim.message"
	AppJob          = "rockim.job"
)

func GenRequest(productId string) *ServiceRequest {
	return &ServiceRequest{ProductId: productId}
}

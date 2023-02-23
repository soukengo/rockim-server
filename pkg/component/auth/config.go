package auth

import "rockimserver/pkg/component/auth/jwt"

type Config struct {
	Jwt *jwt.Config
}

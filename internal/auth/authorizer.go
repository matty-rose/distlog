package auth

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Authorizer struct {
	enforcer *casbin.Enforcer
}

func New(model, policy string) *Authorizer {
	enforcer, _ := casbin.NewEnforcer(model, policy)

	return &Authorizer{
		enforcer: enforcer,
	}
}

func (a *Authorizer) Authorize(subject, object, action string) error {
	allowed, err := a.enforcer.Enforce(subject, object, action)
	if err != nil {
		return err
	}

	if !allowed {
		msg := fmt.Sprintf("%s not permitted to %s to %s", subject, action, object)
		st := status.New(codes.PermissionDenied, msg)

		return st.Err()
	}

	return nil
}

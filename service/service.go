package service

import (
	"context"
	"github.com/robojones/cloud-identity/db"
	"github.com/robojones/cloud-identity/env"
	"github.com/robojones/cloud-lib-go/identity"
)

func NewService(db db.CockroachDB, env *env.Env) *Service {
	return &Service{
		db:  db,
		env: env,
	}
}

type Service struct {
	db  db.CockroachDB
	env *env.Env
}

func (*Service) DeleteAccount(context.Context, *identity.DeleteAccountRequest) (*identity.DeleteAccountResponse, error) {
	panic("implement me")
}

func (*Service) Login(context.Context, *identity.LoginRequest) (*identity.LoginResponse, error) {
	panic("implement me")
}

func (*Service) ChangePassword(context.Context, *identity.ChangePasswordRequest) (*identity.ChangePasswordResponse, error) {
	panic("implement me")
}

func (*Service) GetSessions(context.Context, *identity.GetSessionsRequest) (*identity.GetSessionsResponse, error) {
	panic("implement me")
}

func (*Service) TerminateSession(context.Context, *identity.TerminateSessionRequest) (*identity.TerminateSessionResponse, error) {
	panic("implement me")
}

func (*Service) CreateTwoFactorKey(context.Context, *identity.CreateTwoFactorKeyRequest) (*identity.CreateTwoFactorKeyResponse, error) {
	panic("implement me")
}

func (*Service) ActivateTwoFactorAuthentication(context.Context, *identity.ActivateTwoFactorAuthenticationRequest) (*identity.ActivateTwoFactorAuthenticationResponse, error) {
	panic("implement me")
}

func (*Service) DeactivateTwoFactorAuthentication(context.Context, *identity.DeactivateTwoFactorAuthenticationRequest) (*identity.DeactivateTwoFactorAuthenticationResponse, error) {
	panic("implement me")
}

func (*Service) Recover(context.Context, *identity.RecoverRequest) (*identity.RecoverResponse, error) {
	panic("implement me")
}

func (*Service) AddEmail(context.Context, *identity.AddEmailRequest) (*identity.AddEmailResponse, error) {
	panic("implement me")
}

func (*Service) RemoveEmail(context.Context, *identity.RemoveEmailRequest) (*identity.RemoveEmailResponse, error) {
	panic("implement me")
}

func (*Service) GetEmails(context.Context, *identity.GetEmailsRequest) (*identity.GetEmailsResponse, error) {
	panic("implement me")
}

func (*Service) ValidateEmail(context.Context, *identity.ValidateEmailRequest) (*identity.ValidateEmailResponse, error) {
	panic("implement me")
}

func (*Service) SetPrimaryEmail(context.Context, *identity.SetPrimaryEmailRequest) (*identity.SetPrimaryEmailResponse, error) {
	panic("implement me")
}

func (*Service) FindUser(context.Context, *identity.FindUserRequest) (*identity.FindUserResponse, error) {
	panic("implement me")
}

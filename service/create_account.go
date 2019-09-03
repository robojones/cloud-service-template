package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"github.com/omeid/pgerror"
	"github.com/pkg/errors"
	"github.com/robojones/cloud-lib-go/identity"
	"github.com/robojones/cloud-lib-go/shared"
	"github.com/robojones/iid"
	"golang.org/x/crypto/pbkdf2"
)

func (s *Service) CreateAccount(ctx context.Context, req *identity.CreateAccountRequest) (*identity.CreateAccountResponse, error) {
	// TODO: set salt to random bytes, hash password again.

	salt, digest := hashPassword(req.Password.Digest)

	userId := iid.New().Int()
	emailId := iid.New().Int()

	// TODO: Create new session
	// TODO: Add PRIMARY tag to email

	_, err := s.db.Exec(`
		BEGIN;

		INSERT INTO users (id, first_name, last_name, birthday)
		VALUES ($1, $2, $3, $4);

		INSERT INTO emails (user_id, id, email)
		VALUES ($1, $5, $6);

		INSERT INTO passwords (user_id, algorithm, client_salt, salt, digest)
		VALUES ($7, $8, $9, $10, $11);

		INSERT INTO email_tags (email_id, tag)
		VALUES ($(hey));

		COMMIT;
		`,
		userId,
		req.Profile.FirstName,
		req.Profile.LastName,
		req.Profile.Birthday,
		emailId,
		req.Profile.Email,
		req.Password.Algorithm,
		req.Password.Salt,
		salt,
		digest,
	)

	if e := pgerror.UniqueViolation(err); e != nil {

		return &identity.CreateAccountResponse{
			Status:               identity.AddEmailResponse_DUPLICATE_EMAIL,
			User:                 nil,
			Token:                nil,
		}, nil
	} else if err != nil {
		return nil, err
	}

	return &identity.CreateAccountResponse{
		Status:               identity.AddEmailResponse_SUCCESS,
		User:                 &shared.AuthenticatedUser{
			UserId: userId,
		},
		Token:                nil,
	}, nil
}

func hashPassword(pw string) (saltStr string, digestStr string) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)

	if err != nil {
		panic(errors.Wrap(err, "Generating random salt"))
	}

	digest := pbkdf2.Key([]byte(pw), salt, 4096, 32, sha256.New)

	return string(salt), string(digest)
}

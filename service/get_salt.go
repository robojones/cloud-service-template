package service

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"github.com/robojones/cloud-lib-go/identity"
)

func (s *Service) GetSalt(ctx context.Context, req *identity.GetSaltRequest) (*identity.GetSaltResponse, error) {
	rows, err := s.db.Query(`
		SELECT u.salt
		FROM users AS u,
		     emails AS m
		WHERE m.email = $1
		AND m.user_id = u.id`, req.Email)

	if err == sql.ErrNoRows {
		salt := generateSalt(req.Email, s.env.SaltSeed)
		return &identity.GetSaltResponse{
			Salt: salt,
		}, nil
	}

	if err != nil {
		return nil, err
	}

	var salt string
	rows.Next()
	err = rows.Scan(&salt)

	if err != nil {
		return nil, err
	}

	return &identity.GetSaltResponse{
		Salt: salt,
	}, nil
}

func generateSalt(email string, salt string) string {
	msg := append([]byte(salt), []byte(email)...)

	digest := sha256.Sum256(msg)
	return base64.StdEncoding.EncodeToString(digest[:])
}

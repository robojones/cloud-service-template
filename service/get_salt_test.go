package service

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/robojones/cloud-identity/env"
	"github.com/robojones/cloud-lib-go/identity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetSalt_FoundInDB(t *testing.T) {
	const (
		email        = "test@test.com"
		salt         = "test salt"
		testSaltSeed = "test salt seed"
	)

	mockDB, mock, _ := sqlmock.New()
	s := NewService(mockDB, &env.Env{
		SaltSeed: testSaltSeed,
	})

	mock.ExpectQuery(".*SELECT.*salt.*FROM.*").
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"salt"}).AddRow(salt))

	r, err := s.GetSalt(context.Background(), &identity.GetSaltRequest{
		Email: email,
	})

	assert.NoError(t, err)
	assert.Equal(t, salt, r.Salt)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_GetSalt_GenerateNew(t *testing.T) {
	const (
		email        = "test@test.com"
		testSaltSeed = "test salt seed"
		// salt is the SHA-256 digest derived from the email and the testSaltSeed.
		salt = "YF3G6TcN2sfgB8/wEW/sNE5sb705vFWsM04ZLZvThGo="
	)

	mockDB, mock, _ := sqlmock.New()
	s := NewService(mockDB, &env.Env{
		SaltSeed: testSaltSeed,
	})

	mock.ExpectQuery(".*SELECT.*salt.*FROM.*").
		WithArgs(email).
		WillReturnError(sql.ErrNoRows)

	r, err := s.GetSalt(context.Background(), &identity.GetSaltRequest{
		Email: email,
	})

	assert.NoError(t, err)
	assert.Equal(t, salt, r.Salt)
	assert.NoError(t, mock.ExpectationsWereMet())
}

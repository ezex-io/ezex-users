package integration

import (
	"testing"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/adapter/database"
	"github.com/ezex-io/ezex-users/internal/adapter/database/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/interactor"
	"github.com/ezex-io/ezex-users/internal/test/testdb"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecurityImageIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Setup is automatically called by NewPostgresDB
	dbConn := testdb.NewPostgresDB(t)
	defer dbConn.Cleanup()

	// create a SecurityImage interactor
	securityImageInteractor := interactor.NewSecurityImage(database.NewSecurityImage(dbConn.Queries))

	// create a test user first
	ctx := t.Context()
	userID := uuid.New()
	email := "foo@bar.com"
	firebaseUID := "firebase-123"

	err := dbConn.Queries.CreateUser(ctx, gen.CreateUserParams{
		ID:           userID,
		Email:        email,
		FirebaseUuid: firebaseUID,
	})
	require.NoError(t, err)

	t.Run("SaveAndRetrieveSecurityImage", func(t *testing.T) {
		securityImage := "image.jpg"
		securityPhrase := "My secret phrase"

		saveReq := &users.SaveSecurityImageRequest{
			Email:          email,
			SecurityImage:  securityImage,
			SecurityPhrase: securityPhrase,
		}

		saveResp, err := securityImageInteractor.SaveSecurityImage(ctx, saveReq)
		require.NoError(t, err)
		assert.Equal(t, email, saveResp.Email)

		getReq := &users.GetSecurityImageRequest{
			Email: email,
		}

		getResp, err := securityImageInteractor.GetSecurityImage(ctx, getReq)
		require.NoError(t, err)
		assert.Equal(t, securityImage, getResp.SecurityImage)
		assert.Equal(t, securityPhrase, getResp.SecurityPhrase)
	})

	t.Run("UpdateExistingSecurityImage", func(t *testing.T) {
		// first save
		saveReq1 := &users.SaveSecurityImageRequest{
			Email:          email,
			SecurityImage:  "original-image-data",
			SecurityPhrase: "original phrase",
		}
		_, err := securityImageInteractor.SaveSecurityImage(ctx, saveReq1)
		require.NoError(t, err)

		// then update with new values
		saveReq2 := &users.SaveSecurityImageRequest{
			Email:          email,
			SecurityImage:  "updated-image-data",
			SecurityPhrase: "updated phrase",
		}
		_, err = securityImageInteractor.SaveSecurityImage(ctx, saveReq2)
		require.NoError(t, err)

		// verify it was updated
		getReq := &users.GetSecurityImageRequest{
			Email: email,
		}
		getResp, err := securityImageInteractor.GetSecurityImage(ctx, getReq)
		require.NoError(t, err)
		assert.Equal(t, "updated-image-data", getResp.SecurityImage)
		assert.Equal(t, "updated phrase", getResp.SecurityPhrase)
	})

	t.Run("NonExistentUser", func(t *testing.T) {
		// try to get security image for non-existent user
		getReq := &users.GetSecurityImageRequest{
			Email: "nonexistent@example.com",
		}
		_, err := securityImageInteractor.GetSecurityImage(ctx, getReq)
		assert.Error(t, err)
	})

	t.Run("DirectDatabaseAccess", func(t *testing.T) {
		saveReq := &users.SaveSecurityImageRequest{
			Email:          email,
			SecurityImage:  "test-direct-image",
			SecurityPhrase: "test-direct-phrase",
		}
		_, err := securityImageInteractor.SaveSecurityImage(ctx, saveReq)
		require.NoError(t, err)

		newImage := "sql-updated-image"
		newPhrase := "sql-updated-phrase"

		err = dbConn.Queries.SaveSecurityImage(ctx, gen.SaveSecurityImageParams{
			Email: email,
			SecurityImage: pgtype.Text{
				String: newImage,
				Valid:  true,
			},
			SecurityPhrase: pgtype.Text{
				String: newPhrase,
				Valid:  true,
			},
		})
		require.NoError(t, err)

		getReq := &users.GetSecurityImageRequest{
			Email: email,
		}
		getResp, err := securityImageInteractor.GetSecurityImage(ctx, getReq)
		require.NoError(t, err)
		assert.Equal(t, newImage, getResp.SecurityImage)
		assert.Equal(t, newPhrase, getResp.SecurityPhrase)
	})
}

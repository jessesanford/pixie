package testingutils

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	pb "pixielabs.ai/pixielabs/src/shared/services/proto"
	"pixielabs.ai/pixielabs/src/shared/services/utils"
)

// TestOrgID is a test org valid UUID
const TestOrgID string = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"

// TestUserID is a test user valid UUID
const TestUserID string = "7ba7b810-9dad-11d1-80b4-00c04fd430c8"

// GenerateTestClaimsWithDuration generates valid test user claims for a specified duration.
func GenerateTestClaimsWithDuration(t *testing.T, duration time.Duration) *pb.JWTClaims {
	claims := utils.GenerateJWTForUser(TestUserID, TestOrgID, "test@test.com", time.Now().Add(duration))
	return claims
}

// GenerateTestClaims generates valid test user claims valid for 60 minutes
func GenerateTestClaims(t *testing.T) *pb.JWTClaims {
	return GenerateTestClaimsWithDuration(t, time.Minute*60)
}

// GenerateTestJWTToken generates valid tokens for testing.
func GenerateTestJWTToken(t *testing.T, signingKey string) string {
	return GenerateTestJWTTokenWithDuration(t, signingKey, time.Minute*60)
}

// GenerateTestJWTTokenWithDuration generates valid tokens for testing with the specified duration.
func GenerateTestJWTTokenWithDuration(t *testing.T, signingKey string, timeout time.Duration) string {
	claims := GenerateTestClaimsWithDuration(t, timeout)

	return SignPBClaims(t, claims, signingKey)
}

// SignPBClaims signs our protobuf claims after converting to json.
func SignPBClaims(t *testing.T, claims *pb.JWTClaims, signingKey string) string {
	mc := utils.PBToMapClaims(claims)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, mc).SignedString([]byte(signingKey))
	if err != nil {
		t.Fatal("failed to generate token")
	}
	return token
}

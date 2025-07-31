package middleware

import (
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/edorguez/payment-reminder/pkg/firebase"
	"github.com/gin-gonic/gin"
)

const FirebaseUIDKey = "firebase_uid"

type FirebaseClaims struct {
	FirbaseUID string
	Email      string
	Name       string
}

func ExtractClaims(c *gin.Context) (*FirebaseClaims, bool) {
	decoded, ok := c.Get("decodedToken")
	if !ok {
		return nil, false
	}
	t := decoded.(*auth.Token)

	firebaseUID, _ := t.Claims["user_id"].(string)
	email, _ := t.Claims["email"].(string)
	name, _ := t.Claims["name"].(string)

	return &FirebaseClaims{
		FirbaseUID: firebaseUID,
		Email:      email,
		Name:       name,
	}, true
}

func FirebaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Grab the Authorization header.
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}

		// 2. Strip the "Bearer " prefix.
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. Verify the token with Firebase.
		decoded, err := firebase.Client().VerifyIDToken(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		// 4. Store the UID (for existing code) and the full token (for ExtractClaims).
		c.Set(FirebaseUIDKey, decoded.UID)
		c.Set("decodedToken", decoded)
		c.Next()
	}
}

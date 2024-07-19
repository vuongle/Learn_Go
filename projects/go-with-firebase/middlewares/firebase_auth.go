package middlewares

import (
	"context"
	"net/http"

	// "net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

const valName = "FIREBASE_ID_TOKEN"

// FirebaseAuthMiddleware is middleware for Firebase Authentication
type FirebaseAuthMiddleware struct {
	cli          *auth.Client
	unAuthorized func(c *gin.Context)
}

// New is constructor of the middleware
func New(credFileName string, unAuthorized func(c *gin.Context)) (*FirebaseAuthMiddleware, error) {
	opt := option.WithCredentialsFile(credFileName)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return &FirebaseAuthMiddleware{
		cli:          auth,
		unAuthorized: unAuthorized,
	}, nil
}

// MiddlewareFunc is function to verify token
func (fam *FirebaseAuthMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token from query string
		firebaseTk := c.Query("auth-token")
		idToken, err := fam.cli.VerifyIDToken(context.Background(), firebaseTk)
		if err != nil {
			if fam.unAuthorized != nil {
				fam.unAuthorized(c)
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"message": http.StatusText(http.StatusUnauthorized),
				})
			}
		}
		c.Set(valName, idToken) // set a value (firebase token) to the gin context (via the name "FIREBASE_ID_TOKEN")
		c.Next()

		// // Get token from header
		// authHeader := c.Request.Header.Get("Authorization")
		// _, token, found := strings.Cut(authHeader, " ")
		// if !found {
		// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		// 		"status":  http.StatusForbidden,
		// 		"message": http.StatusText(http.StatusForbidden),
		// 	})
		// 	return
		// }
		// idToken, err := fam.cli.VerifyIDToken(context.Background(), token)
		// if err != nil {
		// 	if fam.unAuthorized != nil {
		// 		fam.unAuthorized(c)
		// 	} else {
		// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// 			"status":  http.StatusUnauthorized,
		// 			"message": http.StatusText(http.StatusUnauthorized),
		// 		})
		// 	}
		// }
		// c.Set(valName, idToken)
		// c.Next()
	}
}

// ExtractClaims extracts claims
func ExtractClaims(c *gin.Context) *auth.Token {
	idToken, ok := c.Get(valName)
	if !ok {
		return new(auth.Token)
	}
	return idToken.(*auth.Token)
}

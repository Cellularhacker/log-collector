package tokenAuth

import (
	"context"
	"errors"
	"log-collector/util/token"
	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/julienschmidt/httprouter"
)

type key int

const (
	//UserIDKey is used in context
	UserIDKey             key = iota
	keyNotAuthorized          = "Not Authorized"
	keyAccessTokenCapital     = "Access-Token"
	keyAccessTokenLower       = "access-token"
)

//Auth middleware checks for token and validates, if invalid returns error
func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenString := r.Header.Get(keyAccessTokenCapital)
		if tokenString == "" {
			tokenString = r.Header.Get(keyAccessTokenLower)
		}
		if tokenString == "" {
			http.Error(w, keyNotAuthorized, 401)
			return
		}

		claims, isValid := token.Validate(tokenString)
		if !isValid {
			http.Error(w, keyNotAuthorized, 401)
			return
		}

		uid := claims["id"]
		if uid == nil {
			http.Error(w, keyNotAuthorized, 401)
			return
		}

		userIDHex := uid.(string)
		if !bson.IsObjectIdHex(userIDHex) {
			http.Error(w, keyNotAuthorized, 401)
			return
		}

		userID := bson.ObjectIdHex(userIDHex)

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next(w, r.WithContext(ctx), ps)
	}
}

func GetUserIDFromToken(tokenString string) (string, error) {
	claims, isValid := token.Validate(tokenString)
	if !isValid {
		return "", errors.New("invalid token")
	}
	return claims["id"].(string), nil
}

func GetUserID(r *http.Request) bson.ObjectId {
	return r.Context().Value(UserIDKey).(bson.ObjectId)
}

func GetIfAvailable(r *http.Request) (userID bson.ObjectId, exists bool) {
	tokenString := r.Header.Get(keyAccessTokenCapital)
	if tokenString == "" {
		tokenString = r.Header.Get(keyAccessTokenLower)
	}
	if tokenString == "" {
		return "", false
	}
	uid, _ := GetUserIDFromToken(tokenString)
	if bson.IsObjectIdHex(uid) {
		return bson.ObjectIdHex(uid), true
	}
	return "", false
}

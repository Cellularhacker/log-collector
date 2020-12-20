package parseInfo

import (
	"errors"
	"github.com/globalsign/mgo/bson"
)

var store Store

const (
	TypeGoogleFirebase = "firebase"
	TypeFacebook       = "facebook"
	KKeyCookie         = "cookie"
	KKeySession        = "session"

	KeyID                  = "_id"
	KeyType                = "type"
	KeyKKey                = "key"
	KeyValue               = "value"
	KeyCreatedAt           = "created_at"
	KeyIsActive            = "is_active"
	KeyUpdatedAt           = "updated_at"
	KeyFirstFetchedAt      = "first_fetched_at"
	KeyLastFetchedAt       = "last_fetched_at"
	KeyLastActiveChangedAt = "last_active_changed_at"
	KeyFetchedCount        = "fetched_count"
	KeyUsedCount           = "used_count"
)

type Store interface {
	Create(a *ParseInfo) error
	GetBy(bson.M) ([]ParseInfo, error)
	GetByDesc(bson.M) ([]ParseInfo, error)
	GetBySortLimit(query bson.M, sort string, limit int) ([]ParseInfo, error)
	UpdateSet(what, set bson.M) error
	Delete(what bson.M, all bool) error
	CountBy(bson.M) (int, error)
}

func SetStore(s Store) {
	store = s
}

func GetLatestFirebaseCookie() (*ParseInfo, error) {
	pi, err := store.GetBySortLimit(bson.M{KeyType: TypeGoogleFirebase, KeyKKey: KKeyCookie, KeyIsActive: true}, "-created_at", 1)
	if err != nil {
		return nil, err
	}
	if len(pi) > 0 {
		return &pi[0], nil
	}

	return nil, errors.New("parseInfo missing")
}

func GetByID(id bson.ObjectId) (*ParseInfo, error) {
	as, err := store.GetByDesc(bson.M{KeyID: id})
	if err != nil {
		return nil, err
	}
	if len(as) > 0 {
		return &as[0], nil
	}
	return nil, errors.New("key missing")
}

func GetLatestOne() (*ParseInfo, error) {
	keys, err := store.GetByDesc(nil)
	if err != nil {
		return nil, err
	}
	if len(keys) > 0 {
		return &keys[0], nil
	}
	return nil, errors.New("parseInfo missing")
}

func GetLatestOneByType(_type string) (*ParseInfo, error) {
	cs, err := store.GetByDesc(bson.M{"type": _type})
	if err != nil {
		return nil, err
	}
	if len(cs) > 0 {
		return &cs[0], nil
	}
	return nil, errors.New("parseInfo missing")
}

func GetLast24ByType(_type string) ([]ParseInfo, error) {
	return store.GetBySortLimit(bson.M{"type": _type}, "-gte", 24)
}

func GetBySortLimit(by bson.M, sort string, limit int) ([]ParseInfo, error) {
	return store.GetBySortLimit(by, sort, limit)
}

func GetAll() ([]ParseInfo, error) {
	return store.GetBy(nil)
}

func DeleteBy(by bson.M) error {
	return store.Delete(by, true)
}

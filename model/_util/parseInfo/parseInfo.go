package parseInfo

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type ParseInfo struct {
	ID                  bson.ObjectId `bson:"_id" json:"id"`
	Type                string        `bson:"type" json:"type"`
	Key                 string        `bson:"key" json:"key"`
	Value               string        `bson:"value" json:"value"`
	IsActive            bool          `bson:"is_active" json:"is_active"`
	CreatedAt           int64         `bson:"created_at" json:"created_at"`
	UpdatedAt           int64         `bson:"updated_at" json:"updated_at"`
	FirstFetchedAt      int64         `bson:"first_fetched_at" json:"first_fetched_at"`
	LastFetchedAt       int64         `bson:"last_fetch_at" json:"last_fetched_at"`
	LastActiveChangedAt int64         `bson:"last_active_changed_at" json:"last_active_changed_at"`
	FetchedCount        int           `bson:"fetched_count" json:"fetched_count"`
	UsedCount           int           `bson:"used_count" json:"used_count"`
}

func New() *ParseInfo {
	return &ParseInfo{}
}

func (a *ParseInfo) Create() error {
	a.CreatedAt = time.Now().Unix()
	a.FirstFetchedAt = 0
	a.LastFetchedAt = 0
	a.UpdatedAt = a.CreatedAt
	return store.Create(a)
}

func (a *ParseInfo) UpdateFetched() error {
	now := time.Now().Unix()
	set := bson.M{}
	a.FetchedCount += 1

	if a.FirstFetchedAt <= 0 { //Only update when it is firstly using.
		set[KeyFirstFetchedAt] = now
	}
	set[KeyLastFetchedAt] = now
	set[KeyFetchedCount] = a.FetchedCount

	return store.UpdateSet(bson.M{KeyID: a.ID}, set)
}

func (a *ParseInfo) Active() error {
	return a.setActiveOrNot(true)
}

func (a *ParseInfo) InActive() error {
	return a.setActiveOrNot(false)
}

func (a *ParseInfo) setActiveOrNot(isActive bool) error {
	a.UsedCount++
	now := time.Now().Unix()
	set := bson.M{KeyIsActive: isActive}
	if a.IsActive != isActive {
		set[KeyLastActiveChangedAt] = now
	}

	return store.UpdateSet(bson.M{KeyID: a.ID}, set)
}

package storages

import "sync"

type RatingStore interface {
	Add(laptopID string, score float64) (*Rating, error)
}

type Rating struct {
	Count uint32
	Sum   float64
}

type InMemoryRatingStore struct {
	mutex   sync.Mutex
	ratings map[string]*Rating
}

func NewInMemoryRatingStore() *InMemoryRatingStore {
	return &InMemoryRatingStore{
		ratings: make(map[string]*Rating),
	}
}

// Add adds a rating for a laptop with the given ID and score.
//
// Parameters:
// - laptopID: the ID of the laptop to add the rating for.
// - score: the score to add for the laptop.
//
// Returns:
// - *Rating: the updated rating for the laptop, or nil if the laptop does not exist.
// - error: an error if the operation failed.
func (store *InMemoryRatingStore) Add(laptopID string, score float64) (*Rating, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	// find a rating in the map by laptop ID
	rating := store.ratings[laptopID]

	// If not exist -> create a new one
	// else increase the count and sum then update the map
	if rating == nil {
		rating = &Rating{
			Count: 1,
			Sum:   score,
		}
	} else {
		rating.Count++
		rating.Sum += score
	}

	store.ratings[laptopID] = rating

	return rating, nil
}

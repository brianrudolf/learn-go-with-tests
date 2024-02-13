package context

import (
	"fmt"
	"net/http"
)

// type StubStore struct {
// 	response string
// }

// func (s *StubStore) Fetch() string {
// 	return s.response
// }

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()

		// data := make(chan (string, error), 1)

		// go func() {
		// 	data, err <- store.Fetch()
		// }()

		// select {
		// case d := <-data:
		// 	fmt.Fprint(w, d)
		// case <-ctx.Done():
		// 	store.Cancel()
		// }
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log something
		}
		fmt.Fprint(w, data)
	}
}

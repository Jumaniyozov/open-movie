package data

import (
	"encoding/json"
	"fmt"
	"time"
)

// Note that there are no struct tags on the Movie struct itself.
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"-"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func (m Movie) MarshalJSON() ([]byte, error) {
	// Create a variable holding the custom runtime string, just like before.
	var runtime string
	if m.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}
	// Define a MovieAlias type which has the underlying type Movie. Due to the way that
	// Go handles type definitions (https://golang.org/ref/spec#Type_definitions) the
	// MovieAlias type will contain all the fields that our Movie struct has but,
	// importantly, none of the methods.
	type MovieAlias Movie
	// Embed the MovieAlias type inside the anonymous struct, along with a Runtime field
	// that has the type string and the necessary struct tags. It's important that we
	// embed the MovieAlias type here, rather than the Movie type directly, to avoid
	// inheriting the MarshalJSON() method of the Movie type (which would result in an
	// infinite loop during encoding).
	aux := struct {
		MovieAlias
		Runtime string `json:"runtime,omitempty"`
	}{
		MovieAlias: MovieAlias(m),
		Runtime:    runtime,
	}
	return json.Marshal(aux)
}
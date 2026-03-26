package hackerapi

import "time"

type SearchBasedResp  struct {
	Exhaustive struct {
		NbHits bool `json:"nbHits"`
		Typo   bool `json:"typo"`
	} `json:"exhaustive"`
	ExhaustiveNbHits bool `json:"exhaustiveNbHits"`
	ExhaustiveTypo   bool `json:"exhaustiveTypo"`
	Hits             []struct {
		HighlightResult struct {
			Author struct {
				MatchLevel   string `json:"matchLevel"`
				MatchedWords []any  `json:"matchedWords"`
				Value        string `json:"value"`
			} `json:"author"`
			Title struct {
				FullyHighlighted bool     `json:"fullyHighlighted"`
				MatchLevel       string   `json:"matchLevel"`
				MatchedWords     []string `json:"matchedWords"`
				Value            string   `json:"value"`
			} `json:"title"`
			URL struct {
				MatchLevel   string `json:"matchLevel"`
				MatchedWords []any  `json:"matchedWords"`
				Value        string `json:"value"`
			} `json:"url"`
		} `json:"_highlightResult,omitempty"`
		Tags             []string  `json:"_tags"`
		Author           string    `json:"author"`
		Children         []int     `json:"children"`
		CreatedAt        time.Time `json:"created_at"`
		CreatedAtI       int       `json:"created_at_i"`
		NumComments      int       `json:"num_comments"`
		ObjectID         string    `json:"objectID"`
		Points           int       `json:"points"`
		StoryID          int       `json:"story_id"`
		Title            string    `json:"title"`
		UpdatedAt        time.Time `json:"updated_at"`
		URL              string    `json:"url,omitempty"`
		HighlightResult0 struct {
			Author struct {
				MatchLevel   string `json:"matchLevel"`
				MatchedWords []any  `json:"matchedWords"`
				Value        string `json:"value"`
			} `json:"author"`
			StoryText struct {
				FullyHighlighted bool     `json:"fullyHighlighted"`
				MatchLevel       string   `json:"matchLevel"`
				MatchedWords     []string `json:"matchedWords"`
				Value            string   `json:"value"`
			} `json:"story_text"`
			Title struct {
				FullyHighlighted bool     `json:"fullyHighlighted"`
				MatchLevel       string   `json:"matchLevel"`
				MatchedWords     []string `json:"matchedWords"`
				Value            string   `json:"value"`
			} `json:"title"`
		} `json:"_highlightResult,omitempty"`
		StoryText        string `json:"story_text,omitempty"`
		HighlightResult1 struct {
			Author struct {
				MatchLevel   string `json:"matchLevel"`
				MatchedWords []any  `json:"matchedWords"`
				Value        string `json:"value"`
			} `json:"author"`
			StoryText struct {
				FullyHighlighted bool     `json:"fullyHighlighted"`
				MatchLevel       string   `json:"matchLevel"`
				MatchedWords     []string `json:"matchedWords"`
				Value            string   `json:"value"`
			} `json:"story_text"`
			Title struct {
				FullyHighlighted bool     `json:"fullyHighlighted"`
				MatchLevel       string   `json:"matchLevel"`
				MatchedWords     []string `json:"matchedWords"`
				Value            string   `json:"value"`
			} `json:"title"`
		} `json:"_highlightResult,omitempty"`
	} `json:"hits"`
	HitsPerPage         int    `json:"hitsPerPage"`
	NbHits              int    `json:"nbHits"`
	NbPages             int    `json:"nbPages"`
	Page                int    `json:"page"`
	Params              string `json:"params"`
	ProcessingTimeMS    int    `json:"processingTimeMS"`
	ProcessingTimingsMS struct {
		Request struct {
			RoundTrip int `json:"roundTrip"`
		} `json:"_request"`
		Fetch struct {
			Query int `json:"query"`
			Total int `json:"total"`
		} `json:"fetch"`
		Total int `json:"total"`
	} `json:"processingTimingsMS"`
	Query        string `json:"query"`
	ServerTimeMS int    `json:"serverTimeMS"`
}

type FrontPageRes struct {
	Exhaustive struct {
		NbHits bool `json:"nbHits"`
		Typo   bool `json:"typo"`
	} `json:"exhaustive"`
	ExhaustiveNbHits bool `json:"exhaustiveNbHits"`
	ExhaustiveTypo   bool `json:"exhaustiveTypo"`
	Hits             []struct {
		HighlightResult struct {
			Author struct {
				MatchLevel   string        `json:"matchLevel"`
				MatchedWords []interface{} `json:"matchedWords"`
				Value        string        `json:"value"`
			} `json:"author"`
			Title struct {
				MatchLevel   string        `json:"matchLevel"`
				MatchedWords []interface{} `json:"matchedWords"`
				Value        string        `json:"value"`
			} `json:"title"`
			URL struct {
				MatchLevel   string        `json:"matchLevel"`
				MatchedWords []interface{} `json:"matchedWords"`
				Value        string        `json:"value"`
			} `json:"url"`
		} `json:"_highlightResult"`
		Tags        []string  `json:"_tags"`
		Author      string    `json:"author"`
		Children    []int     `json:"children"`
		CreatedAt   time.Time `json:"created_at"`
		CreatedAtI  int       `json:"created_at_i"`
		NumComments int       `json:"num_comments"`
		ObjectID    string    `json:"objectID"`
		Points      int       `json:"points"`
		StoryID     int       `json:"story_id"`
		Title       string    `json:"title"`
		UpdatedAt   time.Time `json:"updated_at"`
		URL         string    `json:"url"`
		StoryText   string    `json:"story_text,omitempty"`
	} `json:"hits"`
	HitsPerPage         int    `json:"hitsPerPage"`
	NbHits              int    `json:"nbHits"`
	NbPages             int    `json:"nbPages"`
	Page                int    `json:"page"`
	Params              string `json:"params"`
	ProcessingTimeMS    int    `json:"processingTimeMS"`
	ProcessingTimingsMS struct {
		Request struct {
			RoundTrip int `json:"roundTrip"`
		} `json:"_request"`
		Total int `json:"total"`
	} `json:"processingTimingsMS"`
	Query        string `json:"query"`
	ServerTimeMS int    `json:"serverTimeMS"`
}
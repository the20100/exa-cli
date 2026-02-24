package exa

// --- Request types ---

type SearchRequest struct {
	Query              string          `json:"query"`
	NumResults         *int            `json:"num_results,omitempty"`
	StartPublishedDate *string         `json:"start_published_date,omitempty"`
	EndPublishedDate   *string         `json:"end_published_date,omitempty"`
	IncludeDomains     []string        `json:"include_domains,omitempty"`
	ExcludeDomains     []string        `json:"exclude_domains,omitempty"`
	IncludeText        []string        `json:"include_text,omitempty"`
	ExcludeText        []string        `json:"exclude_text,omitempty"`
	Category           *string         `json:"category,omitempty"`
	Type               *string         `json:"type,omitempty"`
	Contents           *ContentsOptions `json:"contents,omitempty"`
}

type FindSimilarRequest struct {
	URL                  string          `json:"url"`
	NumResults           *int            `json:"num_results,omitempty"`
	ExcludeSourceDomain  bool            `json:"exclude_source_domain,omitempty"`
	IncludeDomains       []string        `json:"include_domains,omitempty"`
	ExcludeDomains       []string        `json:"exclude_domains,omitempty"`
	StartPublishedDate   *string         `json:"start_published_date,omitempty"`
	EndPublishedDate     *string         `json:"end_published_date,omitempty"`
	Contents             *ContentsOptions `json:"contents,omitempty"`
}

type GetContentsRequest struct {
	IDs      []string         `json:"ids"`
	Contents *ContentsOptions `json:"contents,omitempty"`
}

type AnswerRequest struct {
	Query        string  `json:"query"`
	Model        *string `json:"model,omitempty"`
	Text         *bool   `json:"text,omitempty"`
	SystemPrompt *string `json:"system_prompt,omitempty"`
	UserLocation *string `json:"user_location,omitempty"`
}

type ResearchRequest struct {
	Instructions string `json:"instructions"`
}

// --- Contents options ---

type ContentsOptions struct {
	Text        *TextOptions       `json:"text,omitempty"`
	Summary     *SummaryOptions    `json:"summary,omitempty"`
	Highlights  *HighlightOptions  `json:"highlights,omitempty"`
	Livecrawl   *string            `json:"livecrawl,omitempty"`
	MaxAgeHours *int               `json:"max_age_hours,omitempty"`
	Subpages    *int               `json:"subpages,omitempty"`
}

type TextOptions struct {
	MaxCharacters *int   `json:"max_characters,omitempty"`
	IncludeHtmlTags *bool `json:"include_html_tags,omitempty"`
}

type SummaryOptions struct {
	Query *string `json:"query,omitempty"`
}

type HighlightOptions struct {
	NumSentences      *int `json:"num_sentences,omitempty"`
	HighlightsPerURL  *int `json:"highlights_per_url,omitempty"`
}

// --- Response types ---

type SearchResponse struct {
	Results            []Result `json:"results"`
	ResolvedSearchType *string  `json:"resolved_search_type,omitempty"`
	AutoDate           *string  `json:"auto_date,omitempty"`
	RequestID          *string  `json:"requestId,omitempty"`
}

type Result struct {
	URL           string   `json:"url"`
	ID            string   `json:"id"`
	Title         *string  `json:"title,omitempty"`
	Score         *float64 `json:"score,omitempty"`
	PublishedDate *string  `json:"published_date,omitempty"`
	Author        *string  `json:"author,omitempty"`
	Image         *string  `json:"image,omitempty"`
	Favicon       *string  `json:"favicon,omitempty"`
	Text          *string  `json:"text,omitempty"`
	Summary       *string  `json:"summary,omitempty"`
	Highlights    []string `json:"highlights,omitempty"`
}

type AnswerResponse struct {
	Answer    interface{}    `json:"answer"`
	Citations []AnswerResult `json:"citations,omitempty"`
	RequestID *string        `json:"requestId,omitempty"`
}

type AnswerResult struct {
	URL           string   `json:"url"`
	ID            string   `json:"id"`
	Title         *string  `json:"title,omitempty"`
	PublishedDate *string  `json:"published_date,omitempty"`
	Author        *string  `json:"author,omitempty"`
}

type ResearchResponse struct {
	Data      interface{} `json:"data"`
	RequestID *string     `json:"requestId,omitempty"`
}

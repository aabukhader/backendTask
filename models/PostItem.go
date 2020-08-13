package models

// SearchBody struct
type SearchBody struct {
	Status string `json:"status"`
	Data   *Data  `json:"data"`
}

// Data search body struct
type Data struct {
	Q                    string            `json:"q"`
	TotalCount           int               `json:"total_count"`
	SpellingAlternatives map[string]string `json:"spelling_alternatives"`
	Items                []*PostItem       `json:"items"`
}

// PostItem Struct
type PostItem struct {
	ID                    int               `json:"id"`
	Name                  string            `json:"name"`
	Type                  string            `json:"type"`
	URL                   string            `json:"url"`
	GlobalID              string            `json:"global_id"`
	Description           string            `json:"description"`
	PublishedAt           string            `json:"published_at"`
	Publisher             string            `json:"publisher"`
	CitationHTML          string            `json:"citationHtml"`
	IdentifierOfDataverse string            `json:"identifier_of_dataverse"`
	Citation              string            `json:"citation"`
	StorageIdentifier     string            `json:"storageIdentifier"`
	NameOfDataverse       string            `json:"name_of_dataverse"`
	Keywords              map[string]string `json:"keywords"`
	Subjects              map[string]string `json:"subjects"`
	FileCount             int               `json:"fileCount"`
	VersionID             string            `json:"versionId"`
	VersionState          string            `json:"versionState"`
	MajorVersion          int               `json:"majorVersion"`
	MinorVersion          int               `json:"minorVersion"`
	CreatedAt             string            `json:"createdAt"`
	UpdatedAt             string            `json:"updatedAt"`
	Contacts              map[string]string `json:"contacts"`
	Authors               map[string]string `json:"authers"`
}


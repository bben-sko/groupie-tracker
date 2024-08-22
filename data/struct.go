package d

type Relation struct {
	DatesLocations map[string]interface{} `json:"datesLocations"`
}
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type AutoGenerated struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}
type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

type SearchResult struct {
	Image        string `json:"image"`
	CreationDate int    `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

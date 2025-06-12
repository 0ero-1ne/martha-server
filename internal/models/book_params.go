package models

type BookUrlParams struct {
	Limit         int
	Offset        int
	WithTags      bool
	WithAuthors   bool
	WithComments  bool
	WithChapters  bool
	WithBookRates bool
	Query         string
	Tags          string
	Statuses      string
	StartYear     int
	EndYear       int
}

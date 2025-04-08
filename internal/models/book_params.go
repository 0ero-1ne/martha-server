package models

type BookUrlParams struct {
	Limit        int
	Offset       int
	WithTags     bool
	WithAuthors  bool
	WithComments bool
	WithChapters bool
}

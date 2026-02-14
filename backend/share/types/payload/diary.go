package payload

type DiaryOverviewQuery struct {
	Year *int `json:"year_no" validate:"min=2001,max=2100"`
}

type DiaryDayGraph struct {
	Date    *string `json:"date"`
	Title   *string `json:"title"`
	Color   *int    `json:"graph"`
	Summary *string `json:"summary"`
}

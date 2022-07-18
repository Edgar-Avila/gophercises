package story

type StoryArc struct {
    Title string `json:"title"`
    Story []string `json:"story"`
    Options []Option `json:"options"`
}

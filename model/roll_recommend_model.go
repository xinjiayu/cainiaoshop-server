package model

//RollRecommendModel 首页滚动推荐实体
type RollRecommendModel struct {
	ID       string `json:"id"`
	ImageURL string `json:"image_url"`
	Dest     string `json:"dest"`
}

package model

//TypeRecommend 类别推荐
type TypeRecommend struct {
	RecommendModel [3]*RollRecommendModel `json:"recommend"`
}

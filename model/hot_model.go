package model

//HotModel 热门商品
type HotModel struct {
	ID     string  `json:"id"`
	Price  float32 `json:"price"`
	PicURL string  `json:"pic_url"`
	Dest   string  `json:"dest"`
}

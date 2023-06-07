package Model

type Data struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Percent string `json:"percent"`
}

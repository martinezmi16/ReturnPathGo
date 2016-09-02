package api

//Creates animal type that will be encoded into JSON
type Animal struct {

    Id int `json:"id"`
    Name string `json:"name"`
    LegCount int `json:"leg_count"`
    LifeSpan int `json:"lifespan"`
    IsEndangered bool `json:"is_endangered"`
    CreatedTime string `json:"created_at"`
    UpdatedTime string `json:"updated_at"`

}

type Animals []Animal



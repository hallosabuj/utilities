package calculator

type AddNumbers struct {
	Numbers []float32 `json:"numbers"`
}
type AddResult struct {
	Numbers []float32 `json:"numbers"`
	Sum     float32   `json:"sum"`
}

type SubNumbers struct {
	Number1 float32 `json:"num1"`
	Number2 float32 `json:"num2"`
}
type SubResult struct {
	Number1    float32 `json:"num1"`
	Number2    float32 `json:"num2"`
	Difference float32 `json:"differnece"`
}

type MultNumbers struct {
	Numbers []float32 `json:"numbers"`
}
type MultResult struct {
	Numbers []float32 `json:"numbers"`
	Product float32   `json:"product"`
}

type DivNumbers struct {
	Number1 float32 `json:"num1"`
	Number2 float32 `json:"num2"`
}
type DivResult struct {
	Number1  float32 `json:"num1"`
	Number2  float32 `json:"num2"`
	Quotient float32 `json:"quotient"`
}

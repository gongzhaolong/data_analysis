package constant

type Data struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    []struct {
		MaxCpu []struct {
			Data1 []float64 `json:"data"`
		} `json:"max_cpu"`
		AvgCpu []struct {
			Data1 []float64 `json:"data"`
		} `json:"avg_cpu"`
		MaxMem []struct {
			Data1 []float64 `json:"data"`
		} `json:"max_mem"`
		AvgMem []struct {
			Data1 []float64 `json:"data"`
		} `json:"avg_mem"`
	} `json:"data"`
}

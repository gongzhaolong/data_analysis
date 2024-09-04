package functions

import (
	"math"
	"sort"
)

// 计算最大值
func Max(data []float64) float64 {
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// 计算均值
func Mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// 计算方差
func Variance(data []float64) float64 {
	sum := 0.0
	mean := Mean(data)
	for _, value := range data {
		sum += (value - mean) * (value - mean)
	}
	return sum / float64(len(data))
}

// 计算相对标准差（变异系数）
func RelativeStandardDeviation(data []float64) float64 {
	std := math.Sqrt(Variance(data))
	mean := Mean(data)
	return std / mean
}

// 计算滚动相对标准差（变异系数）
func RollingVariance(data []float64, windowSize int) (rollingVars float64) {
	rollingVarsList := make([]float64, len(data)-windowSize+1)
	variance := RelativeStandardDeviation(data)
	for i := 0; i < len(rollingVarsList); i++ {
		window := data[i : i+windowSize]
		rollingVarsList[i] = RelativeStandardDeviation(window)
	}
	for _, windows_var := range rollingVarsList {
		if windows_var > variance {
			rollingVars += windows_var - variance
		} else {
			rollingVars += variance - windows_var
		}
	}
	return rollingVars
}

// 求全局最大值
func MaxVal(data []float64) (max_val float64, max_time int) {
	max_val = data[0]
	max_time = 0
	for index, val := range data {
		if val > max_val {
			max_val = val
			max_time = index
		}
	}
	return
}

// 求全局最小值
func MinVal(data []float64) (min_val float64, min_time int) {
	min_val = data[0]
	min_time = 0
	for index, val := range data {
		if val < min_val {
			min_val = val
			min_time = index
		}
	}
	return
}

// normalize 函数将切片数据归一化到 [0, 1] 之间
func Normalize(data []float64) []float64 {
	minVal := math.MaxFloat64
	maxVal := -math.MaxFloat64

	// 找到切片中的最小值和最大值
	for _, value := range data {
		if value < minVal {
			minVal = value
		}
		if value > maxVal {
			maxVal = value
		}
	}
	//如果数据本身已经在[0,1]不需要归一化
	if maxVal <= 1 && minVal >= 0 {
		return data
	}
	// 归一化切片中的每个值
	normalized := make([]float64, len(data))
	for i, value := range data {
		if maxVal != minVal {
			normalized[i] = (value - minVal) / (maxVal - minVal)
		}
	}

	return normalized
}

//求单调递增、单调递减区间

// 区间的起始和结束位置
type Interval struct {
	Start int
	End   int
}

// ByStart 存储区间
type ByStart []Interval

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }

// findMonotonicRanges 寻找float64切片中的单调递增和单调递减区间
func FindMonotonicRanges(arr []float64) (increasing, decreasing []Interval) {
	if len(arr) == 0 {
		return
	}

	start := 0
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i-1]-arr[i] >= 0 {
			// Found the end of an increasing range
			if i-1 > start {
				increasing = append(increasing, Interval{Start: start, End: i - 1})
			}
			start = i

		}
	}

	start = 0
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i]-arr[i-1] >= 0 {
			// Found the end of a decreasing range
			if i-1 > start {
				decreasing = append(decreasing, Interval{Start: start, End: i - 1})
			}
			start = i
		}
	}
	sort.Sort(ByStart(increasing))
	sort.Sort(ByStart(decreasing))
	return
}

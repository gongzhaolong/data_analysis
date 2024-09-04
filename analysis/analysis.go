package analysis

import (
	"data_analysis/constant"
	"data_analysis/functions"
	"data_analysis/generate_number"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var IncSharpTimes int
var IncSlightSharpTimes int
var IncDiscSlowTimes int
var DecSharpTimes int
var DecSlightSharpTimes int
var DecDiscSlowTimes int

func Analyze(origin_data []float64) (grade int, catgory int) {
	//step1 平滑数据,去除毛刺
	smooth_data := Smooth(origin_data)
	//step2 平稳性检测
	is_stable := isStable(smooth_data) //计算全局、滚动相对标准差（变异系数）判断是否平稳

	//step3 非平稳数据单调性趋势判断，输出数据的特征，得出结论
	if !is_stable {
		normalized_data := functions.Normalize(smooth_data) //数据归一化到[0,1]
		grade, catgory = incAndDecRange(normalized_data)
	} else {
		grade = constant.Risk0
		catgory = constant.C0
	}
	//step4 返回风险等级和整体趋势
	return
}

// 平滑曲线
func Smooth(data []float64) []float64 {
	for i := 1; i < len(data)-1; i++ {
		//去除毛刺
		if data[i] > 2*(data[i+1]+data[i-1])/2 {
			data[i] = (data[i+1] + data[i-1]) / 2
		}
	}
	for i := 1; i < len(data)-1; i++ {
		//平滑曲线
		if (data[i] >= data[i+1] && data[i] >= data[i-1]) || (data[i] <= data[i+1] && data[i] <= data[i-1]) {
			data[i] = (data[i+1] + data[i-1]) / 2
		}
	}
	return data
}

// 平稳性检测
func isStable(columns []float64) (is_stable bool) {
	MaxValue := functions.Max(columns)
	if MaxValue <= constant.MaxThreshold {
		return true
	} //如果最大值小于0.4则认为无风险
	RelativeStd := functions.RelativeStandardDeviation(columns)
	RollRelativeStd := functions.RollingVariance(columns, constant.WindowSize)
	if RelativeStd < constant.StableVar && RollRelativeStd < constant.StableRollingVar {
		//max_val, max_time := functions.MaxVal(columns)
		//min_val, min_time := functions.MinVal(columns)
		//mean := functions.Mean(columns)
		//variance := functions.Variance(columns)
		//fmt.Println("该时序数据为平稳序列:")
		//fmt.Printf("均值: %.4f\n", mean)
		//fmt.Printf("方差: %.4f\n", variance)
		//fmt.Printf("全局峰值: %.4f 采样时刻: %v \n", max_val, max_time)
		//fmt.Printf("全局谷值: %.4f 采样时刻: %v \n", min_val, min_time)
		return true
	}
	return false
}

// 单调递增、递减区间描述
func incAndDecRange(smooth_data []float64) (grade int, catgory int) {
	increasingRanges, decreasingRanges := functions.FindMonotonicRanges(smooth_data)
	for _, interval := range increasingRanges {
		//这里加入了更加精细地判断，4个采样点计算一次斜率
		start, end := interval.Start, interval.End
		if end-start+1 <= 4 {
			TraversalInc(smooth_data, start, end)
			continue
		}
		for i := start; i <= end-3; i += 3 {
			TraversalInc(smooth_data, i, i+3)
		}
	}
	for _, interval := range decreasingRanges {
		//这里加入了更加精细地判断，4个采样点计算一次斜率
		start, end := interval.Start, interval.End
		if end-start+1 <= 4 {
			TraversalDec(smooth_data, start, end)
			continue
		}
		for i := start; i <= end-3; i += 3 {
			TraversalDec(smooth_data, i, i+3)
		}
	}
	grade, catgory = Output()

	fmt.Printf("IncSharpTimes:    %d\n", IncSharpTimes)
	fmt.Printf("IncSlightSharpTimes: %d\n", IncSlightSharpTimes)
	fmt.Printf("IncDiscSlowTimes:   %d\n", IncDiscSlowTimes)
	fmt.Printf("DecSharpTimes: %d\n", DecSharpTimes)
	fmt.Printf("DecSlightSharpTimes:%d\n", DecSlightSharpTimes)
	fmt.Printf("DecDiscSlowTimes:   %d\n", DecDiscSlowTimes)
	IncSharpTimes, IncSlightSharpTimes, IncDiscSlowTimes = 0, 0, 0
	DecSharpTimes, DecSlightSharpTimes, DecDiscSlowTimes = 0, 0, 0
	return
}

// 针对每个区间作进一步分析
func TraversalInc(smooth_data []float64, start, end int) {
	k := (smooth_data[end] - smooth_data[start]) / float64(end-start+1) * 100
	//rateChange := functions.RateChange(k)
	if k >= constant.Sharp {
		IncSharpTimes++
		fmt.Printf("k:%v", k)
		fmt.Printf("[%d, %d] 上升\n", start+1, end+1)
		//fmt.Printf("k:%v", k)
		//fmt.Printf("[%d, %d] 增长趋势:%v  区间极大值:%.4f (可能存在风险！)\n", start+1, end+1, rateChange, origin_data[end])
	} else if k >= constant.SlightlySharp && k < constant.Sharp {
		IncSlightSharpTimes++
		fmt.Printf("k:%v", k)
		fmt.Printf("[%d, %d] 上升\n", start+1, end+1)
		//fmt.Printf("k:%v", k)
		//fmt.Printf("[%d, %d] 增长趋势:%v  区间极大值:%.4f \n", start+1, end+1, rateChange, origin_data[end])
	} else if k >= constant.SlightlySlow && k < constant.SlightlySharp {
		IncDiscSlowTimes++
		fmt.Printf("k:%v", k)
		fmt.Printf("[%d, %d] 上升\n", start+1, end+1)
		//fmt.Printf("k:%v", k)
		//fmt.Printf("[%d, %d] 增长趋势:%v  区间极大值:%.4f \n", start+1, end+1, rateChange, origin_data[end])
	} else {
		fmt.Printf("k:%v", k)
		fmt.Printf("[%d, %d] 下降\n", start+1, end+1)
	}
}

func TraversalDec(smooth_data []float64, start, end int) {
	k := -(smooth_data[end] - smooth_data[start]) / float64(end-start+1) * 100
	if k >= constant.Sharp {
		DecSharpTimes++
		fmt.Printf("k:%v", k)
		fmt.Printf("[%d, %d] 下降\n", start+1, end+1)
	} else if k >= constant.SlightlySharp && k < constant.Sharp {
		DecSlightSharpTimes++
		fmt.Printf("k:%v", k)
		fmt.Printf("[%d, %d] 下降\n", start+1, end+1)
	} else if k >= constant.SlightlySlow && k < constant.SlightlySharp {
		DecDiscSlowTimes++
		fmt.Printf("k:%v", k)
		fmt.Printf("[%d, %d] 下降\n", start+1, end+1)
	} else {
		fmt.Printf("k:%v", k)
		fmt.Printf("[%d, %d] 下降\n", start+1, end+1)
	}
}

func Output() (risk int, catgory int) {
	if (IncSharpTimes > 0 || IncSlightSharpTimes > 0 || IncDiscSlowTimes > 0) && (DecSharpTimes == 0 && DecSlightSharpTimes == 0 && DecDiscSlowTimes <= 1) {
		catgory = constant.C1
	} else if (DecSharpTimes > 0 || DecSlightSharpTimes > 0 || DecDiscSlowTimes > 0) && (IncSharpTimes == 0 && IncSlightSharpTimes == 0 && IncDiscSlowTimes <= 1) {
		catgory = constant.C2
	} else {
		catgory = constant.C0
	}
	switch {
	case IncSharpTimes > 0 || DecSharpTimes > 0 || IncSlightSharpTimes > 1 || IncDiscSlowTimes > 2:
		risk = constant.Risk4
	case IncSlightSharpTimes > 0 || DecSlightSharpTimes > 0 || IncDiscSlowTimes > 1 || DecDiscSlowTimes > 2:
		risk = constant.Risk3
	case IncDiscSlowTimes > 0 || DecDiscSlowTimes > 0:
		risk = constant.Risk2
	default:
		risk = constant.Risk1
	}
	return
}

/*输出结果（示例输出）：
单调递增区间:
    [1, 5]   增长趋势:缓慢       区间极大值:0.7298
	[10, 14] 增长趋势:剧烈       区间极大值:0.7827 (可能存在风险！)
	[18, 19] 增长趋势:缓慢       区间极大值:0.6481
	[24, 28] 增长趋势:剧烈       区间极大值:0.9067 (可能存在风险！)
	[33, 40] 增长趋势:较缓慢      区间极大值:0.5860
单调递减区间:
	[5, 10]  下降趋势:剧烈        区间极小值:0.4664 (可能存在风险！)
	[14, 18] 下降趋势:较缓慢      区间极小值:0.6325
	[19, 24] 下降趋势:缓慢        区间极小值:0.5431
	[28, 33] 下降趋势:剧烈        区间极小值:0.3640 (可能存在风险！)
*/

func JsonDataTest(json_name string) {
	// 打开 JSON 文件
	jsonFile, err := os.Open(json_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	// 读取文件内容
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// 解析 JSON 数据
	var data constant.Data
	json.Unmarshal(byteValue, &data)

	//遍历json数据
	for i, value := range data.Data {
		fmt.Printf("%v:\n", i)
		for _, v := range value.MaxCpu {
			var orignal_data []float64
			copy(orignal_data, v.Data1)
			grade, catgory := Analyze(v.Data1)
			//fmt.Printf("%v ", grade)
			//fmt.Printf("%v ", grade)
			if grade > 0 {
				fmt.Printf("%v \n", grade)
				fmt.Printf("%v \n", catgory)
				Display(orignal_data, i, grade, "MaxCpu", catgory)
			}
		}
		for _, v := range value.AvgCpu {

			grade, catgory := Analyze(v.Data1)
			//fmt.Printf("%v ", grade)
			if grade > 0 {
				fmt.Printf("%v \n", grade)
				fmt.Printf("%v \n", catgory)
				Display(v.Data1, i, grade, "AvgCpu", catgory)
			}
		}
		for _, v := range value.MaxMem {
			grade, catgory := Analyze(v.Data1)
			//fmt.Printf("%v ", grade)
			if grade > 0 {
				fmt.Printf("%v \n", grade)
				fmt.Printf("%v \n", catgory)
				Display(v.Data1, i, grade, "MaxMem", catgory)
			}
		}
		for _, v := range value.AvgMem {
			grade, catgory := Analyze(v.Data1)
			//fmt.Printf("%v ", grade)
			if grade > 0 {
				fmt.Printf("%v \n", grade)
				fmt.Printf("%v \n", catgory)
				//	Display(v.Data1, i, grade, "AvgMem", catgory)
			}
		}
		fmt.Println()
	}
}

var count int

func Display(Data []float64, number, grade int, metric string, catgory int) {
	//smooth := Smooth(Data)
	//metric = metric + "smooth"
	count++
	if count <= 26 {
		generate_number.WriteToExcel(Data, number, grade, metric, catgory)
	}

}

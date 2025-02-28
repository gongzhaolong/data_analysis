package main

import (
	"data_analysis/analysis"
	"data_analysis/constant"
)

func main() {
	//读取40个时序数据（示例输入）
	//origin_data := []float64{
	//	0.176,
	//	0.20675,
	//	0.222125,
	//	0.2375,
	//	0.25675,
	//	0.250925,
	//	0.2480125,
	//	0.2451,
	//	0.2349,
	//	0.2323,
	//	0.231,
	//	0.2297,
	//	0.22675,
	//	0.2238,
	//	0.21195,
	//	0.2001,
	//	0.1984,
	//	0.21315,
	//	0.220525,
	//	0.2279,
	//	0.2853,
	//	0.3427,
	//	0.3734,
	//	0.4041,
	//	0.4106,
	//	0.4134,
	//	0.4148,
	//	0.41275,
	//	0.411725,
	//	0.4168125,
	//	0.41935625,
	//	0.415428125,
	//	0.413464063,
	//	0.4115,
	//	0.41075,
	//	0.41,
	//	0.34915,
	//	0.2883,
	//	0.2596,
	//	0.24525,
	//	0.2309,
	//}
	//origin_data := []float64{0.3, 1.4333, 4.15, 13.8167, 33.0056, 68.1333, 108.3, 139.8167,
	//	161.5722, 171.7278, 177.7333, 183.2222, 185.4333, 189.1389, 192.1111,
	//	189.1889, 200.0833, 193.1334, 189.65, 195.7667, 188.6166, 193.2389,
	//	191.4167, 185.8278, 184.8, 183.7166, 184.6833, 188.8333, 186.6, 188.9611,
	//	185.1333, 191.9333, 193.3167, 192.55, 194.5167, 183.1167, 183.1, 183.1, 183.1}

	//columns := generate_number.ReadData()
	//origin_data := columns[7]
	//fmt.Println(origin_data)
	//generate_number.WriteToExcel(origin_data)
	//数据分析,输出结果
	analysis.JsonDataTest(constant.JsonName)
	//grade, catgory := analysis.Analyze(origin_data)
	//fmt.Printf("grade is %v,catgory is %v\n", grade, catgory)

}

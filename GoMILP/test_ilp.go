//File  : test_ilp.go
//Author: 裴彬
//Date  : 2020/12/4

package ilp

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gonum.org/v1/gonum/mat"
	"math"
	"strconv"
	"time"
)

// TbRealMonitoring ...
type TbRealMonitoring struct {
	ID int `json:"id"`
	// Bm1C 1#锅炉蒸发量
	Bm1C float64 `json:"bm1_c"`
	// Bm1O 1#锅炉蒸发量优化值
	Bm1O float64 `json:"bm1_o"`
	// Bp1 1#锅炉压力值
	Bp1 float64 `json:"bp1"`
	// Bm2C 2#锅炉蒸发量
	Bm2C float64 `json:"bm2_c"`
	// Bm2O 2#锅炉蒸发量优化值
	Bm2O float64 `json:"bm2_o"`
	// Bp2 2#锅炉压力值
	Bp2 float64 `json:"bp2"`
	// Bm3C 3#锅炉蒸发量
	Bm3C float64 `json:"bm3_c"`
	// Bm3O 3#锅炉蒸发量优化值
	Bm3O float64 `json:"bm3_o"`
	// Bp3 3#锅炉压力值
	Bp3 float64 `json:"bp3"`
	// Bm4C 4#锅炉蒸发量
	Bm4C float64 `json:"bm4_c"`
	// Bm4O 4#锅炉蒸发量优化值
	Bm4O float64 `json:"bm4_o"`
	// Bp4 4#锅炉压力值
	Bp4 float64 `json:"bp4"`
	// Bm5C 5#锅炉蒸发量
	Bm5C float64 `json:"bm5_c"`
	// Bm5O 5#锅炉蒸发量优化值
	Bm5O float64 `json:"bm5_o"`
	// Bp5 5#锅炉压力值
	Bp5 float64 `json:"bp5"`
	// Bm6C 6#锅炉蒸发量
	Bm6C float64 `json:"bm6_c"`
	// Bm6O 6#锅炉蒸发量优化值
	Bm6O float64 `json:"bm6_o"`
	// Bp6 6#锅炉压力值
	Bp6 float64 `json:"bp6"`
	// Bm7C 7#锅炉蒸发量
	Bm7C float64 `json:"bm7_c"`
	// Bm7O 7#锅炉蒸发量优化值
	Bm7O float64 `json:"bm7_o"`
	// Bp7 7#锅炉#锅炉压力值值
	Bp7 float64 `json:"bp7"`
	// Tm1InC 1#汽机进汽量
	Tm1InC float64 `json:"tm1_in_c"`
	// Tm1InO 1#汽机进汽量优化值
	Tm1InO float64 `json:"tm1_in_o"`
	// Tp1C 1#汽机压力值
	Tp1C float64 `json:"tp1_c"`
	// Tcos1 1#汽机功率因子
	Tcos1 float64 `json:"tcos1"`
	// Tp1O 1#汽机压力优化值
	Tp1O float64 `json:"tp1_o"`
	// Tm1EC 1#汽机抽汽量
	Tm1EC float64 `json:"tm1_e_c"`
	// Tm1EO 1#汽机抽汽量优化值
	Tm1EO float64 `json:"tm1_e_o"`
	// Tm2InC 2#汽机进汽量
	Tm2InC float64 `json:"tm2_in_c"`
	// Tm2InO 2#汽机进汽量优化值
	Tm2InO float64 `json:"tm2_in_o"`
	// Tp2C 2#汽机压力值
	Tp2C float64 `json:"tp2_c"`
	// Tcos2 2#汽机功率因子
	Tcos2 float64 `json:"tcos2"`
	// Tp2O 2#汽机压力优化值
	Tp2O float64 `json:"tp2_o"`
	// Tm2EC 2#汽机抽汽量
	Tm2EC float64 `json:"tm2_e_c"`
	// Tm2EO 2#汽机抽汽量优化值
	Tm2EO float64 `json:"tm2_e_o"`
	// Tm3InC 3#汽机进汽量
	Tm3InC float64 `json:"tm3_in_c"`
	// Tm3InO 3#汽机进汽量优化值
	Tm3InO float64 `json:"tm3_in_o"`
	// Tp3C 3#汽机压力值
	Tp3C float64 `json:"tp3_c"`
	// Tcos3 3#汽机功率因子
	Tcos3 float64 `json:"tcos3"`
	// Tp3O 3#汽机压力优化值
	Tp3O float64 `json:"tp3_o"`
	// Tm3EC 3#汽机抽汽量
	Tm3EC float64 `json:"tm3_e_c"`
	// Tm3EO 3#汽机抽汽量优化值
	Tm3EO float64 `json:"tm3_e_o"`
	// Tm4InC 4#汽机进汽量
	Tm4InC float64 `json:"tm4_in_c"`
	// Tm4InO 4#汽机进汽量优化值
	Tm4InO float64 `json:"tm4_in_o"`
	// Tp4C 4#汽机压力值
	Tp4C float64 `json:"tp4_c"`
	// Tcos4 4#汽机功率因子
	Tcos4 float64 `json:"tcos4"`
	// Tp4O 4#汽机压力优化值
	Tp4O float64 `json:"tp4_o"`
	// Tm4EC 4#汽机抽汽量
	Tm4EC float64 `json:"tm4_e_c"`
	// Tm4EO 4#汽机抽汽量优化值
	Tm4EO float64 `json:"tm4_e_o"`
	// C1VC 1#减温减压入口蒸汽阀门开度
	C1VC float64 `json:"c1_v_c"`
	// C1VO 1#减温减压入口蒸汽阀门开度优化值
	C1VO float64 `json:"c1_v_o"`
	// C1WC 1#减温减压水阀门开度
	C1WC float64 `json:"c1_w_c"`
	// C1WO 1#减温减压水阀门开度优化值
	C1WO float64 `json:"c1_w_o"`
	// Cm1VOutC 1#减温减压出口蒸汽流量
	Cm1VOutC float64 `json:"cm1_v_out_c"`
	// Cm1VOutO 1#减温减压出口蒸汽流量优化值
	Cm1VOutO float64 `json:"cm1_v_out_o"`
	// C2VC 2#减温减压入口蒸汽阀门开度
	C2VC float64 `json:"c2_v_c"`
	// C2VO 2#减温减压入口蒸汽阀门开度优化值
	C2VO float64 `json:"c2_v_o"`
	// C2WC 2#减温减压水阀门开度
	C2WC float64 `json:"c2_w_c"`
	// C2WO 2#减温减压水阀门开度优化值
	C2WO float64 `json:"c2_w_o"`
	// Cm2VOutC 2#减温减压出口蒸汽流量
	Cm2VOutC float64 `json:"cm2_v_out_c"`
	// Cm2VOutO 2#减温减压出口蒸汽流量优化值
	Cm2VOutO float64 `json:"cm2_v_out_o"`
	// C4VC 4#减温减压入口蒸汽阀门开度
	C4VC float64 `json:"c4_v_c"`
	// C4VO 4#减温减压入口蒸汽阀门开度优化值
	C4VO float64 `json:"c4_v_o"`
	// C4WC 4#减温减压水阀门开度
	C4WC float64 `json:"c4_w_c"`
	// C4WO 4#减温减压水阀门开度优化值
	C4WO float64 `json:"c4_w_o"`
	// Cm4VOutC 4#减温减压出口蒸汽流量
	Cm4VOutC float64 `json:"cm4_v_out_c"`
	// Cm4VOutO 4#减温减压出口蒸汽流量优化值
	Cm4VOutO float64 `json:"cm4_v_out_o"`
	// C5VC 5#减温减压入口蒸汽阀门开度
	C5VC float64 `json:"c5_v_c"`
	// C5VO 5#减温减压入口蒸汽阀门开度优化值
	C5VO float64 `json:"c5_v_o"`
	// C5WC 5#减温减压水阀门开度
	C5WC float64 `json:"c5_w_c"`
	// C5WO 5#减温减压水阀门开度优化值
	C5WO float64 `json:"c5_w_o"`
	// Cm5VOutC 5#减温减压出口蒸汽流量
	Cm5VOutC float64 `json:"cm5_v_out_c"`
	// Cm5VOutO 5#减温减压出口蒸汽流量优化值
	Cm5VOutO float64 `json:"cm5_v_out_o"`
	// H1BuyC 1#需电量卖出
	H1BuyC float64 `json:"h1_buy_c"`
	// H1BuyO 1#需电量买入
	H1BuyO float64 `json:"h1_buy_o"`
	// H1Use 1#线使用
	H1Use float64 `json:"h1_use"`
	// H2BuyC 2#需电量卖出
	H2BuyC float64 `json:"h2_buy_c"`
	// H2BuyO 2#需电量买入
	H2BuyO float64 `json:"h2_buy_o"`
	// H2Use 2#线使用
	H2Use float64 `json:"h2_use"`
	// Y1BuyC 3#需电量卖出
	Y1BuyC float64 `json:"y1_buy_c"`
	// Y1BuyO 3#需电量买入
	Y1BuyO float64 `json:"y1_buy_o"`
	// Y1Use 3#线使用
	Y1Use float64 `json:"y1_use"`
	// HightMC 高压蒸汽量实际值
	HightMC float64 `json:"hight_m_c"`
	// HightMO 高压蒸汽量优化值
	HightMO float64 `json:"hight_m_o"`
	// MiddleMC 中压蒸汽量实际值
	MiddleMC float64 `json:"middle_m_c"`
	// MiddleMO 中压蒸汽量优化值
	MiddleMO float64 `json:"middle_m_o"`
	// LowMC 低压蒸汽量实际值
	LowMC float64 `json:"low_m_c"`
	// LowMO 低压蒸汽量优化值
	LowMO float64 `json:"low_m_o"`
	// HistoryCost 历史总成本值
	HistoryCost float64 `json:"history_cost"`
	// ActualCost 实际总成本值
	ActualCost float64 `json:"actual_cost"`
	// OptCost 优化总成本
	OptCost float64 `json:"opt_cost"`
	// HistoryCostBm1 1#锅炉历史值
	HistoryCostBm1 float64 `json:"history_cost_bm1"`
	// HistoryCostBm2 2#锅炉历史值
	HistoryCostBm2 float64 `json:"history_cost_bm2"`
	// HistoryCostBm3 3#锅炉历史值
	HistoryCostBm3 float64 `json:"history_cost_bm3"`
	// HistoryCostBm4 4#锅炉历史值
	HistoryCostBm4 float64 `json:"history_cost_bm4"`
	// HistoryCostBm5 5#锅炉历史值
	HistoryCostBm5 float64 `json:"history_cost_bm5"`
	// HistoryCostBm6 6#锅炉历史值
	HistoryCostBm6 float64 `json:"history_cost_bm6"`
	// HistoryCostBm7 7#锅炉历史值
	HistoryCostBm7 float64 `json:"history_cost_bm7"`
	// ActualCostBm1 1#锅炉实际值
	ActualCostBm1 float64 `json:"actual_cost_bm1"`
	// ActualCostBm2 2#锅炉实际值
	ActualCostBm2 float64 `json:"actual_cost_bm2"`
	// ActualCostBm3 3#锅炉实际值
	ActualCostBm3 float64 `json:"actual_cost_bm3"`
	// ActualCostBm4 4#锅炉实际值
	ActualCostBm4 float64 `json:"actual_cost_bm4"`
	// ActualCostBm5 5#锅炉实际值
	ActualCostBm5 float64 `json:"actual_cost_bm5"`
	// ActualCostBm6 6#锅炉实际值
	ActualCostBm6 float64 `json:"actual_cost_bm6"`
	// ActualCostBm7 7#锅炉实际值
	ActualCostBm7 float64 `json:"actual_cost_bm7"`
	// HistoryCostTm1 1#历史同期值
	HistoryCostTm1 float64 `json:"history_cost_tm1"`
	// HistoryCostTm2 2#历史同期值
	HistoryCostTm2 float64 `json:"history_cost_tm2"`
	// HistoryCostTm3 3#历史同期值
	HistoryCostTm3 float64 `json:"history_cost_tm3"`
	// HistoryCostTm4 4#历史同期值
	HistoryCostTm4 float64 `json:"history_cost_tm4"`
	// ActualCostTm1 1#汽机实际值
	ActualCostTm1 float64 `json:"actual_cost_tm1"`
	// ActualCostTm2 2#汽机实际值
	ActualCostTm2 float64 `json:"actual_cost_tm2"`
	// ActualCostTm3 3#汽机实际值
	ActualCostTm3 float64 `json:"actual_cost_tm3"`
	// ActualCostTm4 4#汽机实际值
	ActualCostTm4 float64 `json:"actual_cost_tm4"`
	// HistoryCostH1 1#线优化历史值
	HistoryCostH1 float64 `json:"history_cost_h1"`
	// HistoryCostH2 2#线优化历史值
	HistoryCostH2 float64 `json:"history_cost_h2"`
	// HistoryCostY1 3#线优化历史值
	HistoryCostY1 float64 `json:"history_cost_y1"`
	// ActualCostH1 1#线优化实际值
	ActualCostH1 float64 `json:"actual_cost_h1"`
	// ActualCostH2 2#线优化实际值
	ActualCostH2 float64 `json:"actual_cost_h2"`
	// ActualCostY1 3#线优化实际值
	ActualCostY1 float64 `json:"actual_cost_y1"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
	// InsertTimestamp 插入数据的时间
	InsertTimestamp int64 `json:"insert_timestamp"`
	// OptCostBm1 1#锅炉优化目标值
	OptCostBm1 float64 `json:"opt_cost_bm1"`
	// OptCostBm2 2#锅炉优化目标值
	OptCostBm2 float64 `json:"opt_cost_bm2"`
	// OptCostBm3 3#锅炉优化目标值
	OptCostBm3 float64 `json:"opt_cost_bm3"`
	// OptCostBm4 4#锅炉优化目标值
	OptCostBm4 float64 `json:"opt_cost_bm4"`
	// OptCostBm5 5#锅炉优化目标值
	OptCostBm5 float64 `json:"opt_cost_bm5"`
	// OptCostBm6 6#锅炉优化目标值
	OptCostBm6 float64 `json:"opt_cost_bm6"`
	// OptCostBm7 7#锅炉优化目标值
	OptCostBm7 float64 `json:"opt_cost_bm7"`
	// OptCostTm1 1#汽机优化目标值
	OptCostTm1 float64 `json:"opt_cost_tm1"`
	// OptCostTm2 2#汽机优化目标值
	OptCostTm2 float64 `json:"opt_cost_tm2"`
	// OptCostTm3 3#汽机优化目标值
	OptCostTm3 float64 `json:"opt_cost_tm3"`
	// OptCostTm4 4#汽机优化目标值
	OptCostTm4 float64 `json:"opt_cost_tm4"`
	// OptCostH1 1#线优化目标值
	OptCostH1 float64 `json:"opt_cost_h1"`
	// OptCostH2 2#线优化目标值
	OptCostH2 float64 `json:"opt_cost_h2"`
	// OptCostY1 袁博线优化目标值
	OptCostY1 float64 `json:"opt_cost_y1"`
}



type Constrain struct {
	Name []string
	Value []float64
}

type stringSlice []string

func (slice stringSlice) pos(value string) int {
	for p,v := range slice {
		if (v == value) {
			return p
		}
	}
	return -1
}

//type FloatSlice []float64
func sum(floatSlice1,floatSlice2 []float64) float64{
	if len(floatSlice1) == 0 || len(floatSlice1) != len(floatSlice2) {
		return -1
	}
	sum_ret := 0.0
	for p,_ := range floatSlice1 {
		sum_ret += floatSlice1[p] * floatSlice2[p]
	}
	return sum_ret
}

func Test_ilp() {
	prob := MilpProblem{
		C: []float64{-1, -2, 0, 0},
		//A: mat.NewDense(2, 4, []float64{
		//	-1, 2, 1, 0,
		//	3, 1, 0, 1,
		//}),
		//b: []float64{4, 9},
		G: mat.NewDense(2, 4, []float64{
			-0.5, 2, 1.5, 0,
			3.1, 1.3, 0, 1.1,
		}),
		H: []float64{15, 18},
		IntegralityConstraints: []bool{false, true, false, false},
	}

	// Solve the problem with 1 worker and a one-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	got, _ := prob.Solve(ctx, 1, DummyMiddleware{})
	fmt.Println(got.Z)
	fmt.Println(got.X)
	// dump the logged tree to a DOT-file
}

func Test_nilp() {
	db, err := sql.Open("mysql", "root:123456@/hoko_python")
	if err != nil {
		panic(err)
	}
	defer db.Close()    //关闭数据库
	err = db.Ping()      //连接数据库
	if err!=nil{
		fmt.Println("数据库连接失败")
		return
	}
	var tm1_in_c float64
	//var tb_table_monitoring TbRealMonitoring
	row := db.QueryRow("select tm1_in_c from tb_real_monitoring order by id desc limit 1")
	fmt.Println(row)
	row.Scan(&tm1_in_c)
	fmt.Println(tm1_in_c)
	turbine_data := []float64{143.14,171.04,175.36,330.2,63.3,123.74,79.5,151.05,27.92,24.68,28.48,49.09,0.0,0.95,0.0,2.17,13.3,26.97,22.51,90.87}
	if turbine_data[12] < 0{
	turbine_data[12] = math.Abs(turbine_data[12])
	turbine_data[13] = 0
	}else if turbine_data[12] > 0{
	turbine_data[13] = turbine_data[12]
	turbine_data[12] = 0
	}
	if turbine_data[14] < 0{
	turbine_data[14] = math.Abs(turbine_data[14])
	turbine_data[15] = 0
	}else if turbine_data[14] > 0{
	turbine_data[15] = turbine_data[14]
	turbine_data[14] = 0
	turbine_data[16] = math.Abs(turbine_data[16])
	}
	//fmt.Println("turbine_data = ",turbine_data)

	var my_ctype []bool
	my_ctype_num := []int{18,2}
	//for range my_ctype_num {
		c_num,i_num := my_ctype_num[0],my_ctype_num[1]
		for j := 0;j<c_num;j++{
		my_ctype = append(my_ctype, false)
		}
		for j := 0;j<i_num;j++{
			my_ctype = append(my_ctype, true)
		}
	//}
	my_obj := []float64{140.56,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.8888*1000,-0.3941*1000,0.8888*1000,-0.3941*1000,0.8663*1000,0.0,0.0} //目标函数系数  min(c1*q1+c2*e1+c2*e3+c3*e5-c4*e2-c4*e4)
	my_rhs := []float64{26.97,22.51,90.87,0.0,151.05,261.33,8.3,7.45,24.16,-0.39,0.0,5.0,0.0,5.0,0.0,0.0,0.0,0.0} //约束条件的值相当于b
	my_colnames := stringSlice{"q1","x1","x2","x3","x4","x5","x6","x7","x8","w1","w2","w3","w4","e1","e2","e3","e4","e5","y1","y2"} //column names列向量的名字
	my_ub :=[]float64{850.0,250.0,250.0,350.0,350.0,95.0,200.0,200.0,200.0,29.0,29.0,55.0,55.0,5.0,5.0,5.0,5.0,40.0,1.0,1.0} //变量的约束条件上限
	my_lb :=[]float64{500.0,0.0  ,0.0  ,0.0  ,0.0  ,0.0  ,0.0  ,0.0  ,0.0  ,10.0,10.0,10.0,10.0,0.0,0.0,0.0,0.0,0.0 ,0.0,0.0} //变量的约束条件下限

	//fmt.Println(my_ctype)
	//fmt.Println(my_rhs)
	//fmt.Println(my_colnames)

	x1 := turbine_data[0]  // 1#汽轮机进汽量当前值
	x2 := turbine_data[1]  // 2#汽轮机进汽量当前值
	x3 := turbine_data[2]  // 3#汽轮机进汽量当前值
	x4 := turbine_data[3]  // 4#汽轮机进汽量当前值
	q1 := x1 + x2 + x3 +x4
	x5 := turbine_data[4]    // 1#汽轮机抽汽量当前值
	x6 := turbine_data[5]    // 2#汽轮机抽汽量当前值
	x7 := turbine_data[6]    // 3#汽轮机抽汽量当前值
	x8 := turbine_data[7]    // 4#汽轮机抽汽量当前值
	w1 := turbine_data[8]    // 1#汽轮机发电量当前值
	w2 := turbine_data[9]    // 2#汽轮机发电量当前值
	w3 := turbine_data[10]   // 3#汽轮机发电量当前值
	w4 := turbine_data[11]   // 4#汽轮机发电量当前值
	e1 := turbine_data[12]  // 环能一线买电量当前值
	e2 := turbine_data[13]  // 环能一线卖电量当前值
	e3 := turbine_data[14]  // 环能二线买电量当前值
	e4 := turbine_data[15]  // 环能二线卖电量当前值
	e5 := turbine_data[16]  // 袁博线买电量当前值
	m1 := turbine_data[17]  // 环能一线需电量
	m2 := turbine_data[18]  // 环能一线需电量
	//m3 := turbine_data[19]  // 袁博线需电量
	y1,_ := strconv.ParseFloat(strconv.FormatBool((m1 - w1)>0),64)
	y2,_ := strconv.ParseFloat(strconv.FormatBool((m2 - w2)>0),64)
	ori_x := []float64{q1,x1,x2,x3,x4,x5,x6,x7,x8,w1,w2,w3,w4,e1,e2,e3,e4,e5,y1,y2}
	ori_cost := sum(ori_x , my_obj)
	ori_steam_cost := ori_x[0]*my_obj[0]
	ori_power_cost := sum(ori_x[1:],my_obj[1:])


	rows_eq := []Constrain{{[]string{"w1","e1","e2"},[]float64{1.0,1.0,-1.0}},
	{[]string{"w2","e3","e4"},[]float64{1.0,1.0,-1.0}},
	{[]string{"w3","w4","e5"},[]float64{1.0,1.0,1.0}},
	{[]string{"x1","x2","x3","x4","q1"},[]float64{1.0,1.0,1.0,1.0,-1.0}},
	{[]string{"x8"},[]float64{1.0}},
	{[]string{"x5","x6","x7"},[]float64{1.0,1.0,1.0}},
	{[]string{"x1","x5","w1"},[]float64{1.0,-0.41,-3.9}},
	{[]string{"x2","x6","w2"},[]float64{1.0,-0.61,-3.57}},
	{[]string{"x3","x7","w3"},[]float64{1.0,-0.82,-3.02}},
	{[]string{"x4","x8","w4"},[]float64{1.0,-0.89,-3.98}}}

	rows_ineq := []Constrain{
		{[]string{"e1","y1"},[]float64{1.0,-5.0}},
		{[]string{"e2","y1"},[]float64{1.0,5.0}},
		{[]string{"e3","y2"},[]float64{1.0,-5.0}},
		{[]string{"e4","y2"},[]float64{1.0,5.0}},
		{[]string{"x5","x1"},[]float64{1.0,-0.85}},
		{[]string{"x6","x2"},[]float64{1.0,-0.85}},
		{[]string{"x7","x3"},[]float64{1.0,-0.85}},
		{[]string{"x8","x4"},[]float64{1.0,-0.85}}}

	cons_val_eq_one := make([]float64,len(my_colnames))
	cons_eq_val := [][]float64{}
	cons_eq_ret := []float64{}
	cons_val_ineq_one := make([]float64,len(my_colnames))
	cons_ineq_val := [][]float64{}
	cons_ineq_ret := []float64{}
	//fmt.Println(cons_val_eq_one)
	//fmt.Println(rows_eq)
	for i := 0;i<len(rows_eq);i++{
		//fmt.Println(rows_eq[i].Name)
		for j := 0;j<len(rows_eq[i].Name);j++ {
			//my_colnames(my_colnames,j)
			idx := my_colnames.pos(rows_eq[i].Name[j])
			//fmt.Println("idx = ",idx)
			if idx != -1 {
				cons_val_eq_one[idx] = rows_eq[i].Value[j]
			}
		}
		cons_eq_val = append(cons_eq_val,cons_val_eq_one)
		cons_val_eq_one = make([]float64,len(my_colnames))
	}
	for i := 0;i<len(rows_ineq);i++{
		//fmt.Println(rows_ineq[i].Name)
		for j := 0;j<len(rows_ineq[i].Name);j++ {
			//my_colnames(my_colnames,j)
			idx := my_colnames.pos(rows_ineq[i].Name[j])
			//fmt.Println("idx = ",idx)
			if idx != -1 {
				cons_val_ineq_one[idx] = rows_ineq[i].Value[j]
			}
		}
		cons_ineq_val = append(cons_ineq_val,cons_val_ineq_one)
		cons_val_ineq_one = make([]float64,len(my_colnames))
	}
	//fmt.Println(cons_eq_val)
	//fmt.Println(cons_ineq_val)
	for i := 0;i<len(cons_eq_val);i++ {
		for j := 0;j<len(cons_eq_val[0]);j++{
			cons_eq_ret = append(cons_eq_ret, cons_eq_val[i][j])
		}
	}
	for i := 0;i<len(cons_ineq_val);i++ {
		for j := 0;j<len(cons_ineq_val[0]);j++{
			cons_ineq_ret = append(cons_ineq_ret, cons_ineq_val[i][j])
		}
	}
	//fmt.Println(cons_eq_ret)
	//fmt.Println(cons_ineq_ret)

	my_rhs_eq := my_rhs[:len(rows_eq)]
	my_rhs_ineq := my_rhs[len(rows_eq):]
	//fmt.Println(my_rhs_eq)
	//fmt.Println(my_rhs_ineq)
	if len(my_ub) != 0 && len(my_ub) == len(my_lb) {
		for p, _ := range my_ub {
			cons_cp := make([]float64, len(my_colnames))
			cons_val_ineq_one[p] = 1
			my_rhs_ineq = append(my_rhs_ineq, my_ub[p])
			cons_cp[p] = -1
			my_rhs_ineq = append(my_rhs_ineq, -my_lb[p])
			//fmt.Println(cons_val_ineq_one)
			//fmt.Println(cons_cp)
			//fmt.Println(my_rhs_ineq)
			//fmt.Println(len(my_rhs_ineq))
			for _,v1 := range cons_val_ineq_one {
				cons_ineq_ret = append(cons_ineq_ret, v1)
			}
			for _,v2 := range cons_cp {
				cons_ineq_ret = append(cons_ineq_ret, v2)
			}
			cons_val_ineq_one = make([]float64, len(my_colnames))
		}
	}

	prob := MilpProblem{
		C: my_obj, //目标函数系数  min(c1*q1+c2*e1+c2*e3+c3*e5-c4*e2-c4*e4)
		A: mat.NewDense(len(rows_eq), len(my_colnames), cons_eq_ret),
		B: my_rhs_eq,
		G: mat.NewDense(len(rows_ineq)+len(my_ub)*2, len(my_colnames), cons_ineq_ret),
		H: my_rhs_ineq,
		//h: []float64{25.5, 20.6,20,30},
		IntegralityConstraints: my_ctype,
	}

	// Solve the problem with 1 worker and a one-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	got, _ := prob.Solve(ctx, 1, DummyMiddleware{})
	//fmt.Println(prob.G)
	//fmt.Println(got.z)
	//fmt.Println(got.x)
	opt_x := got.X
	opt_cost := got.Z
	opt_steam_cost := my_obj[0]*opt_x[0]
	opt_power_cost := sum(my_obj[1:],opt_x[1:])
	opt_cost_delta := ori_cost - opt_cost
	opt_cost_degree := opt_cost_delta/ori_cost
	fmt.Println("ori_cost = ", ori_cost)
	fmt.Println("opt_cost = ", )
	fmt.Println("ori_x = ", ori_x)
	fmt.Println("opt_x = ", opt_x)
	fmt.Println("ori_steam_cost = ", ori_steam_cost)
	fmt.Println("opt_steam_cost = ", opt_steam_cost)
	fmt.Println("ori_power_cost = ", ori_power_cost)
	fmt.Println("opt_power_cost = ", opt_power_cost)
	fmt.Println("opt_cost_delta = ", opt_cost_delta)
	fmt.Println("opt_cost_degree = ", opt_cost_degree)


	//prob := MilpProblem{
	//	c: []float64{-1.2,2.3,5.2,-1.5}, //目标函数系数  min(c1*q1+c2*e1+c2*e3+c3*e5-c4*e2-c4*e4)
	//	//A: mat.NewDense(2, 4, []float64{
	//	//	-1, 2, 1, 0,
	//	//	3, 1, 0, 1,
	//	//}),
	//	//b: []float64{4, 9.2},
	//	G: mat.NewDense(12, 4, []float64{
	//		-1,0,0,0,
	//		1,0,0,0,
	//		0,-1,0,0,
	//		0,1,0,0,
	//		0,0,-1,0,
	//		0,0,1,0,
	//		0,0,0,-1,
	//		0,0,0,1,
	//		-0.5, 0, 1.5, 0,
	//		3.1, 5, 0, 1.1,
	//		1.1, 0, 0.5, 2.1,
	//		-2.1, -1.2, 0, -1.1,
	//	}),
	//	h: []float64{12,12.5,10,10.8,10,10.5,2,2.8,25.5, 20.6,20,30},
	//	//h: []float64{25.5, 20.6,20,30},
	//	IntegralityConstraints: []bool{false, false, false, false},
	//}
	//
	//// Solve the problem with 1 worker and a one-second timeout
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//
	//got, _ := prob.Solve(ctx, 1, DummyMiddleware{})
	//fmt.Println(prob.G)
	//fmt.Println(got.z)
	//fmt.Println(got.x)
	//// dump the logged tree to a DOT-file
}

type NewMilp struct {
	my_obj []float64
	my_colnames []string
	rows_eq [][]float64
	cons_eq_ret []float64
	my_rhs_eq []float64
	rows_ineq [][]float64
	my_rhs_ineq []float64
	cons_ineq_ret []float64
	my_ub []float64
	my_lb []float64
	my_ctype []bool
}

func Test()  {
	var newMilps []NewMilp
	var newMilp NewMilp
	newMilp.rows_eq = [][]float64{
		{3, 1, 0, 1},
	}
	newMilp.cons_eq_ret = []float64{3, 1, 0, 1}
	newMilp.my_rhs_eq = []float64{9.2}
	newMilp.rows_ineq = [][]float64{
		{-0.5, 0, 1.5, 0},
		{3.1, 5, 0, 1.1},
		{1.1, 0, 0.5, 2.1},
		{-2.1, -1.2, 0, -1.1},
	}
	newMilp.cons_ineq_ret = []float64{
		-0.5, 0, 1.5, 0,
		3.1, 5, 0, 1.1,
		1.1, 0, 0.5, 2.1,
		-2.1, -1.2, 0, -1.1,
	}
	newMilp.my_rhs_ineq = []float64{2.8,25.5, 20.6,20,30}
	newMilps = append(newMilps, newMilp)
	TestMilpAuto(newMilps)
}

func TestMilpAuto(newMilps []NewMilp) {
	my_obj := []float64{-1.2,2.3,5.2,-1.5}
	my_colnames := []string{"x1","x2","x3","x4"}
	rows_eq := [][]float64{
		{-1, 2, 1, 0},
	}

	cons_eq_ret := []float64{
		-1, 2, 1, 0,
	}
	my_rhs_eq := []float64{4}
	rows_ineq := [][]float64{
		{-1,0,0,0},
		{1,0,0,0},
		{0,-1,0,0},
		{0,1,0,0},
		{0,0,-1,0},
		{0,0,1,0},
		{0,0,0,-1},
		{0,0,0,1},

	}
	my_rhs_ineq := []float64{12,12.5,10,10.8,10,10.5,2}
	cons_ineq_ret :=[]float64{
		-1,0,0,0,
		1,0,0,0,
		0,-1,0,0,
		0,1,0,0,
		0,0,-1,0,
		0,0,1,0,
		0,0,0,-1,
		0,0,0,1,

	}
	//my_ub := []float64{}
	//my_lb := []float64{}
	my_ctype := []bool{false,false,false,false}

	//配置
	for _,oneMilp := range newMilps {
		for _,item := range oneMilp.rows_eq {
			rows_eq = append(rows_eq, item)
		}
		for _,item := range oneMilp.cons_eq_ret {
			cons_eq_ret = append(cons_eq_ret, item)
		}
		for _,item := range oneMilp.my_rhs_eq {
			my_rhs_eq = append(my_rhs_eq, item)
		}
		for _,item := range oneMilp.rows_ineq {
			rows_ineq = append(rows_ineq, item)
		}
		for _,item := range oneMilp.my_rhs_ineq {
			my_rhs_ineq = append(my_rhs_ineq, item)
		}
		for _,item := range oneMilp.cons_ineq_ret {
			cons_ineq_ret = append(cons_ineq_ret, item)
		}
	}

	prob := MilpProblem{
		C: my_obj, //目标函数系数  min(c1*q1+c2*e1+c2*e3+c3*e5-c4*e2-c4*e4)
		A: mat.NewDense(len(rows_eq), len(my_colnames), cons_eq_ret),
		B: my_rhs_eq,
		//G: mat.NewDense(len(rows_ineq)+len(my_ub)*2, len(my_colnames), cons_ineq_ret),
		G: mat.NewDense(len(rows_ineq), len(my_colnames), cons_ineq_ret),
		H: my_rhs_ineq,
		//h: []float64{25.5, 20.6,20,30},
		IntegralityConstraints: my_ctype,
	}

	// Solve the problem with 1 worker and a one-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	got, _ := prob.Solve(ctx, 1, DummyMiddleware{})
	fmt.Println(prob.G)
	fmt.Println(got.Z)
	fmt.Println(got.X)
	// dump the logged tree to a DOT-file
}

func TestMilpAuto2() {
	//prob := MilpProblem{
	//	c: []float64{-1.2,2.3,5.2,-1.5}, //目标函数系数  min(c1*q1+c2*e1+c2*e3+c3*e5-c4*e2-c4*e4)
	//	A: mat.NewDense(2, 4, []float64{
	//		-1, 2, 1, 0,
	//		3, 1, 0, 1,
	//	}),
	//	b: []float64{4, 9.2},
	//	G: mat.NewDense(12, 4, []float64{
	//		-1,0,0,0,
	//		1,0,0,0,
	//		0,-1,0,0,
	//		0,1,0,0,
	//		0,0,-1,0,
	//		0,0,1,0,
	//		0,0,0,-1,
	//		0,0,0,1,
	//		-0.5, 0, 1.5, 0,
	//		3.1, 5, 0, 1.1,
	//		1.1, 0, 0.5, 2.1,
	//		-2.1, -1.2, 0, -1.1,
	//	}),
	//	h: []float64{12,12.5,10,10.8,10,10.5,2,2.8,25.5, 20.6,20,30},
	//	//h: []float64{25.5, 20.6,20,30},
	//	IntegralityConstraints: []bool{false, false, false, false},
	//}
	//my_ctype_num := []int{18,2}
	my_obj := []float64{-1.2,2.3,5.2,-1.5}
	my_colnames := []string{"x1","x2","x3","x4"}
	rows_eq := [][]float64{
		{-1, 2, 1, 0},
		{3, 1, 0, 1},
	}

	cons_eq_ret := []float64{
		-1, 2, 1, 0,
		3, 1, 0, 1,
	}
	my_rhs_eq := []float64{4, 9.2}
	rows_ineq := [][]float64{
		{-1,0,0,0},
		{1,0,0,0},
		{0,-1,0,0},
		{0,1,0,0},
		{0,0,-1,0},
		{0,0,1,0},
		{0,0,0,-1},
		{0,0,0,1},
		{-0.5, 0, 1.5, 0},
		{3.1, 5, 0, 1.1},
		{1.1, 0, 0.5, 2.1},
		{-2.1, -1.2, 0, -1.1},
	}
	my_rhs_ineq := []float64{12,12.5,10,10.8,10,10.5,2,2.8,25.5, 20.6,20,30}
	cons_ineq_ret :=[]float64{
		-1,0,0,0,
		1,0,0,0,
		0,-1,0,0,
		0,1,0,0,
		0,0,-1,0,
		0,0,1,0,
		0,0,0,-1,
		0,0,0,1,
		-0.5, 0, 1.5, 0,
		3.1, 5, 0, 1.1,
		1.1, 0, 0.5, 2.1,
		-2.1, -1.2, 0, -1.1,
	}
	//my_ub := []float64{}
	//my_lb := []float64{}
	my_ctype := []bool{false,false,false,false}
	prob := MilpProblem{
		C: my_obj, //目标函数系数  min(c1*q1+c2*e1+c2*e3+c3*e5-c4*e2-c4*e4)
		A: mat.NewDense(len(rows_eq), len(my_colnames), cons_eq_ret),
		B: my_rhs_eq,
		//G: mat.NewDense(len(rows_ineq)+len(my_ub)*2, len(my_colnames), cons_ineq_ret),
		G: mat.NewDense(len(rows_ineq), len(my_colnames), cons_ineq_ret),
		H: my_rhs_ineq,
		//h: []float64{25.5, 20.6,20,30},
		IntegralityConstraints: my_ctype,
	}

	// Solve the problem with 1 worker and a one-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	got, _ := prob.Solve(ctx, 1, DummyMiddleware{})
	fmt.Println(prob.G)
	fmt.Println(got.Z)
	fmt.Println(got.X)
	// dump the logged tree to a DOT-file
}
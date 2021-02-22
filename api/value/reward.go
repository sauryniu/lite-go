package value

// Reward is 奖励
type Reward struct {
	Count       int64
	Route       string
	TargetType  TargetTypeValue
	TargetIndex int
	ValueType   TypeValue
}

// TargetTypeValue is 目标类型值
type TargetTypeValue int

// TypeValue is 数值类型值
type TypeValue int

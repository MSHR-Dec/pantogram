package model

import (
	"math"
)

type Prefecture struct {
	Tag   PrefectureTag
	Field PrefectureField
}

type PrefectureTag struct {
	Name string
}

type PrefectureField struct {
	Total int
	Count int
	Ratio float64
}

func (f *PrefectureField) CalcRatio() {
	ratio := float64(f.Count) / float64(f.Total) * 100
	f.Ratio = math.Round(ratio*100) / 100
}

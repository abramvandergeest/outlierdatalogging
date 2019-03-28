package outlierdatalogging

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	DriverName string `md:"driverName,required"`
	PsqlInfo   string `md:"psqlInfo,required"`
}

type Input struct {
	Ind  int64   `md:"ind,required"`
	Act  int64   `md:"act,required"`
	Pred []int64 `md:"pred,required"`
	T    int64   `md:"t,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	indVal, _ := coerce.ToInt64(values["ind"])
	r.Ind = indVal
	actVal, _ := coerce.ToInt64(values["act"])
	r.Act = actVal
	r.Pred = values["pred"].([]int64)
	tVal, _ := coerce.ToInt64(values["t"])
	r.T = tVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ind":  r.Ind,
		"act":  r.Act,
		"pred": r.Pred,
		"t":    r.T,
	}
}

type Output struct {
	Output string `md:"output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["output"])
	o.Output = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Output,
	}
}

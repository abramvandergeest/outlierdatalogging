package outlierdatalogging

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	DriverName string `md:"driverName,required"`
	PsqlInfo   string `md:"psqlInfo,required"`
}

type Input struct {
	Ind  int32         `md:"ind,required"`
	Act  int32         `md:"act,required"`
	Pred []interface{} `md:"pred,required"`
	T    int32         `md:"t,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	indVal, _ := coerce.ToInt32(values["ind"])
	r.Ind = indVal
	actVal, _ := coerce.ToInt32(values["act"])
	r.Act = actVal
	r.Pred = values["pred"].([]interface{})
	tVal, _ := coerce.ToInt32(values["t"])
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

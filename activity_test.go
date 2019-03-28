package outlierdatalogging

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {
	settings := &Settings{DriverName: "postgres", PsqlInfo: "host=localhost port=5432 user=flogo password=flynnRocks dbname=avanderg@tibco.com sslmode=disable"}
	iCtx := test.NewActivityInitContext(settings, nil)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())
	var p []interface{}
	p = append(p, 0)
	input := &Input{Ind: 234, Act: 1, Pred: p, T: 5}
	err = tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{}
	err = tc.GetOutputObject(output)
	assert.Nil(t, err)
	assert.Equal(t, "data has been inserted into database", output.Output)
}

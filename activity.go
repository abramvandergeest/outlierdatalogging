package outlierdatalogging

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	activity.Register(&Activity{}, New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {
	fmt.Println("Gets to the NEW function")
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s, %s", s.DriverName, s.PsqlInfo)
	fmt.Println(s)
	act := &Activity{settings: s} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	settings *Settings
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	drivername := a.settings.DriverName
	psqlInfo := a.settings.PsqlInfo
	ind := input.Ind
	actual := input.Act
	pred := input.Pred[0]
	t := input.T

	ctx.Logger().Debugf("Input: %s, %s, %s,%s", ind, actual, pred, t)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return true, fmt.Errorf("opening %s database failed: %s", drivername, err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return true, fmt.Errorf("opening %s database failed: %s", drivername, err)
	}

	tp := 0
	fp := 0
	fn := 0
	tn := 0
	if actual == pred && actual == 1 {
		tp = 1
	} else if actual == pred && actual == 0 {
		tn = 1
	} else if actual != pred && pred == 1 {
		fp = 1
	} else if actual != pred && pred == 0 {
		fn = 1
	}

	//  CREATE TABLE IF NOT EXISTS test.outlier (ind integer, act integer, pred integer, t integer);
	sqlStatement := `
	INSERT INTO test.outlier (ind, act, pred, t,truepos,falsepos,falseneg,trueneg) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err = db.Exec(sqlStatement, ind, actual, pred, t, tp, fp, fn, tn)
	if err != nil {
		return true, fmt.Errorf("inserting into %s database failed: %s", drivername, err)
	}

	output := &Output{Output: "data has been inserted into database"}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

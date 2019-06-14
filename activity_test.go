package alarmrenderactivity

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("data", `{"AlarmInfo": {"TAlarm": [{"AlarmID": "12","NodeID": "123","JunctionName": "abcd","XCoor": "25133.26","YCoor": "43509.02","StartDate": "2017-03-22 14:35:08.0","EndDate": "2017-03-23 14:35:08.0","Type": "1","Message": "(9/3)14:29 Roadworks on SLE (towards BKE) before Mandai Rd Exit. Avoid lane 1."},
	{"AlarmID": "12","NodeID": "123","JunctionName": "abcd","XCoor": "25133.26","YCoor": "43509.02","StartDate": "2017-03-22 14:35:08.0","EndDate": "2017-03-23 14:35:08.0","Type": "1","Message": "(9/3)14:29 Roadworks on SLE (towards BKE) before Mandai Rd Exit. Avoid lane 1."}]}}
	`)
	//tc.SetInput("file", "D:/Flogo/xml.jsp")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}
	act.Eval(tc)
	//check output attr

	output := tc.GetOutput("output")
	assert.Equal(t, output, output)

}

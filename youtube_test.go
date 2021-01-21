package youtube

import (
	"testing"
)

func TestNoErrorWithCorrectAPIKey(t *testing.T) {

	_, err := InitYouTubeAPI("AIzaSyDLgJuC5rUCpzlqVA2QBRIFa9WpDaSBITg")

	if err != nil {
		t.Errorf("Error returned.")
	}
}

func TestReturnChannelJSONWithId(t *testing.T) {

	service, err := InitYouTubeAPI("AIzaSyDLgJuC5rUCpzlqVA2QBRIFa9WpDaSBITg")

	channel, err := ChannelById(service, []string{"id", "statistics"}, "UCXTpFs_3PqI41qX2d9tL2Rw")

	if err != nil {
		t.Errorf("Error occured: %s", err.Error())
	}

	if channel.Id != "UCXTpFs_3PqI41qX2d9tL2Rw" {
		t.Errorf("Did not return the correct channel data.")
	}

	bytes, err := channel.MarshalJSON()

	t.Log(string(bytes))

}

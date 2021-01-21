package main

import "testing"

func TestLineParse(t *testing.T) {

	testLine := "<td title=\"4tgm6uUNPVw.mp4\"><a href=\"/dl/4tgm6uUNPVw.mp4\" target=\"_blank\">【鬼畜ゲー】一人でおかしいうさぎのゲームする【Super Bunny Man】</a></td>"

	name, id := ParseLine(testLine)

	if name == "" || id == "" {
		t.Errorf("Returned empty with correct input")
	}

	if name != "【鬼畜ゲー】一人でおかしいうさぎのゲームする【Super Bunny Man】" && id != "4tgm6uUNPVw.mp4" {
		t.Errorf("Did not return correct name and id")
	}

	testLine = "<a class=\"navbar-brand\" href=\"\">hololivevideos.download</a>"

	name, id = ParseLine(testLine)

	if name != "" && id != "" {
		t.Errorf("Did not return empty string with incorrect input")
	}
}

func TestParseFile(t *testing.T) {

}

package soundcloudapi_test

import (
	"strings"
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestGetDownloadURL(t *testing.T) {
	dlURL, err := api.GetDownloadURL("https://soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the", "hls")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !strings.Contains(dlURL, "sndcdn.com/") {
		t.Errorf("Invalid download URL returned, received: (%s)", dlURL)
	}
}

func TestGetDownloadURLPublic(t *testing.T) {
	// This track has a public download URL link
	trackInfo, err := api.GetTrackInfo(soundcloudapi.GetTrackInfoOptions{
		URL: "https://soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !trackInfo[0].Downloadable {
		t.Error("Track changed, update the URL")
		return
	}

	dlURL, err := api.GetDownloadURL("https://soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the", "")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !strings.Contains(dlURL, "sndcdn.com/") {
		t.Errorf("Invalid download URL returned, received: (%s)", dlURL)
	}
}

func TestGetStreamURL(t *testing.T) {
	dlURL, err := api.GetStreamURL("https://soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the", "progressive")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !strings.Contains(dlURL, "sndcdn.com/") {
		t.Errorf("Invalid download URL returned, received: (%s)", dlURL)
	}
}

func TestGetDownloadURLNewLink(t *testing.T) {
	// This track has a public download URL link
	trackInfo, err := api.GetTrackInfo(soundcloudapi.GetTrackInfoOptions{
		URL: "https://on.soundcloud.com/t1Jie",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if trackInfo[0].Downloadable {
		t.Error("Track changed, update the URL")
		return
	}
	if trackInfo[0].Title != "Ocean Eyes" {
		t.Error("Invalid track title, did it change?")
		return
	}

	dlURL, err := api.GetDownloadURL("https://on.soundcloud.com/t1Jie", "")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !strings.Contains(dlURL, "sndcdn.com/") {
		t.Errorf("Invalid download URL returned, received: (%s)", dlURL)
	}
}

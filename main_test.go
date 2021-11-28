package main

import (
	"testing"
)

func TestGetCorrespondingMediaName(t *testing.T) {
	// IMG_7286.jpg.json      => IMG_7286.jpg
	// IMG_7286.jpg(1).json   => IMG_7286(1).jpg
	// IMG_7286.jpg(2).json   => IMG_7286(2).jpg

	if getCorrespondingMediaName("IMG_7286.jpg.json") != "IMG_7286.jpg" {
		t.Error(`"IMG_7286.jpg.json" => ! "IMG_7286.jpg"`)
	}
	if getCorrespondingMediaName("IMG_7286.jpg(1).json") != "IMG_7286(1).jpg" {
		t.Error(`"IMG_7286.jpg(1).json" => ! "IMG_7286(1).jpg"`)
	}
	if getCorrespondingMediaName("IMG_7286.jpg(2).json") != "IMG_7286(2).jpg" {
		t.Error(`"IMG_7286.jpg(2).json" => ! "IMG_7286(2).jpg"`)
	}
}

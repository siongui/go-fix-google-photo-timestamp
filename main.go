package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(\(\d+\)$)`)

func getCorrespondingMediaName(jsonFilename string) string {
	// IMG_7286.jpg.json      => IMG_7286.jpg
	// IMG_7286.jpg(1).json   => IMG_7286(1).jpg
	// IMG_7286.jpg(2).json   => IMG_7286(2).jpg

	// jsonNoExt = IMG_7286.jpg(1)
	jsonNoExt := strings.TrimSuffix(jsonFilename, filepath.Ext(jsonFilename))
	if re.MatchString(jsonNoExt) {
		// num = (1)
		num := re.FindStringSubmatch(jsonNoExt)[0]
		// nameNoNum = IMG_7286.jpg
		nameNoNum := strings.TrimSuffix(jsonNoExt, num)
		// ext = .jpg
		nameNoNumExt := filepath.Ext(nameNoNum)
		// nameNoNumNoExt = IMG_7286
		nameNoNumNoExt := strings.TrimSuffix(nameNoNum, nameNoNumExt)

		return nameNoNumNoExt + num + nameNoNumExt
	}
	return jsonNoExt
}

func setMetadata(dir string, file os.DirEntry) (err error) {
	m, err := decodeJsonMetadata(dir, file)
	if err != nil {
		return
	}

	if m.PhotoTakenTime.Timestamp == "" {
		return
	}

	mediaName := getCorrespondingMediaName(file.Name())
	log.Printf("Media: %s\tMetadata: %s\ttitle(original media name): %s\ttimestamp: %s\n", mediaName, file.Name(), m.Title, m.PhotoTakenTime.Timestamp)
	takenTime, err := unixTimestampToGoTime(m.PhotoTakenTime.Timestamp)
	if err != nil {
		log.Printf("Fail to convert timestamp (%s) of media\n", m.PhotoTakenTime.Timestamp)
		return
	}
	wrongTime, err := getFileModTime(dir, mediaName)
	if err != nil {
		log.Printf("Fail to read media (%s) modification time\n", mediaName)
		return
	}
	log.Printf("set file time from %s to %s", wrongTime.String(), takenTime.String())
	err = os.Chtimes(filepath.Join(dir, mediaName), takenTime, takenTime)
	if err == nil {
		log.Println("modification time set successfully!")
		return
	}

	log.Println("Fail to set modification time")
	return
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "./testdata", "The folder that contains photo/video and json metadata from Google Takeout")
	flag.Parse()

	log.Printf("Folder to process: %s\n", dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			err := setMetadata(dir, file)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

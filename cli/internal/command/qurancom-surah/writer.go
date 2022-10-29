package quranComSurah

import (
	"data-quran-cli/internal/util"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func cleanDstDir(dstDir string) error {
	return filepath.WalkDir(dstDir, func(path string, d fs.DirEntry, err error) error {
		// Remove all file suffixed with "-qurancom.json"
		dName := d.Name()
		if d.IsDir() || !strings.HasSuffix(dName, "-qurancom.json") {
			return nil
		}

		dDir := filepath.Base(filepath.Dir(path))
		switch dDir {
		case "surah-info", "surah-translation":
			return os.Remove(path)
		}

		return nil
	})
}

func writeListSurah(dstDir string, language string, data map[string]ListSurahOutput) error {
	// If data is empty, stop
	if len(data) == 0 {
		return nil
	}

	logrus.Printf("writing surah list for %s", language)

	// Prepare destination path
	dstDir = filepath.Join(dstDir, "surah-translation")
	os.MkdirAll(dstDir, os.ModePerm)

	dstPath := fmt.Sprintf("%s-qurancom.json", language)
	dstPath = filepath.Join(dstDir, dstPath)

	// Encode data
	err := util.EncodeSortedKeyJson(dstPath, &data)
	if err != nil {
		return fmt.Errorf("fail to write surah list for %s: %w", language, err)
	}

	return nil
}

func writeSurahInfo(dstDir string, language string, data *AllSurahInfoOutput) error {
	// If data is empty, stop
	if data == nil {
		return nil
	}

	logrus.Printf("writing surah info for %s", language)

	// Prepare destination path
	dstDir = filepath.Join(dstDir, "surah-info")
	os.MkdirAll(dstDir, os.ModePerm)

	dstPath := fmt.Sprintf("%s-qurancom.md", language)
	dstPath = filepath.Join(dstDir, dstPath)

	// Open destination file
	f, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("create file for surah info %s failed: %w", language, err)
	}
	defer f.Close()

	// Write metadata
	f.WriteString("<!--\n")
	f.WriteString(fmt.Sprintf("Language: %s\n", data.Language))
	f.WriteString(fmt.Sprintf("Source  : %s\n", data.Source))
	f.WriteString("-->\n\n")

	// Write each info
	for surah := 1; surah <= 114; surah++ {
		f.WriteString("# ")
		f.WriteString(strconv.Itoa(surah))
		f.WriteString("\n\n")

		text := data.Texts[surah]
		if text == "" {
			f.WriteString("<!-- TODO:MISSING -->\n\n")
			continue
		}

		f.WriteString(text)
		f.WriteString("\n\n")
	}

	// Flush the data
	err = f.Sync()
	if err != nil {
		return fmt.Errorf("write file for surah info %s failed: %w", language, err)
	}

	return nil
}
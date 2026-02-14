package photo

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/buckket/go-blurhash"

	"share/types/extern"
	"share/types/model"
	"share/utils/value"

	_ "golang.org/x/image/tiff"

	"github.com/rwcarlsen/goexif/exif"

	"command/procedures/file"
)

type Subcommand struct {
	Dir string
}

func (r *Subcommand) Name() string {
	return "photo"
}

func (r *Subcommand) Parse(args []string) error {
	fs := flag.NewFlagSet(r.Name(), flag.ContinueOnError)
	fs.StringVar(&r.Dir, "dir", "", "Dir of the photo")

	return fs.Parse(args)
}

func (r *Subcommand) Run() error {
	// fmt.Printf("INFO | Variable DIR: %s\n", r.Dir)

	files, err := file.GetDscFiles(r.Dir)
	if err != nil {
		log.Fatal(err)
	}

	var photos []*model.PhotoItem

	fmt.Println("photo_section_id;shooter_id;image_path;thumbnail_path;raw_path;blurhash;exif;created_at;updated_at")
	for i, fileInfo := range files {
		if i == -1 {
			continue
		}

		imgFile, err := os.Open(filepath.Join(r.Dir, fileInfo.Name()))
		if err != nil {
			log.Fatal(err.Error())
		}

		// * Parse EXIF
		metaData, err := exif.Decode(imgFile)
		if err != nil {
			log.Fatal(err.Error())
		}

		jsonByte, err := metaData.MarshalJSON()
		if err != nil {
			log.Fatal(err.Error())
		}

		var exifRaw *ExifRaw
		err = json.Unmarshal(jsonByte, &exifRaw)

		// * Reset file pointer
		_, err = imgFile.Seek(0, 0)
		if err != nil {
			return err
		}

		// * Parse image
		img, _, err := image.Decode(imgFile)
		if err != nil {
			log.Fatal(err.Error())
		}

		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()

		if exifRaw.Orientation != nil {
			switch exifRaw.Orientation[0] {
			case 1:
				width = bounds.Dx()
				height = bounds.Dy()
			case 3:
				width = bounds.Dx()
				height = bounds.Dy()
			case 6:
				width = bounds.Dy()
				height = bounds.Dx()
			case 8:
				width = bounds.Dy()
				height = bounds.Dx()
			}
		}

		w := width
		h := height

		for w >= 10 && h >= 10 {
			w /= 10
			h /= 10
		}

		if w > 10 {
			w = 9
		}
		if h > 10 {
			h = 9
		}

		hash, err := blurhash.Encode(w, h, img)
		if err != nil {
			log.Fatal(err.Error())
		}

		t, err2 := time.Parse("2006:01:02 15:04:05 -0700", exifRaw.DateTimeOriginal+" +0700")
		if err2 != nil {
			log.Fatal(err.Error())
		}

		var aperture = "1.2"
		if exifRaw.ApertureValue != nil {
			aperture = strings.Split(exifRaw.ApertureValue[0], "/")[0]
		} else if exifRaw.FNumber != nil {
			frac1, _ := strconv.ParseInt(strings.Split(exifRaw.FNumber[0], "/")[0], 10, 64)
			frac2, _ := strconv.ParseInt(strings.Split(exifRaw.FNumber[0], "/")[1], 10, 64)
			aperture = fmt.Sprintf("%.1f", float64(frac1)/float64(frac2))
		}

		fl := "35"
		flff := "52"
		if exifRaw.FocalLength != nil {
			foc := strings.Split(exifRaw.FocalLength[0], "/")
			fl = foc[0]
			if len(foc) > 1 {
				foc0, _ := strconv.ParseInt(foc[0], 10, 32)
				foc1, _ := strconv.ParseInt(foc[1], 10, 32)
				fl = fmt.Sprintf("%d", foc0/foc1)
			}
			if len(exifRaw.FocalLengthIn35mmFilm) > 0 {
				flff = strconv.Itoa(exifRaw.FocalLengthIn35mmFilm[0])
			} else {
				flff = fl
			}
		} else {
			exifRaw.LensModel = "RISESPRAY 35mm F1.2"
		}

		if exifRaw.LensModel == "----" {
			exifRaw.LensModel = "RISESPRAY 35mm F1.2"
		}

		if len(exifRaw.IsoSpeedRatings) == 0 {
			exifRaw.LensModel = "N/A"
			exifRaw.IsoSpeedRatings = []int{0}
			exifRaw.ExposureTime = []string{"N/A"}
		}

		currentTime := time.Now()
		camera := exifRaw.Make + " " + exifRaw.Model
		if strings.HasPrefix(exifRaw.Model, "NIKON") {
			camera = exifRaw.Model
		}

		photos = append(photos, &model.PhotoItem{
			Id:             nil,
			PhotoSection:   nil,
			PhotoSectionId: value.Ptr[uint64](8),
			Shooter:        nil,
			ShooterId:      value.Ptr[uint64](1),
			ImagePath:      value.Ptr("img/" + fileInfo.Name()),
			ThumbnailPath:  value.Ptr("tmb/" + strings.Split(fileInfo.Name(), ".")[0] + ".jpg"),
			RawPath:        value.Ptr("raw/" + strings.Split(fileInfo.Name(), ".")[0] + ".arw"),
			Blurhash:       &hash,
			Exif: &extern.Exif{
				Aperture:      aperture,
				Timestamp:     t.UTC(),
				ShutterSpeed:  exifRaw.ExposureTime[0],
				FocalLength:   fl,
				FocalLengthFF: flff,
				Iso:           strconv.Itoa(exifRaw.IsoSpeedRatings[0]),
				Lens:          exifRaw.LensModel,
				Camera:        camera,
				Software:      exifRaw.Software,
				Width:         width,
				Height:        height,
			},
			CreatedAt: &currentTime,
			UpdatedAt: &currentTime,
		})

		bytes, err := json.Marshal(photos[i].Exif)
		fmt.Printf("%d;%d;%s;%s;%s;\"%s\";%s;%s;%s\n", *photos[i].PhotoSectionId, *photos[i].ShooterId, *photos[i].ImagePath, *photos[i].ThumbnailPath, *photos[i].RawPath, *photos[i].Blurhash, string(bytes), photos[i].CreatedAt.Format(time.RFC3339), photos[i].UpdatedAt.Format(time.RFC3339))
		// fmt.Printf("UPDATE photo_items SET blurhash = '%s' WHERE image_path = '%s'\n", *photos[i].Blurhash, *photos[i].ImagePath)
	}

	// jsbyte, _ := json.Marshal(photos)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	//
	// fmt.Println(string(jsbyte))
	return nil
}

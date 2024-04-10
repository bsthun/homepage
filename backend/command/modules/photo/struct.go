package photo

type ExifRaw struct {
	ApertureValue         []string `json:"ApertureValue"`
	FNumber               []string `json:"FNumber"`
	BrightnessValue       []string `json:"BrightnessValue"`
	Artist                string   `json:"Artist"`
	Copyright             string   `json:"Copyright"`
	DateTimeOriginal      string   `json:"DateTimeOriginal"`
	ExposureTime          []string `json:"ExposureTime"`
	FocalLength           []string `json:"FocalLength"`
	FocalLengthIn35mmFilm []int    `json:"FocalLengthIn35mmFilm"`
	IsoSpeedRatings       []int    `json:"ISOSpeedRatings"`
	LensModel             string   `json:"LensModel"`
	Make                  string   `json:"Make"`
	Model                 string   `json:"Model"`
	PixelXDimension       []int    `json:"PixelXDimension"`
	PixelYDimension       []int    `json:"PixelYDimension"`
	Orientation           []int    `json:"Orientation"`
	Software              string   `json:"Software"`
	WhiteBalance          []int    `json:"WhiteBalance"`
}

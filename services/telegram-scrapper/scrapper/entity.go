package scrapper

// Source describes target url and maximum count of parsing animals
type Source struct {
	Url        string
	MaxRecords int
}

// Animal received structure after parsing web page.
type Animal struct {
	ID     int      `json:"id"`
	Text   string   `json:"text"`
	Photos []string `json:"photos"`
}

// Settings describes scrapper environment
type Settings struct {
	Driver string `json:"CHROME_DRIVER_PATH" mapstructure:"CHROME_DRIVER_PATH"`
	Port   int    `json:"PARSER_PORT" mapstructure:"PARSER_PORT"`
}

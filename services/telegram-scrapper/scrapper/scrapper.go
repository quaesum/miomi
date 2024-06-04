package scrapper

import (
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	environment "madmax/sdk/environment/config"
	"time"
)

type Scrapper struct {
	selenium.WebDriver
	Sources  []Source
	Settings Settings
}

// New creates a new instance of scrapper
func New() *Scrapper {
	v, err := environment.LoadConfig(".env")
	if err != nil {
		log.Fatal(err)
	}

	var settings Settings
	err = v.Unmarshal(&settings)
	if err != nil {
		log.Fatal(err)
	}

	opts := []selenium.ServiceOption{
		selenium.Output(nil),
	}
	_, err = selenium.NewChromeDriverService(settings.Driver, settings.Port, opts...)
	if err != nil {
		log.Fatal(err)
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", settings.Port))
	if err != nil {
		log.Fatal(err)
	}
	return &Scrapper{
		WebDriver: wd,
		Settings:  settings,
	}
}

// AddSource adds new target url with records limit
func (s *Scrapper) AddSource(url string, maxRecords int) {
	s.Sources = append(s.Sources, Source{Url: url, MaxRecords: maxRecords})
}

// Run starts scrapper
func (s *Scrapper) Run() {
	for _, el := range s.Sources {
		processedPosts := make(map[int]struct{})
		var factor = 0.5
		var animals []Animal
		count := 0
		for count < el.MaxRecords {
			animals = nil
			html, err := s.GetHTML(el, factor)
			if err != nil {
				log.Fatal(err)
			}
			part, err := ProcessHTML(html, el.MaxRecords, processedPosts, count)
			if err != nil {
				log.Fatal(err)
			}
			count += len(part)
			if count > el.MaxRecords {
				over := el.MaxRecords - len(animals)
				animals = append(animals, part[:over]...)
			} else {
				animals = append(animals, part...)
			}
			log.Println(len(animals), count)

			factor += 0.5
		}
		var filteredAnimals []Animal
		for _, item := range animals {
			if item.Photos != nil {
				filteredAnimals = append(filteredAnimals, item)
			}
		}
		err := writeToFile(filteredAnimals)
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

// GetHTML parsing page and returns html code as string
func (s *Scrapper) GetHTML(el Source, factor float64) (string, error) {
	err := s.Get(el.Url)
	if err != nil {
		log.Println("Error while getting page:", err)
		return "", err
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < el.MaxRecords; i++ {
		if i%10 == 0 {
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(1 * time.Millisecond)
		}
		_, err = s.ExecuteScript("window.scrollBy(0, -700	)", nil)
		if err != nil {
			log.Println("Scroll error: ", err, "\nRetrying...")
		}
	}

	html, err := s.PageSource()
	if err != nil {
		log.Println("Failed to get html:", err)
		return "", err
	}
	return html, nil
}

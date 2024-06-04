package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
	"log"
	environment "madmax/sdk/environment/config"
	"madmax/services/telegram-scrapper/scrapper"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Gemini struct {
	*genai.GenerativeModel
}

func New() *Gemini {
	v, err := environment.LoadConfig(".env")
	if err != nil {
		log.Fatal(err)
	}

	var settings Settings
	err = v.Unmarshal(&settings)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(settings.ApiKey))
	if err != nil {
		log.Fatal(err)
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetMaxOutputTokens(20000)
	return &Gemini{
		model,
	}

}

func (g *Gemini) Process() {
	var wg sync.WaitGroup
	log.Println("preparing request")
	splitAnimals, err := prepareRequest()
	if err != nil {
		log.Fatal(err)
	}

	fileShort, err := os.Open("animals.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fileShort.Close()

	dataShort, err := io.ReadAll(fileShort)
	if err != nil {
		log.Fatal(err)
	}

	var animalsShort []scrapper.Animal
	err = json.Unmarshal(dataShort, &animalsShort)
	if err != nil {
		log.Fatal(err)
	}

	photosMap := make(map[int][]string)
	for _, animal := range animalsShort {
		photosMap[animal.ID] = animal.Photos
	}

	log.Println("sending text to gemini")
	var animals []AnimalResponse
	//for _, item := range splitAnimals {
	//	var animalsPart []AnimalResponse
	//	str := g.TextToJSON(item)
	//	str = strings.Trim(str, "```")
	//	str = strings.TrimLeft(str, "json")
	//	fmt.Println(str)
	//	err = json.Unmarshal([]byte(str), &animalsPart)
	//	if err != nil {
	//		log.Println("error while unmarshalling animal part", err)
	//		continue
	//	}
	//	log.Println("completing animals part")
	//	animals = append(animals, animalsPart...)
	//}

	for i := range splitAnimals {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			item := splitAnimals[i]
			var animalsPart []AnimalResponse
			str := g.TextToJSON(item)
			str = strings.Trim(str, "```")
			str = strings.TrimLeft(str, "json")
			fmt.Println(str)
			err = json.Unmarshal([]byte(str), &animalsPart)
			if err != nil {
				log.Println("error while unmarshalling animal part", err)
			}
			log.Println("completing animals part")
			animals = append(animals, animalsPart...)
		}(i)
	}
	wg.Wait()

	//for _, animal := range animals {
	//	if animal.Type == "" {
	//		log.Printf("type not found in %d, requesting to gemini", animal.ID)
	//		if photos, exists := photosMap[animal.ID]; exists {
	//			animal.Type = g.fillTypeByPhoto(photos[0])
	//		}
	//
	//	}
	//}

	//for i := range animals {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		animal := &animals[i]
	//		if animal.Type == "" {
	//			log.Printf("Type not found for animal %d, requesting to Gemini", animal.ID)
	//			if photos, exists := photosMap[animal.ID]; exists && len(photos) > 0 {
	//				animal.Type = g.fillTypeByPhoto(photos[0])
	//			}
	//		}
	//	}(i)
	//}
	//
	//wg.Wait()

	dataJSON, err := json.Marshal(animals)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("writing to file")
	err = writeToFile(string(dataJSON))
}

func (g *Gemini) TextToJSON(text string) string {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Minute*2)
	instructions := &genai.Content{
		Parts: []genai.Part{
			genai.Text(prompt),
		},
		Role: "system",
	}
	g.SystemInstruction = instructions

	resp, err := g.GenerateContent(tctx, genai.Text(text))
	if err != nil {
		log.Fatalf("error sending message: %v", err)
	}
	fmt.Println(resp.UsageMetadata.TotalTokenCount)

	for _, part := range resp.Candidates[0].Content.Parts {
		if v, ok := part.(genai.Text); ok {
			return string(v)
		}
	}
	return ""
}

func prepareRequest() ([]string, error) {
	file, err := os.Open("animals.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var animals []AnimalRequest
	err = json.Unmarshal(data, &animals)
	if err != nil {
		return nil, err
	}

	var splitAnimals []string
	size := len(animals)
	part := 50

	for i := 0; i < size; i += part {
		end := i + part
		if end > size {
			end = size
		}
		an, err := json.Marshal(animals[i:end])
		if err != nil {
			return nil, err
		}
		splitAnimals = append(splitAnimals, string(an))
	}

	return splitAnimals, nil
}

func completeFields(animals []AnimalResponse) {
	for _, animal := range animals {
		if animal.Name == "" {
			if animal.Type == "dog" {

			}
		}
	}
}

func writeToFile(animals string) error {
	err := os.WriteFile("animals_gemini.json", []byte(animals), 0666)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gemini) fillTypeByPhoto(url string) string {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Minute*30)
	instructions := &genai.Content{
		Parts: []genai.Part{
			genai.Text(photo_prompt),
		},
		Role: "system",
	}
	g.SystemInstruction = instructions

	imgData, err := downloadImage(url)
	if err != nil {
		log.Fatal(err)
	}

	img := []genai.Part{
		genai.ImageData("jpeg", imgData),
	}

	resp, err := g.GenerateContent(tctx, img...)
	if err != nil {
		log.Printf("error sending message: %v \n%v\n", err, url)
		return ""
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		if v, ok := part.(genai.Text); ok {
			return string(v)
		}
	}

	return ""
}

func downloadImage(url string) ([]byte, error) {
	// Загружаем изображение по URL
	resp, err := getWithRetry(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Считываем содержимое ответа
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Создаём временный файл и записываем в него содержимое изображения
	tmpFile, err := ioutil.TempFile("", "image_")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name()) // Удаление временного файла после использования

	_, err = tmpFile.Write(imgData)
	if err != nil {
		return nil, err
	}

	// Читаем содержимое временного файла и возвращаем его
	_, err = tmpFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	imgData, err = ioutil.ReadAll(tmpFile)
	if err != nil {
		return nil, err
	}

	return imgData, nil
}

func retryDo(fn func() error, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		err := fn()
		if err == nil {
			return nil
		}
		log.Printf("Attempt %d/%d failed: %v\n", i+1, maxRetries, err)
		time.Sleep(2 * time.Second) // Optional: Add some delay between retries
	}
	return errors.New("all retry attempts failed")
}

// getWithRetry performs an HTTP GET request with retries.
func getWithRetry(url string) (*http.Response, error) {
	var resp *http.Response
	var err error

	err = retryDo(func() error {
		resp, err = http.Get(url)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		return nil
	}, 3)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

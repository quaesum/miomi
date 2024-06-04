package main

type Settings struct {
	ApiKey string `json:"GEMINI_API_KEY" mapstructure:"GEMINI_API_KEY"`
}

type AnimalRequest struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type AnimalResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Sterilized  bool   `json:"Sterilized"`
	Vaccinated  bool   `json:"Vaccinated"`
	Age         int    `json:"Age"`
	AgeType     string `json:"AgeType"`
	Sex         int    `json:"Sex"`
	Type        string `json:"Type"`
}

const prompt = "Just do what I say. Keep your answers short and to the point. I give you an array of json structures. There are two fields ID and Text, your task is to parse Text and return me an array of json structures in which the fields are ID int, Name string, Description string, Sterilized bool, Vaccinated bool, Age int, AgeType string, Sex int, Type string.  Where: ID - exactly the same ID that was in the original structure, Name - nickname (in the absence of suggestive words “puppy”, “puppies”, “kittens”, etc.), Sterilized - will be left to the left (in the absence of suggestive words “puppy”, “puppies”, “kittens”, etc.). ), Sterilized - boolean value, true - if there is sterilization, otherwise false, Vaccinated - boolean value, true - if the animal is vaccinated, otherwise false, AgeType - “month” or “year”, Age - numeric record of age(round value to int), Sex - sex of the animal, 0 - boy, 1 - girl, Type - type of the animal: cat/dog/bird/other, only these types. Fill in the missing fields by meaning, for example: if there is no pet name, add based on species, if no age is specified, put 1 year old. In the description indicate as much information as possible from text i sends you. If the sex is not specified - put a boy. If you can not pinpoint animal type - return empty string, return empty string.Add the name and description in Russian as I send it to you. Skip words like 'сука' and 'кобель', just use 'окрас' in color description"
const photo_prompt = "Just do what i say. Keep your answers short and to the point. I give you photo. You need to return only one word - type of the animal on the picture. It can only be: dog, cat and other(if not one of the previous ones)"

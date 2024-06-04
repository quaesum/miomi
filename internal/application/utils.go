package application

import (
	"encoding/json"
	"log"
	"madmax/internal/entity"
	"madmax/internal/utils"
	"madmax/sdk/nats"
	mse "madmax/services/mail/entity"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func calculateServicesScore(service entity.Service, searchTerm string) int64 {
	searchTerm = strings.ToLower(searchTerm)
	nameCount := strings.Count(strings.ToLower(service.Name), searchTerm)
	descriptionCount := strings.Count(strings.ToLower(service.Description), searchTerm)
	return int64(nameCount + descriptionCount)
}

func calculateProductsScore(service entity.Product, searchTerm string) int64 {
	searchTerm = strings.ToLower(searchTerm)
	nameCount := strings.Count(strings.ToLower(service.Name), searchTerm)
	descriptionCount := strings.Count(strings.ToLower(service.Description), searchTerm)
	return int64(nameCount + descriptionCount)
}

func generateEmailVerificationToken(userID int64, userBI *entity.User) (string, error) {
	return utils.GenerateToken(strconv.FormatInt(userID, 10), userBI.Role)
}

func sendEmailVerificationMessage(email, token string) error {

	m := nats.Request{
		Code: mse.ConfirmEmailType,
		To:   email,
		CommonData: mse.CommonData{
			Token: token,
		},
	}
	req, err := json.Marshal(m)
	if err != nil {
		return err
	}

	re, err := mse.Nats.Request("mail_topic", req, 10*time.Second)
	if err != nil {
		log.Println("failed to send email confirmation")
		return err
	}

	resp := nats.Response{}
	err = json.Unmarshal(re.Data, &resp)

	if err != nil {
		log.Println("failed to unmarshal nats response")
		return err
	}

	if resp.Code != http.StatusOK {
		log.Println("response isn't correct")
	}
	return nil
}

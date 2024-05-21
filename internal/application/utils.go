package application

import (
	"context"
	"encoding/json"
	"github.com/wagslane/go-rabbitmq"
	"log"
	"madmax/internal/entity"
	"madmax/internal/utils"
	mse "madmax/services/mail/entity"
	"strconv"
	"strings"
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
	msg := mse.MailQueMessage{
		Type: mse.ConfirmEmailType,
		To:   email,
		CommonData: mse.CommonData{
			Token: token,
		},
	}
	b, err := json.Marshal(msg)
	if err != nil {
		log.Printf("error marshalling email verification message: %v", err)
		return err
	}

	err = mse.EmailPublisher.PublishWithContext(
		context.Background(),
		b,
		[]string{"meme.emails"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("emails"),
	)
	if err != nil {
		log.Printf("error publishing email verification message: %v", err)
	}
	return err
}

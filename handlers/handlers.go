package handlers

import (
	"fmt"
	"github.com/IdekDude/webhookAPI/models"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"sync"
)

var clientReq models.Webhook
var webhooks = make(chan []byte, 2)
func SendWebhook(c *fiber.Ctx) error {
	err := c.BodyParser(&clientReq)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}
	//sendWebhook(c.Body())
	webhooks <- c.Body()
	go WebhookScheduler()
	return nil
}

func WebhookScheduler() {
	var webhooksSent = 0
	var wg sync.WaitGroup
	for webhook := range webhooks {
		wg.Add(1)
		webhooksSent++
		go sendWebhook(webhook, &wg)
		if webhooksSent > 2 {
			wg.Wait()
			webhooksSent--
		}
	}
}

func sendWebhook(webhook []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	hooks := LoadWebhooks("WEBHOOKS")

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	for _, url := range hooks {
		req.SetRequestURI(url)
		req.SetBody(webhook)
		req.Header.SetMethod("POST")
		req.Header.Set("content-type", "application/json")


		err := fasthttp.Do(req, resp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.StatusCode())
	}
}
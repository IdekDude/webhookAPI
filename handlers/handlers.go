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
	fmt.Println("hello")
	var wg sync.WaitGroup
	for webhook := range webhooks {
		fmt.Println(webhooksSent)
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
/*	jsonHook := new(bytes.Buffer)
	err = json.NewEncoder(jsonHook).Encode(webhook)
	if err != nil {
		return
	}

	sendWebhook, err := http.NewRequest("POST", "https://discord.com/api/webhooks/859444343618535454/Q3ADYUqSfzKI-sYqFFcUVi85TwniLGEtjHQwF-tfECloSZP6P-QafULpriNk-Q4R5YmU", jsonHook)
	if err != nil {
		log.Fatal(err)
	}

	sendWebhook.Header.Set("content-type", "application/json")

	time.Sleep(2 * time.Second)

	sendWebhookRes, err := client.Do(sendWebhook)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sendWebhookRes)*/
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
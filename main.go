package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"first_golang_project/config"
	"first_golang_project/routes"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CreditTransfer struct (same as before)
type CreditTransfer struct {
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

// CertificatePayload for receiving PEM string
type CertificatePayload struct {
	PEM string `json:"pem"`
}

func main() {
	config.Init() // <- MUST call Init first
	app := fiber.New()

	// Sample scripts
	// Existing credit-transfer endpoint
	app.Post("/credit-transfer", func(c *fiber.Ctx) error {
		processStart := time.Now()
		locationPH, _ := time.LoadLocation("Asia/Manila")
		ctCreationDateTime := processStart.In(locationPH).Format("2006-01-02 15:04:05")

		CreditTransferData := new(CreditTransfer)
		if err := c.BodyParser(CreditTransferData); err != nil {
			return c.Status(400).SendString("Invalid payload: " + err.Error())
		}

		ctRequest, _ := json.Marshal(CreditTransferData)
		fmt.Println("JSON Payload:", string(ctRequest))

		return c.JSON(fiber.Map{
			"message":   "Credit transfer processed successfully",
			"payload":   CreditTransferData,
			"timestamp": ctCreationDateTime,
		})
	})

	// New endpoint: parse a PEM certificate
	app.Post("/parse-certificate", func(c *fiber.Ctx) error {
		payload := new(CertificatePayload)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(400).SendString("Invalid payload: " + err.Error())
		}

		// Decode PEM
		block, _ := pem.Decode([]byte(payload.PEM))
		if block == nil {
			return c.Status(400).SendString("Failed to parse PEM block")
		}

		// Parse x509 certificate
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return c.Status(400).SendString("Failed to parse certificate: " + err.Error())
		}

		// Return some info about the certificate
		info := map[string]interface{}{
			"subject":         cert.Subject,
			"issuer":          cert.Issuer,
			"not_before":      cert.NotBefore,
			"not_after":       cert.NotAfter,
			"serial_number":   cert.SerialNumber.String(),
			"dns_names":       cert.DNSNames,
			"email_addresses": cert.EmailAddresses,
		}

		return c.JSON(fiber.Map{
			"message": "Certificate parsed successfully",
			"cert":    info,
		})

	})

	// register routes
	routes.UserRoutes(app)
	routes.SupplierRoutes(app)

	fmt.Println("Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}

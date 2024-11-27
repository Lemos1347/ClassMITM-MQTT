package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Lemos1347/ClassMITM-MQTT/internal/caesar_cipher"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func init() {
	envPath := flag.String("env", "./.env", "Path for .env file")
	flag.Parse()
	godotenv.Load(*envPath)
}

func main() {
	broker, ok := os.LookupEnv("BROKER_URL")
	if !ok {
		log.Fatal("BROKER_URL not set")
	}
	clientID := "go-mqtt-publisher"
	topic := "/listener"
	username, ok := os.LookupEnv("BROKER_USERNAME")
	if !ok {
		log.Fatal("BROKER_USERNAME not set")
	}
	password, ok := os.LookupEnv("BROKER_PASSWORD")
	if !ok {
		log.Fatal("BROKER_PASSWORD not set")
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected to MQTT broker!")
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Erro while connecting to broker: %v\n", token.Error())
		return
	}

	defer client.Disconnect(250)

	shift := rand.Intn(26)
	caesarCipher := caesar_cipher.NewCaeserCipher(shift)
	log.Printf("Using shift: %d\n", shift)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Write message you wanna send (or 'quit' to exit program):")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if strings.ToLower(message) == "quit" {
			fmt.Println("Finishing program...")
			break
		}

		encryptedMessage := caesarCipher.Encrypt(message)

		token := client.Publish(topic, 0, false, encryptedMessage)
		token.Wait()

		fmt.Println("Message sent:", message)
		time.Sleep(1 * time.Second)
	}
}

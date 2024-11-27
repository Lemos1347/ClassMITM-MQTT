package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Lemos1347/ClassMITM-MQTT/internal/caesar_cipher"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var shitftPtr *int

func init() {
	envPath := flag.String("env", "./.env", "Path for .env file")
	shitftPtr = flag.Int("shift", 0, "Value used for caeser cipher shift")
	flag.Parse()
	godotenv.Load(*envPath)
}

func main() {
	if shitftPtr == nil {
		log.Fatal("Shift value not defined!")
	}

	shift := *shitftPtr

	if shift == 0 {
		fmt.Println("Erro: you must provide a valid shift value with '-shift'")
		os.Exit(1)
	}

	broker, ok := os.LookupEnv("BROKER_URL")
	if !ok {
		log.Fatal("BROKER_URL not set")
	}
	clientID := "go-mqtt-listener"
	topic := "/caller"
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
		client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			caeserCipher := caesar_cipher.NewCaeserCipher(shift)

			msgDecrypted := caeserCipher.Decrypt(string(msg.Payload()))
			log.Printf("Message received: %s\n", msgDecrypted)
		})
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Erro while connecting to broker: %v\n", token.Error())
		return
	}

	fmt.Println("Waiting for messages... Press Ctrl+C to finish.")
	select {}
}

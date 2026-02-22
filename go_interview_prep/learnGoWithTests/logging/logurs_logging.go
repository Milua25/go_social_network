package main

import (
	"github.com/sirupsen/logrus"
)
func main(){

	logger := logrus.New()

	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	// logging examples
	logger.Info("This is an info type of message")
	logger.Warn("This is a waring type of message")
	logger.Error("This is an error message")

	logger.WithFields(logrus.Fields{
		"username": "John Doe", 
		"method": "GET",
	}).Info("User logged in")

}
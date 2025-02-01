package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// calculate the amount of points earned
func (receipt Receipt) calcPoints() int {
	points := 0
	// One point for every alphanumeric character in the retailer name.
	for _, val := range receipt.Retailer {
		if unicode.IsLetter(val) || unicode.IsDigit(val) {
			points++
		}
	}
	// 50 points if the total is a round dollar amount with no cents.
	totalSplit := strings.Split(receipt.Total, ".")
	cents, err := strconv.Atoi(totalSplit[1])
	if err != nil {
		log.Println(err)
		return points
	}
	if cents == 0 {
		points += 50
	}
	// 25 points if the total is a multiple of 0.25.
	if cents%25 == 0 {
		points += 25
	}
	// 5 points for every two items on the receipt.
	points += (5 * (len(receipt.Items) / 2))
	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		itemDesc := strings.TrimSpace(item.ShortDescription)
		if len(itemDesc)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				log.Println(err)
				return points
			}
			points += int(math.Ceil(price * float64(0.2)))
		}
	}
	// 6 points if the day in the purchase date is odd.
	dateSplit := strings.Split(receipt.PurchaseDate, "-")
	dayOfMonth, err := strconv.Atoi(dateSplit[2])
	if err != nil {
		log.Println(err)
		return points
	}
	if dayOfMonth%2 == 1 {
		points += 6
	}
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	timeSplit := strings.Split(receipt.PurchaseTime, ":")
	time, err := strconv.Atoi(timeSplit[0] + timeSplit[1])
	if err != nil {
		log.Println(err)
		return points
	}
	if time >= 1400 && time < 1600 {
		points += 10
	}
	return points
}

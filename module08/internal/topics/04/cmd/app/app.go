package main

import (
	"Rebrain/module08/internal/topics/04/internal/models"
	"fmt"
	"github.com/goombaio/namegenerator"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

func Start(db *gorm.DB) {
	//createUser(db)
	//selectAndUpdate(db)

	//CreateRandomUsers(db, 30)
	firstSelect(db)
	secondSelect(db)
	thirdSelect(db)

}

func selectAndUpdate(db *gorm.DB) {
	user := &models.User{}

	db.Where("name = ?", "Nikolai").First(&user)
	if user.ID > 0 {
		user.Name = "Dmitri"
		db.Save(user)
	}
}

// спорный....
func firstSelect(db *gorm.DB) {
	fmt.Println("FIRST SELECT")
	var ids []int
	db.Model(&models.Card{}).Select("user_id").Group("user_id").Having("COUNT(*) = 3").Find(&ids)

	var users []models.User
	db.Model(&models.User{}).Where("id IN ?", ids).Find(&users)
	fmt.Println(users)
}

func secondSelect(db *gorm.DB) {
	fmt.Println("SECOND SELECT")
	var users []models.User
	db.Model(&models.User{}).Joins("inner join cards on users.id = cards.user_id").Where(
		"cards.type = 'VISA'").Find(&users)
	fmt.Println(users)
}

func thirdSelect(db *gorm.DB) {
	fmt.Println("THIRD SELECT")
	var users []models.User
	db.Model(&models.User{}).Joins("inner join cards on users.id = cards.user_id").Where(
		"cards.type = 'MIR' AND users.age > 50").Find(&users)
	fmt.Println(users)
}

func createUser(db *gorm.DB) {
	fmt.Println("CREATE USER")
	user := &models.User{
		Name:     "Amigo",
		Age:      30,
		IsVerify: true,
		Cards: []models.Card{models.Card{
			Number: "34534534534",
			Type:   "VISA",
		}},
	}
	result := db.Create(user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func CreateRandomUsers(db *gorm.DB, n int) {
	for i := 0; i < n; i++ {
		CreateRandomUser(db)
	}
}

func CreateRandomUser(db *gorm.DB) {
	randomUser := generateRandomUser()
	result := db.Create(randomUser)
	if result.Error != nil {
		panic(result.Error)
	}
}

func generateRandomUser() *models.User {
	user := &models.User{}
	user.Name = randomName()
	user.Age = randomAge()
	user.IsVerify = randomVerify()
	user.Cards = randomCards()

	return user
}

func randomName() string {
	nameGenerator := namegenerator.NewNameGenerator(time.Now().UnixNano())
	name := nameGenerator.Generate()
	return name
}

func randomAge() int {
	age := rand.Intn(53) + 18
	return age
}

func randomVerify() bool {
	return rand.Intn(2) == 1
}

func randomCards() []models.Card {
	cardsQuantity := rand.Intn(3) + 1
	cards := make([]models.Card, cardsQuantity)

	for i := 0; i < cardsQuantity; i++ {
		cards[i] = randomCard()
	}

	return cards
}

func randomCard() models.Card {
	types := []string{"VISA", "MIR", "MASTERCARD"}
	card := models.Card{}

	card.Type = types[rand.Intn(3)]
	card.Number = strconv.Itoa(rand.Intn(90000000000) + 1000000000)
	return card
}

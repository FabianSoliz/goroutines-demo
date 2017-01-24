//Controller for managed the resourceOrder of a user
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/mercadolibre/fury_messages-unread-api/src/api/service"
	"time"
	"math/rand"
)



var ResourceService service.ResourceService



//Return purchase
func GetByPurchaseId(c *gin.Context) {

	purchaseId, _ := strconv.Atoi(c.Param("purchase_id"))

	chanTask := make(chan string)

	var responsePurchase []string

	go func() {chanTask <- task(1000 + purchaseId)}()
	go func() {chanTask <- task(2000 + purchaseId)}()
	go func() {chanTask <- task(3000 + purchaseId)}()

	//go rutine con tareas encadenadas
	go func() {

		chanSubTask := make(chan string)

		go func () {chanSubTask <- subtask(4000 + purchaseId)}()
		responseSubTask1 := <-chanSubTask

		go func() {chanSubTask <- task(4500 + purchaseId)}()
		responseSubTask2 := <-chanSubTask

		chanTask <- (responseSubTask1 + "," + responseSubTask2)
	}()

	go func() {chanTask <- task(5000 + purchaseId)}()

	//go rutine con tareas encadenadas
	go func() {

		chanSubTask := make(chan string)

		go func () {chanSubTask <- subtask(6000 + purchaseId)}()
		responseSubTask1 := <-chanSubTask

		go func() {chanSubTask <- task(6500 + purchaseId)}()
		responseSubTask2 := <-chanSubTask

		chanTask <- (responseSubTask1 + "," + responseSubTask2)
	}()

	go func() {chanTask <- task(7000 + purchaseId)}()
	go func() {chanTask <- task(8000 + purchaseId)}()


	for i := 0; i < 8; i++ {
		response := <-chanTask

		responsePurchase = appendResponse(responsePurchase, response)
	}


	c.JSON(http.StatusOK, responsePurchase)
}

func task(id int) string {
	time.Sleep(time.Duration(rand.Int31n(200)) * time.Millisecond)


	return `{"name": "task` + strconv.Itoa(id) + `"}`
}

func subtask(id int) string {
	time.Sleep(time.Duration(rand.Int31n(200)) * time.Millisecond)


	return `{"name": "subtask` + strconv.Itoa(id) + `"}`
}


func appendResponse(original []string, response string) []string {
	return append(original, response)
}
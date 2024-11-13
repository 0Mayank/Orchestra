package main

import (
	pauths "Client/auths"
	pbooking "Client/booking"
	ppayment "Client/payment"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var CLIENT_ID string
var HOTEL_ID int64
var PAYMENT_ID string
var PORT int64
var BOOKING_HOST string
var BOOKING_PORT int64
var PAYMENT_HOST string
var PAYMENT_PORT int64
var AUTH_HOST string
var AUTH_PORT int64

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	CLIENT_ID = os.Getenv("CLIENT_ID")
	HOTEL_ID, err = strconv.ParseInt(os.Getenv("HOTEL_ID"), 10, 64)
	if err != nil {
		log.Fatalf("Error loading HOTEL_ID: %v", err)
	}

	PAYMENT_ID = os.Getenv("PAYMENT_ID")
	PORT, err = strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Error loading PORT: %v", err)
	}

	BOOKING_HOST = os.Getenv("BOOKING_HOST")
	BOOKING_PORT, err = strconv.ParseInt(os.Getenv("BOOKING_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Error loading BOOKING_PORT: %v", err)
	}

	PAYMENT_HOST = os.Getenv("PAYMENT_HOST")
	PAYMENT_PORT, err = strconv.ParseInt(os.Getenv("PAYMENT_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Error loading PAYMENT_PORT: %v", err)
	}

	AUTH_HOST = os.Getenv("AUTH_HOST")
	AUTH_PORT, err = strconv.ParseInt(os.Getenv("AUTH_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Error loading AUTH_PORT: %v", err)
	}

	router := gin.Default()

	bconn, err := grpc.NewClient(
		fmt.Sprintf("%v:%v", BOOKING_HOST, BOOKING_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer bconn.Close()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	aconn, err := grpc.NewClient(
		fmt.Sprintf("%v:%v", AUTH_HOST, AUTH_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer bconn.Close()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	pconn, err := grpc.NewClient(
		fmt.Sprintf("%v:%v", PAYMENT_HOST, PAYMENT_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer bconn.Close()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	bclient := pbooking.NewBookingServiceClient(bconn)
	aclient := pauths.NewAuthServiceClient(aconn)
	pclient := ppayment.NewPaymentServiceClient(pconn)

	router.POST("/api/register", handleRegister(&aclient))
	router.POST("/api/login", handleLogin(&aclient))

	router.GET("/api/book/createBooking", handleCreateBooking(&bclient))
	router.POST("/api/book/getAllBookings", handleGetAllBookings(&bclient))

	router.POST("/api/payment/makePayment", handleMakePayment(&pclient))

	router.Use(cors.Default())
	log.Fatal(router.Run(":8080"))
}

func handleMakePayment(client *ppayment.PaymentServiceClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body struct {
			Email      string  `json:"email"`
			Amount     float64 `json:"amount"`
			CardNumber string  `json:"card_number"`
			ExpiryDate string  `json:"expiry_date"`
			CVV        string  `json:"cvv"`
		}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		req := ppayment.CreatePaymentRequest{
			UserId:   body.Email,
			Amount:   body.Amount,
			Currency: "USD",  // Assuming currency is USD, change if needed
			Method:   "card", // Assuming method is card, change if needed
		}

		response, err := (*client).CreatePayment(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"payment_id": response.PaymentId,
			"message":    response.Message,
		})
	}
}

func handleRegister(client *pauths.AuthServiceClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
			Age      int32  `json:"age"`
			Gender   string `json:"gender"`
		}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		req := pauths.SignupRequest{
			ClientId: CLIENT_ID,
			UserData: map[string]string{
				"email":    body.Email,
				"password": body.Password,
				"name":     body.Name,
				"age":      fmt.Sprintf("%d", body.Age),
				"gender":   body.Gender,
			},
			PrimaryKeyField: "email",
		}

		response, err := (*client).Signup(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": response.Message})
	}
}

func handleLogin(client *pauths.AuthServiceClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		req := pauths.LoginRequest{
			ClientId:        CLIENT_ID,
			PrimaryKeyField: "email",
			PrimaryKeyValue: body.Email,
			Password:        body.Password,
		}

		response, err := (*client).Login(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.UserDetails)
	}
}

func handleCreateBooking(client *pbooking.BookingServiceClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body struct {
			CustomerName  string `json:"customer_name"`
			CustomerEmail string `json:"customer_email"`
			TransactionID int64  `json:"transaction_id"`
			RoomType      int32  `json:"room_type"`
			CheckInDate   string `json:"check_in_date"`
			CheckOutDate  string `json:"check_out_date"`
		}

		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		req := pbooking.GetRoomsRequest{
			HotelId:      HOTEL_ID,
			CheckInDate:  body.CheckInDate,
			CheckOutDate: body.CheckOutDate,
		}
		response, err := (*client).GetRooms(c, &req)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": "No Room Available"})
			return
		}

		var selectedRoom *pbooking.RoomInfo
		for _, room := range response.Rooms {
			if room.RoomType == body.RoomType && room.Available {
				selectedRoom = room
				break
			}
		}

		if selectedRoom == nil {
			c.JSON(
				http.StatusNotFound,
				gin.H{"msg": "No rooms available for the requested type"},
			)
			return
		}

		req2 := pbooking.CreateBookingRequest{
			CustomerName:  body.CustomerName,
			CustomerEmail: body.CustomerEmail,
			CheckInDate:   body.CheckInDate,
			CheckOutDate:  body.CheckOutDate,
			RoomId:        selectedRoom.RoomId,
			TransactionId: body.TransactionID,
		}
		response2, err2 := (*client).CreateBooking(c, &req2)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "xdd"})
			fmt.Println("xdd")
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{"booking_id": response2.BookingId, "room_no": selectedRoom.RoomNumber},
		)
	}
}

func handleGetAllBookings(client *pbooking.BookingServiceClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		customerEmail := c.Query("customer_email")
		req := pbooking.ListCustomerBookingsRequest{
			CustomerEmail: customerEmail,
			HotelId:       HOTEL_ID,
		}
		response, err := (*client).ListCustomerBookings(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.Bookings)
	}
}

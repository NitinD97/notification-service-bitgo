package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"notification-server-bitgo/models"
)

//Create a crypto notification service as an HTTP Rest API server
//Endpoints:
//- Create a Notification (Input parameters: Current Price of Bitcoin, Daily Percentage Change, Trading Volume, etc)
//- Send a notification to email/emails
//- List notifications (Sent, Pending, Failed)
//- Update/Delete notification

type App struct {
	pgConn  *pgxpool.Pool
	queries *models.Queries
	router  *gin.Engine
}

var userid = uuid.New()

func main() {
	fmt.Println("Starting application...")
	dbConn := InitPostgres()
	queries := models.New(dbConn)

	err := queries.TruncateUsers(context.Background())
	if err != nil {
		panic(err)
	}
	u, err := queries.CreateUser(context.Background(), models.CreateUserParams{
		Name:  "Test1",
		Email: "test1@test.com",
	})
	if err != nil {
		panic(err)
	}

	userid = u.ID

	r := gin.Default()

	app := App{
		pgConn:  dbConn,
		queries: queries,
		router:  r,
	}

	app.InitializeRoutes()
	err = r.Run()
	if err != nil {
		panic(err)
	}
}

func (app *App) InitializeRoutes() {
	fmt.Println("Initializing routes...")

	app.router.POST("/notification", app.createNotification)
	app.router.GET("/notifications", app.listNotifications)

	app.router.POST("/notification/send", app.sendNotification)
}

func (app *App) createNotification(c *gin.Context) {
	var input CreateNotificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: check validations

	id := uuid.New()
	currentPrice := pgtype.Numeric{}
	percentChange := pgtype.Numeric{}
	err := currentPrice.Scan(fmt.Sprintf("%.2f", input.CurrentPrice))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = percentChange.Scan(fmt.Sprintf("%.2f", input.PercentChange))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.queries.CreateNotification(c.Request.Context(), models.CreateNotificationParams{
		ID:            id,
		CurrentPrice:  currentPrice,
		PercentChange: percentChange,
		Volume:        int32(input.Volume),
		UserID:        userid,
	})

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": id})
}

func (app *App) listNotifications(c *gin.Context) {
	notifications, err := app.queries.GetNotificationsByUserId(c.Request.Context(), userid)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notifications})
}

func (app *App) sendNotification(c *gin.Context) {
	var input SendNotificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: validate request

	// NOTE: expect the smtp logic here

	tx, err := app.pgConn.Begin(c.Request.Context())
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}
	defer tx.Rollback(c.Request.Context())
	err = app.queries.WithTx(tx).UpdateNotificationStatusById(c.Request.Context(), models.UpdateNotificationStatusByIdParams{
		Status: models.NotificationStatusSENT,
		ID:     input.NotificationId,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	for _, email := range input.Emails {
		err = app.queries.CreateEmailNotification(c.Request.Context(), models.CreateEmailNotificationParams{
			NotificationID: input.NotificationId,
			SentTo:         email,
		})
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
			return
		}
	}

	err = tx.Commit(c.Request.Context())
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input.NotificationId})
}

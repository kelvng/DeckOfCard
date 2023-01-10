package main

import (
	"DeckofCards/module/deck"
	"DeckofCards/module/players"
	"DeckofCards/structs"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	// Enable CORS for all routes
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	// Create a new deck of cards
	r.POST("/new", func(c *gin.Context) {
		deck.NewDeck()
		players.InitPlayers()
		c.JSON(http.StatusOK, fmt.Sprintf("Created a new deck of cards with %d cards.", len(structs.Decks)))
	})

	// Shuffle the deck
	r.POST("/shuffle", func(c *gin.Context) {
		if len(structs.Decks) == 0 {
			deck.NewDeck()
			players.InitPlayers()
			c.JSON(http.StatusOK, gin.H{"message": "\"Created a new deck of cards"})
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		deck.Shuffle()
		c.JSON(http.StatusOK, fmt.Sprintf("Shuffled the deck of cards: %v", structs.Decks))
	})

	// Get the top card of the deck
	r.GET("/top", func(c *gin.Context) {
		if len(structs.Decks) > 0 {
			c.JSON(http.StatusOK, structs.Decks[0])
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	// Draw cards from the top of the deck
	r.DELETE("/drawn", func(c *gin.Context) {
		count := 3

		if len(structs.Decks) >= count {
			for i := 0; i < len(structs.Players); i++ {
				structs.Players[i].Hand = []structs.Card{}
				structs.Players[i].Hand = append(structs.Players[i].Hand, structs.Decks[:count]...)
				structs.Decks = structs.Decks[count:]

				if deck.CalculatePoints(structs.Players[i].Hand) < 10 {
					structs.Players[i].Points = deck.CalculatePoints(structs.Players[i].Hand)
				} else {
					structs.Players[i].Points = deck.CalculatePoints(structs.Players[i].Hand) % 10
				}
			}

			structs.DrawnCards = structs.Decks[:count] // update the drawn cards
			structs.Decks = structs.Decks[count:]      // remove the drawn cards from the deck

			// Determine the winner of the game.
			winner := players.FindWinner(structs.Players)

			// Subtract coins from the losing players.
			players.SubtractCoins(winner, structs.Players)

			c.Status(http.StatusOK)
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not enough cards in the deck. Please shuffle the deck and try again."})
		}
	})

	// Define the "Reveal" route
	r.GET("/reveal", func(c *gin.Context) {
		// Determine the winning and losing players
		winningPlayer, losingPlayer := players.DetermineWinnersAndLosers(structs.Players)

		// Return the result to the client
		c.JSON(http.StatusOK, gin.H{
			"drawnCards": structs.DrawnCards,
			"winner":     winningPlayer.Name,
			"loser":      losingPlayer.Name,
		})
	})

	// Get player information
	r.GET("/player/:id", func(c *gin.Context) {
		id := c.Param("id")
		playerID, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if playerID < 1 || playerID > len(structs.Players) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		player := structs.Players[playerID-1]
		c.JSON(http.StatusOK, player)
	})

	r.GET("/players", func(c *gin.Context) {
		c.JSON(http.StatusOK, structs.Players)
	})

	r.POST("/reset", func(c *gin.Context) {
		deck.NewDeck()
		deck.Shuffle()
		players.InitPlayers()
		c.JSON(http.StatusOK, gin.H{"message": "Game reset"})
	})

	r.Run()
}

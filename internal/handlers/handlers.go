package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nutcase/dao-golang/pkg/client"
)

// TaskRequest represents the request body for creating a task
type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Bounty      string `json:"bounty"` // Amount in tokens
}

// ProofRequest represents the request body for submitting proof
type ProofRequest struct {
	ProofURI string `json:"proof_uri"`
}

// WalletRequest represents the request body for creating a wallet
type WalletRequest struct {
	Name string `json:"name"`
}

// CreateTask handles task creation
func CreateTask(c *client.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req TaskRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: Implement task creation using client
		ctx.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
	}
}

// ListTasks handles listing all tasks
func ListTasks(c *client.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Implement task listing using client
		ctx.JSON(http.StatusOK, gin.H{"tasks": []string{}})
	}
}

// GetTask handles getting a single task by ID
func GetTask(c *client.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		// TODO: Implement get task using client
		ctx.JSON(http.StatusOK, gin.H{"task": id})
	}
}

// ClaimTask handles claiming a task
func ClaimTask(c *client.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		// TODO: Implement task claiming using client
		ctx.JSON(http.StatusOK, gin.H{"message": "Task claimed successfully"})
	}
}

// SubmitProof handles submitting proof for a task
func SubmitProof(c *client.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		var req ProofRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: Implement proof submission using client
		ctx.JSON(http.StatusOK, gin.H{"message": "Proof submitted successfully"})
	}
}

// ApproveTask handles approving a task
func ApproveTask(c *client.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		// TODO: Implement task approval using client
		ctx.JSON(http.StatusOK, gin.H{"message": "Task approved successfully"})
	}
}

// CreateWallet handles wallet creation
func CreateWallet(c *client.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req WalletRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create wallet using client
		address, err := c.CreateWallet(req.Name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"address": address})
	}
}

// GetBalance handles getting wallet balance
func GetBalance(c *client.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		address := ctx.Param("address")
		denom := ctx.DefaultQuery("denom", "stake") // Default to 'stake' token

		// Get balance using client
		balance, err := c.GetBalance(address, denom)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"balance": balance})
	}
}

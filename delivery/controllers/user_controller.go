package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/Afomiat/Loan-Tracker-API/usecases"
    "github.com/Afomiat/Loan-Tracker-API/domain"
    "net/http"
)

type UserController struct {
    UserUsecases usecases.UserUsecases
}

func (uc *UserController) Register(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := uc.UserUsecases.RegisterUser(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) VerifyEmail(c *gin.Context) {
    token := c.Query("token")
    email := c.Query("email")

    err := uc.UserUsecases.VerifyEmail(email, token)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

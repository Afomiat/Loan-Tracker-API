package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/Afomiat/Loan-Tracker-API/usecases"
    "net/http"
)

type AdminController struct {
    AdminUsecases *usecases.AdminUsecases
}

func (ac *AdminController) ViewAllUsers(c *gin.Context) {
    users, err := ac.AdminUsecases.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}

func (ac *AdminController) DeleteUser(c *gin.Context) {
    id := c.Param("id")

    err := ac.AdminUsecases.DeleteUser(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

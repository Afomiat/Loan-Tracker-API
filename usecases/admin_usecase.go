package usecases

import (
    "github.com/Afomiat/Loan-Tracker-API/domain"
    "github.com/Afomiat/Loan-Tracker-API/repositories"
)

type AdminUsecases struct {
    UserRepo *repositories.UserRepository
}

func (au *AdminUsecases) GetAllUsers() ([]domain.User, error) {
    // Implement the logic to get all users
    return nil, nil
}

func (au *AdminUsecases) DeleteUser(id string) error {
    // Implement the logic to delete a user by ID
    return nil
}

package usecases

import (
    "github.com/Afomiat/Loan-Tracker-API/domain"
    "github.com/Afomiat/Loan-Tracker-API/repositories"
    "github.com/Afomiat/Loan-Tracker-API/infrastructure"
    
    
)

type UserUsecases struct {
    UserRepo      *repositories.UserRepository
    EmailService  *infrastructure.EmailService
    TokenService  *infrastructure.TokenService
}

func (uc *UserUsecases) RegisterUser(user *domain.User) error {
    existingUser, _ := uc.UserRepo.FindByEmail(user.Email)
    if existingUser != nil {
        return domain.ErrEmailAlreadyUsed
    }

    err := uc.UserRepo.Create(user)
    if err != nil {
        return err
    }

    token, err := uc.TokenService.GenerateAccessToken(user.ID)
    if err != nil {
        return err
    }

    err = uc.EmailService.SendVerificationEmail(user.Email, token)
    return err
}

func (uc *UserUsecases) VerifyEmail(email, token string) error {
    user, err := uc.UserRepo.FindByEmail(email)
    if err != nil {
        return err
    }
    if user == nil {
        return domain.ErrUserNotFound
    }

    err = uc.UserRepo.VerifyUser(email, token)
    return err
}

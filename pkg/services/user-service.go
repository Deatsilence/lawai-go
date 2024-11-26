package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	businesslogic "github.com/Deatsilence/lawai-go/pkg/business-logic"
	models "github.com/Deatsilence/lawai-go/pkg/models/mongo"
	"github.com/Deatsilence/lawai-go/pkg/repositories"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo *repositories.UserRepository
	bl   *businesslogic.UserBL
}

var validateUser = validator.New()

func NewUserService(repo *repositories.UserRepository, bl *businesslogic.UserBL) *UserService {
	return &UserService{repo: repo, bl: bl}
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (primitive.ObjectID, int, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	if validationErr := validateUser.Struct(user); validationErr != nil {
		return primitive.NilObjectID, http.StatusBadRequest, validationErr
	}

	count, err := s.repo.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return primitive.NilObjectID, http.StatusInternalServerError, err
	}

	if count > 0 {
		isVerified, err := s.DeleteUnverified(ctx, *user.Email)
		if err != nil {
			return primitive.NilObjectID, http.StatusInternalServerError, err
		}
		if isVerified {
			return primitive.NilObjectID, http.StatusConflict, fmt.Errorf("email already exists")
		}
	}

	password, err := s.bl.HashPassword(*user.Password)
	if err != nil {
		return primitive.NilObjectID, http.StatusInternalServerError, err
	}

	user.Password = &password

	insertedID, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return primitive.NilObjectID, http.StatusInternalServerError, err
	}
	return insertedID, http.StatusCreated, nil
}

func (s *UserService) DeleteUnverified(ctx context.Context, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	var user models.User
	foundUser, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	user = *foundUser

	if !user.IsVerified {
		err := s.repo.DeleteUser(ctx, user.ID)
		if err != nil {
			return user.IsVerified, err
		}
	}
	return user.IsVerified, nil
}

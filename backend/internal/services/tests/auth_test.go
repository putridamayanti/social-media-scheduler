package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"social-media-scheduler/internal/dtos"
	"social-media-scheduler/internal/models"
	"social-media-scheduler/internal/services"
	"testing"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepo) Create(ctx context.Context, u *models.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

type MockAuthRepo struct {
	mock.Mock
}

func (m *MockAuthRepo) CreateSession(ctx context.Context, userId uuid.UUID) (*models.Session, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(*models.Session), args.Error(1)
}

func (m *MockAuthRepo) RemoveSession(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestAuthService_Login_Success(t *testing.T) {
	userRepo := new(MockUserRepo)
	authRepo := new(MockAuthRepo)

	service := services.NewAuthService(authRepo, userRepo)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	userId, _ := uuid.Parse("605f2149-89ea-413f-97c9-d82273cbb6a1")
	user := &models.User{
		ID:       userId,
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	sessionId, _ := uuid.Parse("00276e3c-1cc0-4184-8584-a1bb6d23b0c8")
	session := &models.Session{
		ID:     sessionId,
		UserID: userId,
	}

	request := dtos.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	userRepo.On("GetByEmail", mock.Anything, request.Email).Return(user, nil)

	authRepo.On("CreateSession", mock.Anything, user.ID).Return(session, nil)

	result, err := service.Login(context.Background(), request)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.NoError(t, err)
	assert.Equal(t, session, result)
}

func TestAuthService_Login_InvalidEmailOrPassword(t *testing.T) {
	userRepo := new(MockUserRepo)
	authRepo := new(MockAuthRepo)

	service := services.NewAuthService(authRepo, userRepo)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correct"), bcrypt.DefaultCost)

	userId, _ := uuid.Parse("605f2149-89ea-413f-97c9-d82273cbb6a1")
	user := &models.User{
		ID:       userId,
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	req := dtos.LoginRequest{
		Email:    "test@mail.com",
		Password: "wrong",
	}

	userRepo.
		On("GetByEmail", mock.Anything, req.Email).
		Return(user, nil)

	result, err := service.Login(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, result)

	authRepo.AssertNotCalled(t, "CreateSession")
}

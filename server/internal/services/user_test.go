package services

import (
	"errors"
	"reflect"
	"server/internal/models"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

var employee2 = models.User{
	Username:  "khang",
	Password:  "khangpham",
	Role:      1,
	Telephone: "0123456789",
	DName:     "Finance",
	Login:     false,
}

var asset2 = models.Asset{
	Id:          1,
	Name:        "table",
	Type:        "funiture",
	Status:      "working",
	StatusNote:  "",
	Description: "This is a long table",
	UserId:      1,
	CreatedAt:   time.Date(2023, time.January, 2, 15, 4, 5, 0, time.UTC),
	UpdatedAt:   time.Date(2023, time.January, 2, 15, 4, 5, 0, time.UTC),
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		id   uint32
		user *models.User
	}
	tests := []struct {
		name     string
		args     args
		mockRepo func(ctrl *gomock.Controller) models.IUserRepository
		wantErr  error
	}{
		{
			name: "Should return error when updating non existing user",
			args: args{
				id:   2,
				user: &employee2,
			},
			mockRepo: func(ctrl *gomock.Controller) models.IUserRepository {
				m := models.NewMockIUserRepository(ctrl)
				m.EXPECT().GetUserById(uint32(2)).Return(nil, gorm.ErrRecordNotFound)
				return m
			},
			wantErr: gorm.ErrRecordNotFound,
		},
		{
			name: "Should update success when user exists",
			args: args{
				id:   2,
				user: &employee2,
			},
			mockRepo: func(ctrl *gomock.Controller) models.IUserRepository {
				m := models.NewMockIUserRepository(ctrl)
				m.EXPECT().GetUserById(uint32(2)).Return(&models.User{
					ID:        2,
					Username:  employee2.Username,
					Password:  employee2.Password,
					Role:      employee2.Role,
					Telephone: employee2.Telephone,
					DName:     employee2.DName,
					Login:     employee2.Login,
				}, nil)
				m.EXPECT().UpdateUser(&models.User{
					ID:        2,
					Username:  employee2.Username,
					Password:  employee2.Password,
					Role:      employee2.Role,
					Telephone: employee2.Telephone,
					DName:     employee2.DName,
					Login:     employee2.Login,
				}).Return(nil)
				return m
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userService := &UserService{
				Repository: tt.mockRepo(ctrl),
			}

			err := userService.UpdateUser(tt.args.id, tt.args.user)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_GetUserById(t *testing.T) {
	type args struct {
		id uint32
	}
	tests := []struct {
		name     string
		args     args
		mockRepo func(ctrl *gomock.Controller) models.IUserRepository
		want     *models.User
		wantErr  error
	}{
		{
			name: "Should return error when user not exist",
			args: args{
				id: 2,
			},
			mockRepo: func(ctrl *gomock.Controller) models.IUserRepository {
				m := models.NewMockIUserRepository(ctrl)
				m.EXPECT().GetUserById(gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
				return m
			},
			wantErr: gorm.ErrRecordNotFound,
			want:    nil,
		},
		{
			name: "Should return user info when user exists",
			args: args{
				id: 1,
			},
			mockRepo: func(ctrl *gomock.Controller) models.IUserRepository {
				m := models.NewMockIUserRepository(ctrl)
				m.EXPECT().GetUserById(gomock.Any()).Return(&employee2, nil)
				return m
			},
			wantErr: nil,
			want:    &employee2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userService := &UserService{
				Repository: tt.mockRepo(ctrl),
			}

			got, err := userService.GetUserById(tt.args.id)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got user info: %v, wantUser = %v", got, tt.want)
			}
		})
	}
}

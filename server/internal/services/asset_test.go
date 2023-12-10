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

var employee1 = models.User{
	Username:  "mai",
	Password:  "123",
	Role:      1,
	Telephone: "",
	DName:     "Finance",
	Login:     false,
}

var asset1 = models.Asset{
	Id:          1,
	Name:        "table",
	Type:        "funiture",
	Status:      "working",
	StatusNote:  "",
	Description: "",
	UserId:      1,
	CreatedAt:   time.Date(2023, time.January, 2, 15, 4, 5, 0, time.UTC),
	UpdatedAt:   time.Date(2023, time.January, 2, 15, 4, 5, 0, time.UTC),
}

func TestUpdateAsset(t *testing.T) {
	type args struct {
		id    uint32
		asset *models.Asset
	}
	tests := []struct {
		name     string
		args     args
		mockRepo func(ctrl *gomock.Controller) models.IAssetRepository
		wantErr  error
	}{
		{
			name: "Should return error when updating non existing asset",
			args: args{
				id: 2,
				asset: &models.Asset{
					Status:      "not working",
					StatusNote:  "2 days",
					Description: "need fix",
				},
			},
			mockRepo: func(ctrl *gomock.Controller) models.IAssetRepository {
				m := models.NewMockIAssetRepository(ctrl)
				m.EXPECT().GetAssetById(gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
				return m
			},
			wantErr: gorm.ErrRecordNotFound,
		},
		{
			name: "Should update success when asset exists",
			args: args{
				id: 1,
				asset: &models.Asset{
					Status:      "not working",
					StatusNote:  "2 days",
					Description: "need fix",
				},
			},
			mockRepo: func(ctrl *gomock.Controller) models.IAssetRepository {
				m := models.NewMockIAssetRepository(ctrl)
				m.EXPECT().GetAssetById(gomock.Any()).Return(&asset1, nil)
				m.EXPECT().UpdateAsset(gomock.Any()).Return(nil)
				return m
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			assetService := &AssetService{
				Repository: tt.mockRepo(ctrl),
			}

			err := assetService.UpdateAsset(tt.args.id, tt.args.asset)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UpdateAsset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAssetService_GetAssetById(t *testing.T) {
	type args struct {
		id uint32
	}
	tests := []struct {
		name     string
		args     args
		mockRepo func(ctrl *gomock.Controller) models.IAssetRepository
		want     *models.Asset
		wantErr  error
	}{
		{
			name: "Should return error when asset not exist",
			args: args{
				id: 2,
			},
			mockRepo: func(ctrl *gomock.Controller) models.IAssetRepository {
				m := models.NewMockIAssetRepository(ctrl)
				m.EXPECT().GetAssetById(gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
				return m
			},
			wantErr: gorm.ErrRecordNotFound,
			want:    nil,
		},
		{
			name: "Should return asset info when asset exist",
			args: args{
				id: 1,
			},
			mockRepo: func(ctrl *gomock.Controller) models.IAssetRepository {
				m := models.NewMockIAssetRepository(ctrl)
				m.EXPECT().GetAssetById(gomock.Any()).Return(&asset1, nil)
				return m
			},
			wantErr: nil,
			want:    &asset1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			assetService := &AssetService{
				Repository: tt.mockRepo(ctrl),
			}

			got, err := assetService.GetAssetById(tt.args.id)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetAssetById() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got asset info: %v, wantAsset = %v", got, tt.want)
			}
		})
	}
}

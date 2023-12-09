package services

import (
	"errors"
	"reflect"
	"server/internal/db"
	"server/internal/db/migrations"
	"server/internal/models"
	"testing"
	"time"

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
		name    string
		args    args
		wantErr error
		want    models.Asset
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
			wantErr: gorm.ErrRecordNotFound,
			want: models.Asset{
				Id:          asset1.Id,
				Name:        asset1.Name,
				Type:        asset1.Type,
				Status:      asset1.Status,
				StatusNote:  asset1.StatusNote,
				Description: asset1.Description,
				UserId:      asset1.UserId,
				CreatedAt:   asset1.CreatedAt,
				UpdatedAt:   asset1.UpdatedAt,
			},
		},
		{
			name: "Should update more than 1 field in asset when asset exists",
			args: args{
				id: 1,
				asset: &models.Asset{
					Status:      "not working",
					StatusNote:  "2 days",
					Description: "need fix",
				},
			},
			wantErr: nil,
			want: models.Asset{
				Id:          asset1.Id,
				Name:        asset1.Name,
				Type:        asset1.Type,
				Status:      "not working",
				StatusNote:  "2 days",
				Description: "need fix",
				UserId:      asset1.UserId,
				CreatedAt:   asset1.CreatedAt,
				UpdatedAt:   asset1.UpdatedAt,
			},
		},
		{
			name: "Should update 1 field in asset when asset exists",
			args: args{
				id: 1,
				asset: &models.Asset{
					UserId: 2,
				},
			},
			wantErr: nil,
			want: models.Asset{
				Id:          asset1.Id,
				Name:        asset1.Name,
				Type:        asset1.Type,
				Status:      asset1.Status,
				StatusNote:  asset1.StatusNote,
				Description: asset1.Description,
				UserId:      2,
				CreatedAt:   asset1.CreatedAt,
				UpdatedAt:   asset1.UpdatedAt,
			},
		},
	}

	db.ConnectDB()

	assetService := NewAssetService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			migrations.MigrationDown(db.DB)
			migrations.MigrateUp(db.DB)
			prepareTestUpdateAssetData()

			err := assetService.UpdateAsset(tt.args.id, tt.args.asset)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UpdateAsset() error = %v, wantErr %v", err, tt.wantErr)
			}

			var assetAfterUpdate models.Asset
			_ = db.DB.First(&assetAfterUpdate, asset1.Id).Error
			if !reflect.DeepEqual(models.Asset{
				Id:          assetAfterUpdate.Id,
				Name:        assetAfterUpdate.Name,
				Type:        assetAfterUpdate.Type,
				Status:      assetAfterUpdate.Status,
				StatusNote:  assetAfterUpdate.StatusNote,
				Description: assetAfterUpdate.Description,
				UserId:      assetAfterUpdate.UserId,
				CreatedAt:   tt.want.CreatedAt,
				UpdatedAt:   tt.want.UpdatedAt,
			}, tt.want) {
				t.Errorf("Asset after update: %v, wantAsset = %v", assetAfterUpdate, tt.want)
			}
		})
	}
}

func prepareTestUpdateAssetData() {
	addUser(&employee1)
	addAsset(&asset1)
}

func addUser(user *models.User) {
	db.DB.Model(&models.User{}).Create(&user)
}

func addAsset(asset *models.Asset) {
	db.DB.Model(&models.Asset{}).Create(&asset)
}

func prepareTestGetAssetById() {
	addUser(&employee1)
	addAsset(&asset1)
}

func TestAssetService_GetAssetById(t *testing.T) {
	type args struct {
		id uint32
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Asset
		wantErr error
	}{
		{
			name: "Should return error when asset not exist",
			args: args{
				id: 2,
			},
			wantErr: gorm.ErrRecordNotFound,
			want:    nil,
		},
		{
			name: "Should return asset info when asset exist",
			args: args{
				id: 1,
			},
			wantErr: nil,
			want:    &asset1,
		},
	}

	db.ConnectDB()

	assetService := NewAssetService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			migrations.MigrationDown(db.DB)
			migrations.MigrateUp(db.DB)
			prepareTestGetAssetById()

			got, err := assetService.GetAssetById(tt.args.id)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetAssetById() error = %v, wantErr %v", err, tt.wantErr)
			}
			switch tt.name {
			case "Should return asset info when asset exist":
				if !reflect.DeepEqual(models.Asset{
					Id:          got.Id,
					Name:        got.Name,
					Type:        got.Type,
					Status:      got.Status,
					StatusNote:  got.StatusNote,
					Description: got.Description,
					UserId:      got.UserId,
				}, models.Asset{
					Id:          tt.want.Id,
					Name:        tt.want.Name,
					Type:        tt.want.Type,
					Status:      tt.want.Status,
					StatusNote:  tt.want.StatusNote,
					Description: tt.want.Description,
					UserId:      tt.want.UserId,
				}) {
					t.Errorf("Got asset info: %v, wantAsset = %v", got, tt.want)
				}
			case "Should return error when asset not exist":
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Got asset info: %v, wantAsset = %v", got, tt.want)
				}
			}
		})
	}
}

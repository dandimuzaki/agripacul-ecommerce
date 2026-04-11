package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/pkg/utils"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SeedProvinces(db *gorm.DB, log *zap.Logger) error {
	type ProvinceAPI struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	var res utils.APIResponse[ProvinceAPI]
	req, err := http.NewRequest(
		"GET",
		"https://wilayah.id/api/provinces.json",
		nil,
	)
	data, err := utils.FetchJSON(req, res, log)
	if err != nil {
		log.Error("Error fetch provinces", zap.Error(err))
		return err
	}

	for _, p := range data {
		err :=db.FirstOrCreate(
			&entity.Province{},
			entity.Province{
				Code: p.Code,
				Name: p.Name,
			},
		).Error
		if err != nil {
			log.Error("Error create provinces", zap.Error(err))
			return err
		}
	}

	return nil
}

func SeedRegencies(db *gorm.DB, log *zap.Logger) error {
	var provinces []entity.Province
	err := db.Find(&provinces).Error
	if err != nil {
		log.Error("Error get provinces", zap.Error(err))
		return err
	}

	for _, prov := range provinces {
		type RegencyAPI struct {
			Code string `json:"code"`
			Name string `json:"name"`
		}

		var res utils.APIResponse[RegencyAPI]
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("https://wilayah.id/api/regencies/%s.json", prov.Code),
			nil,
		)
		data, err := utils.FetchJSON(req, res, log)
		if err != nil {
			log.Error("Error fetch regencies", zap.Error(err))
			return err
		}

		for _, c := range data {
			cityType := "kabupaten"
			if strings.Contains(strings.ToLower(c.Name), "kota") {
				cityType = "kota"
			}

			err := db.FirstOrCreate(&entity.Regency{}, entity.Regency{
				Code:       c.Code,
				ProvinceID: prov.ID,
				Name:       c.Name,
				Type:       cityType,
			}).Error
			if err != nil {
				log.Error("Error create regencies", zap.Error(err))
				return err
			}
		}
	}

	return nil
}

func SeedDistricts(db *gorm.DB, log *zap.Logger) error {
	var regencies []entity.Regency
	err := db.Find(&regencies).Error
	if err != nil {
		log.Error("Error get regencies", zap.Error(err))
		return err
	}

	for _, regency := range regencies {
		type DistrictAPI struct {
			Code   string `json:"code"`
			Name string `json:"name"`
		}

		var res utils.APIResponse[DistrictAPI]
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("https://wilayah.id/api/districts/%s.json", regency.Code),
			nil,
		)
		data, err := utils.FetchJSON(req, res, log)
		if err != nil {
			log.Error("Error fetch districts", zap.Error(err))
			return err
		}

		for _, d := range data {
			err := db.FirstOrCreate(&entity.District{}, entity.District{
				Code:   d.Code,
				RegencyID: regency.ID,
				Name:   d.Name,
			}).Error
			if err != nil {
				log.Error("Error create districts", zap.Error(err))
				return err
			}
		}
	}

	return nil
}

func SeedSubdistricts(db *gorm.DB, log *zap.Logger) error {
	var districts []entity.District
	err := db.Find(&districts).Error
	if err != nil {
		log.Error("Error get districts", zap.Error(err))
		return err
	}

	for _, dist := range districts {
		type SubdistrictAPI struct {
			Code   string `json:"code"`
			Name string `json:"name"`
		}

		var res utils.APIResponse[SubdistrictAPI]
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("https://wilayah.id/api/villages/%s.json", dist.Code),
			nil,
		)
		data, err := utils.FetchJSON(req, res, log)
		if err != nil {
			log.Error("Error fetch subdistricts", zap.Error(err))
			return err
		}

		for _, s := range data {
			err = db.FirstOrCreate(&entity.Subdistrict{}, entity.Subdistrict{
				Code:       s.Code,
				DistrictID: dist.ID,
				Name:       s.Name,
			}).Error
			if err != nil {
				log.Error("Error create subdistricts", zap.Error(err))
				return err
			}
		}
	}

	return nil
}


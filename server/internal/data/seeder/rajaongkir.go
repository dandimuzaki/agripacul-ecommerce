package seeder

import (
	"debian-ecommerce/internal/data/entity"
	"debian-ecommerce/pkg/utils"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RajaOngkirProvince struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RajaOngkirCity struct {
	ID       uint   `json:"id"`
	Name string `json:"name"`
	ZipCode     string `json:"zip_code"`
}

type RajaOngkirDistrict struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	ZipCode string `json:"zip_code"`
}

func FetchRajaOngkirProvinces(apiKey string, log *zap.Logger) ([]RajaOngkirProvince, error) {
	req, err := http.NewRequest(
		"GET",
		"https://rajaongkir.komerce.id/api/v1/destination/province",
		nil,
	)
	req.Header.Set("key", apiKey)
	if err != nil {
		log.Error("Error request fetch raja ongkir provinces", zap.Error(err))
		return nil, err
	}
	var res utils.APIResponse[RajaOngkirProvince]
	data, err := utils.FetchJSON(req, res, log)
	if err != nil {
		log.Error("Error get response fetch raja ongkir provinces", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func normalizeProvinceName(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "dki", "")
	s = strings.ReplaceAll(s, "daerah istimewa", "")
	s = regexp.MustCompile(`[^a-z ]`).ReplaceAllString(s, "")
	return strings.TrimSpace(s)
}

func SeedRajaOngkirRegencies(
	db *gorm.DB,
	apiKey string,
	log *zap.Logger,
) error {

	roProvinces, err := FetchRajaOngkirProvinces(apiKey, log)
	if err != nil {
		log.Error("Error request fetch raja ongkir provinces", zap.Error(err))
		return err
	}

	fmt.Println("roProvinces", roProvinces)

	var provinces []entity.Province
	err = db.Find(&provinces).Error
	if err != nil {
		log.Error("Error get provinces", zap.Error(err))
		return err
	}

	fmt.Println("provinces", provinces)

	provinceMap := make(map[string]uint)
	for _, rp := range roProvinces {
		provinceMap[normalizeProvinceName(rp.Name)] = rp.ID
	}

	fmt.Println("provinceMap", provinceMap)

	for _, p := range provinces {
		err := db.Model(&entity.Province{}).
			Where("LOWER(name) LIKE ?", "%"+normalizeProvinceName(p.Name)+"%").
			Where("raja_ongkir_id IS NULL").
			Update("raja_ongkir_id", p.ID).
			Error

		if err != nil {
			log.Error("Failed update province", zap.Error(err))
		}

		roProvinceID, ok := provinceMap[normalizeProvinceName(p.Name)]
		if !ok {
			log.Warn("Province not matched", zap.String("province", normalizeProvinceName(p.Name)))
			log.Info("map", zap.Any("map", provinceMap))
			continue
		}

		cities, err := FetchRajaOngkirCitiesByProvince(apiKey, roProvinceID, log)
		if err != nil {
			log.Error("Error fetch raja ongkir cities", zap.Error(err))
			return err
		}

		for _, c := range cities {
			matchAndUpdateRegency(db, c, log)
		}
	}

	return nil
}

func FetchRajaOngkirCitiesByProvince(
	apiKey string,
	provinceRajaOngkirID uint,
	log *zap.Logger,
) ([]RajaOngkirCity, error) {

	url := fmt.Sprintf(
		"https://rajaongkir.komerce.id/api/v1/destination/city/%d",
		provinceRajaOngkirID,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("Error request fetch raja ongkir cities", zap.Error(err))
		return nil, err
	}
	req.Header.Set("key", apiKey)

	var res utils.APIResponse[RajaOngkirCity]
	data, err := utils.FetchJSON(req, res, log)
	if err != nil {
		log.Error("Error get response fetch raja ongkir cities", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func normalizeCityName(s string) string {
	s = strings.ToLower(s)

	replacements := []string{
		"kota", "",
		"kabupaten", "",
		"kab.", "",
	}

	for i := 0; i < len(replacements); i += 2 {
		s = strings.ReplaceAll(s, replacements[i], replacements[i+1])
	}

	s = regexp.MustCompile(`[^a-z0-9 ]`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")

	return strings.TrimSpace(s)
}

func FetchRajaOngkirDistrictByCity(
	apiKey string,
	cityRajaOngkirID uint,
	log *zap.Logger,
) ([]RajaOngkirDistrict, error) {

	url := fmt.Sprintf(
		"https://rajaongkir.komerce.id/api/v1/destination/district/%d",
		cityRajaOngkirID,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("Error request fetch raja ongkir districts", zap.Error(err))
		return nil, err
	}

	req.Header.Set("key", apiKey)

	var res utils.APIResponse[RajaOngkirDistrict]
	data, err := utils.FetchJSON(req, res, log)
	if err != nil {
		log.Error("Error get response fetch raja ongkir districts", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func normalizeDistrictName(s string) string {
	s = strings.ToLower(s)

	replacements := []string{
		"kecamatan", "",
		"kec.", "",
		"keb.", "",
		"administrasi", "",
	}

	for i := 0; i < len(replacements); i += 2 {
		s = strings.ReplaceAll(s, replacements[i], replacements[i+1])
	}

	s = regexp.MustCompile(`[^a-z0-9 ]`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")

	return strings.TrimSpace(s)
}

func SeedRajaOngkirDistrictByCity(
	db *gorm.DB,
	regency entity.Regency,
	apiKey string,
	log *zap.Logger,
) error {

	// 1. Guard: regency must already have RajaOngkir city ID
	if regency.RajaOngkirID == nil {
		log.Warn(
			"Regency has no RajaOngkir ID, skipped",
			zap.String("regency", regency.Name),
		)
		return nil
	}

	// 2. Fetch RajaOngkir districts by city
	roDistricts, err := FetchRajaOngkirDistrictByCity(
		apiKey,
		*regency.RajaOngkirID,
		log,
	)
	if err != nil {
		return err
	}

	// 3. Normalize RajaOngkir districts into list (NOT map)
	type roEntry struct {
		ID   uint
		Name string
	}

	roList := make([]roEntry, 0, len(roDistricts))
	for _, rd := range roDistricts {
		roList = append(roList, roEntry{
			ID:   rd.ID,
			Name: normalizeDistrictName(rd.Name),
		})
	}

	// 4. Load local districts for this regency
	var districts []entity.District
	if err := db.
		Where("regency_id = ?", regency.ID).
		Find(&districts).
		Error; err != nil {
		return err
	}

	// 5. Match + update
	for _, d := range districts {

		// idempotent: skip already updated rows
		if d.RajaOngkirID != nil {
			continue
		}

		local := normalizeDistrictName(d.Name)

		var matchedID *uint

		for _, ro := range roList {
			// fuzzy bidirectional match
			if strings.Contains(ro.Name, local) ||
				strings.Contains(local, ro.Name) {
				matchedID = &ro.ID
				break
			}
		}

		if matchedID != nil {
			fmt.Println("matchedID", *matchedID)
			if err := db.Model(&entity.District{}).
				Where("id = ?", d.ID).
				Update("raja_ongkir_id", *matchedID).
				Error; err != nil {

				log.Error(
					"Failed updating district RajaOngkir ID",
					zap.String("district", d.Name),
					zap.Error(err),
				)
				return err
			}
		} else {
			fmt.Println("matchedID", matchedID)
			log.Warn(
				"District not matched",
				zap.String("regency", regency.Name),
				zap.String("district", d.Name),
			)
		}
	}

	return nil
}

func SeedRajaOngkirDistrictsWithLimit(
	db *gorm.DB,
	apiKey string,
	log *zap.Logger,
	maxRequests int,
) error {

	var regencies []entity.Regency
	err := db.
		Where("raja_ongkir_id IS NOT NULL").
		Find(&regencies).Error
	if err != nil {
		log.Error("Error get regencies", zap.Error(err))
		return err
	}

	requests := 0

	for _, regency := range regencies {
		if requests >= maxRequests {
			log.Info("API limit reached, stopping")
			break
		}

		// skip if districts already seeded
		var count int64
		err = db.Model(&entity.District{}).
			Where("regency_id = ?", regency.ID).
			Where("raja_ongkir_id IS NOT NULL").
			Count(&count).Error

		if err != nil {
			log.Error("Error get districts count", zap.Error(err))
			return err
		}

		fmt.Println(count)
		// if count > 0 {
		// 	continue
		// }

		if err := SeedRajaOngkirDistrictByCity(
			db,
			regency,
			apiKey,
			log,
		); err != nil {
			log.Error("Error update district", zap.Error(err))
			return err
		}

		requests++
	}

	return nil
}

func matchAndUpdateRegency(
	db *gorm.DB,
	roCity RajaOngkirCity,
	log *zap.Logger,
) {
	normalizedRO := normalizeCityName(roCity.Name)

	var regency entity.Regency
	err := db.
		Where("LOWER(name) LIKE ?", "%"+normalizedRO+"%").
		First(&regency).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn(
				"Regency not matched",
				zap.String("rajaongkir_city", roCity.Name),
			)
			return
		}

		log.Error(
			"Failed querying regency",
			zap.Error(err),
		)
		return
	}

	// Update only if empty (idempotent seeder)
	if regency.RajaOngkirID != nil {
		return
	}

	if err := db.Model(&entity.Regency{}).
		Where("id = ?", regency.ID).
		Update("raja_ongkir_id", roCity.ID).
		Error; err != nil {

		log.Error(
			"Failed updating regency RajaOngkir ID",
			zap.String("regency", regency.Name),
			zap.Error(err),
		)
		return
	}

	log.Info(
		"Regency matched",
		zap.String("regency", regency.Name),
		zap.Uint("raja_ongkir_id", roCity.ID),
	)
}

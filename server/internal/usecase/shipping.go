package usecase

import (
	"context"
	"debian-ecommerce/internal/dto/request"
	"debian-ecommerce/internal/dto/response"
	"debian-ecommerce/pkg/utils"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func (s *checkoutService) FetchShippingOptions(ctx context.Context, req request.ShippingOptionsRequest) ([]response.ShippingOption, error) {
	apiKey := s.config.RajaOngkirConfig.APIKey
	
	// Get user id from context
	userID := ctx.Value("user_id").(uint)

	// Get customer id
	customer, err := s.repo.CustomerRepo.FindCustomerByUserID(ctx, userID)
	if err != nil {
		s.log.Error("Error find customer by user id", zap.Error(err), zap.Uint("user_id", userID))
		return nil, err
	}

	// Get ordered items (items selected in the cart)
	items, err := s.repo.CartRepo.GetSelectedCartItems(ctx, customer.ID)
		
	// Get company shipping origin address
	origin, err := s.repo.CompanyRepo.GetShippingOriginAddress(ctx)
	if err != nil {
		s.log.Error("Failed to get shipping origin address", zap.Error(err))
	}

	// Calculate total weight
	var totalWeight float64
	for _, i := range items {
		sku, err := s.repo.SKURepo.GetByID(ctx, i.SKUID)
		if err != nil {
			s.log.Error("Failed to get SKU", zap.Error(err))
			return nil, err
		}
		totalWeight += sku.Weight * float64(i.Quantity)
	}

	// Get customer address
	destination, err := s.repo.AddressRepo.FindByID(*req.ShippingAddressID)
	if err != nil {
		s.log.Error("Failed to get customer address", zap.Error(err))
		return nil, err
	}

	// Use mock raja ongkir id if nil
	var originID uint
	if origin != nil && origin.District.RajaOngkirID != nil {
		originID = *origin.District.RajaOngkirID
	} else {
		originID = 1338
	}

	// Use mock raja ongkir id if nil
	var destinationID uint
	if destination.District.RajaOngkirID != nil {
		destinationID = *destination.District.RajaOngkirID
	} else {
		destinationID = 487
	}

	form := url.Values{}
	form.Set("origin", strconv.FormatUint(uint64(originID), 10))
	form.Set("destination", strconv.FormatUint(uint64(destinationID), 10))
	form.Set("weight", strconv.Itoa(int(totalWeight)))
	form.Set("courier", "jne:sicepat:ide:sap:jnt:ninja:tiki:lion:anteraja:pos:ncs:rex:rpx:sentral:star:wahana:dse")
	form.Set("price", "lowest")
	
	reqRO, err := http.NewRequest(
		"POST",
		"https://rajaongkir.komerce.id/api/v1/calculate/district/domestic-cost",
		strings.NewReader(form.Encode()),
	)
	reqRO.Header.Set("key", apiKey)
	reqRO.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var options []response.ShippingOption

	// Use mock shipping if cannot fetch API
	shippings, err := LoadShipping("./internal/mock/mock_shipping.json")
	if err != nil {
		s.log.Error("Error load mock shippings", zap.Error(err))
		return nil, err
	}

	var res utils.APIResponse[response.ShippingOption]
	result, err := utils.FetchJSON(reqRO, res, s.log)
	if err != nil {
		s.log.Error("Error decode shipping options", zap.Error(err))
		options = shippings
	} else {
		options = result
	}
	
	return options, nil
}

func LoadShipping(path string) ([]response.ShippingOption, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var shippings []response.ShippingOption
	if err := json.NewDecoder(file).Decode(&shippings); err != nil {
		return nil, err
	}

	return shippings, nil
}
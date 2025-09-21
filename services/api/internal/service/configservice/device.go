package configservice

import (
	"context"
	"github.com/atareversei/quardian/internal/capture"
	"github.com/atareversei/quardian/services/api/internal/dto"
	"github.com/atareversei/quardian/services/api/internal/dto/configdto"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
	"math"
)

func (s *Service) ListDevices(ctx context.Context, req configdto.ListDevicesRequest) (configdto.ListDevicesResponse, error) {
	// TODO: add database connection later
	const op = "configservice.ListDevices"
	rawDevices, err := capture.GetAllDevices()
	if err != nil {
		return configdto.ListDevicesResponse{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithError(err)
	}

	lastPage := int(math.Ceil(float64(len(rawDevices)) / float64(req.PerPage)))
	if req.Page > lastPage {
		return configdto.ListDevicesResponse{}, richerror.New(op).WithKind(richerror.KindNotFound)
	}

	res := configdto.ListDevicesResponse{
		Meta: dto.ListMeta{
			CurrentPage:        req.Page,
			PerPage:            req.PerPage,
			LastPage:           lastPage,
			TotalNumberOfItems: len(rawDevices),
		},
		List: make([]configdto.ListDevicesResponseItem, 0),
	}

	paginatedRawDevices := make([]capture.Device, 0)
	if lastPage == req.Page {
		paginatedRawDevices = rawDevices[(req.Page-1)*req.PerPage:]
	} else {
		paginatedRawDevices = rawDevices[(req.Page-1)*req.PerPage : (req.Page)*req.PerPage]
	}

	for _, d := range paginatedRawDevices {
		dev := configdto.ListDevicesResponseItem{
			Name:        d.Name,
			Description: d.Description,
			IP:          nil,
			Status:      "Active",
		}

		res.List = append(res.List, dev)
	}

	return res, nil
}

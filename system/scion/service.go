package scion

import (
	"context"
	"errors"

	"github.com/cortezaproject/corteza-server/system/scion/types"
	"go.uber.org/zap"
)

type (
	service struct {
		log *zap.Logger

		drivers  types.DriverSet
		sessions types.SessionSet
	}
)

// NewService initializes and creates a new scion service
//
// @todo options?
func NewService(logger *zap.Logger) *service {
	return &service{
		drivers:  make([]types.Driver, 0, 10),
		sessions: make(types.SessionSet, 0, 10),
	}
}

// RegisterDriver registers a new driver for the service to use
func (svc *service) RegisterDriver(d types.Driver) {
	svc.drivers = append(svc.drivers, d)
}

func (svc *service) NewSession(ctx context.Context) (*types.Session, error) {
	return nil, nil
}

// @todo all of those session management bits to bind parameters, templates, ...

func (svc *service) Render(ctx context.Context, sessionID uint64, tt ...types.DocumentType) error {
	s := svc.sessions.FindByID(sessionID)
	if s == nil {
		return errors.New("session not found")
	}

	r := s.ToRequest()

	for _, t := range tt {
		dd := svc.drivers.GetDrivers(t)
		for _, d := range dd {
			out, err := d.Render(ctx, r)
			if err != nil {
				return err
			}
			out = out
		}
	}

	return nil
}

// @todo additional methods to work with templates, get documents, ...

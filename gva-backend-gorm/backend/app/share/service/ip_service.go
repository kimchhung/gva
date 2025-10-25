package service

import (
	admincontext "backend/app/admin/context"
	apperror "backend/app/share/error"
	"backend/env"
	"net"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/oschwald/geoip2-golang"
)

type IPService struct {
	db *geoip2.Reader
}

func NewIPService(cfg *env.Config) *IPService {
	ip_s := IPService{}

	go func() {
		if err := ip_s.OpenDB(); err != nil {
			panic(err)
		}
	}()

	return &ip_s
}

func (s *IPService) OpenDB() error {
	path := "./internal/geo2/GeoLite2-Country.mmdb"
	absPath, err := filepath.Abs(path)

	if err != nil {
		return err
	}

	db, err := geoip2.Open(absPath)
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *IPService) VerifyWhiteListIP(currentIP string, IPWhiteList []string) error {
	fullIpAccess := "0.0.0.0"

	if len(IPWhiteList) == 0 {
		return nil
	}

	for _, ip := range IPWhiteList {
		if ip == fullIpAccess || ip == currentIP {
			return nil
		}
	}
	return apperror.ErrAdminWhitelistInvalid
}

func (s *IPService) RequiredWhiteListIP() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			admin := admincontext.MustAdminContext(c.Request().Context()).Admin
			if admin.IpWhiteList != nil {
				if err := s.VerifyWhiteListIP(s.GetCurrentIP(c), admin.IpWhiteList); err != nil {
					return err
				}
			}
			return next(c)
		}
	}
}

func (s *IPService) GetIPRecord(currentIP string) (*geoip2.Country, error) {
	parseIP := net.ParseIP(currentIP)

	if parseIP == nil {
		return nil, apperror.ErrInvalidIP
	}

	return s.db.Country(net.ParseIP(currentIP))
}

func (s *IPService) GetCurrentIP(c echo.Context) string {
	return c.RealIP()
}

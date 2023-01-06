package db

import (
	"github.com/zu1k/nali/internal/constant"
	"github.com/zu1k/nali/pkg/cdn"
	"github.com/zu1k/nali/pkg/ip2region"
	"github.com/zu1k/nali/pkg/qqwry"
	"path/filepath"
)

func GetDefaultDBList() List {
	return List{
		&DB{
			Name: "qqwry",
			NameAlias: []string{
				"chunzhen",
			},
			Format:       FormatQQWry,
			File:         filepath.Join(constant.DataDirPath, "qqwry.dat"),
			Languages:    LanguagesZH,
			Types:        TypesIPv4,
			DownloadUrls: qqwry.DownloadUrls,
		},
		&DB{
			Name: "zxipv6wry",
			NameAlias: []string{
				"zxipv6",
				"zx",
			},
			Format:    FormatZXIPv6Wry,
			File:      filepath.Join(constant.DataDirPath, "zxipv6wry.db"),
			Languages: LanguagesZH,
			Types:     TypesIPv6,
		},
		&DB{
			Name: "geoip",
			NameAlias: []string{
				"geoip2",
				"geolite",
				"geolite2",
			},
			Format:    FormatMMDB,
			File:      filepath.Join(constant.DataDirPath, "GeoLite2-City.mmdb"),
			Languages: LanguagesAll,
			Types:     TypesIP,
		},
		&DB{
			Name: "dbip",
			NameAlias: []string{
				"db-ip",
			},
			Format:    FormatMMDB,
			File:      filepath.Join(constant.DataDirPath, "dbip.mmdb"),
			Languages: LanguagesAll,
			Types:     TypesIP,
		},
		&DB{
			Name:      "ipip",
			Format:    FormatIPIP,
			File:      filepath.Join(constant.DataDirPath, "ipipfree.ipdb"),
			Languages: LanguagesZH,
			Types:     TypesIP,
		},
		&DB{
			Name: "ip2region",
			NameAlias: []string{
				"i2r",
			},
			Format:       FormatIP2Region,
			File:         filepath.Join(constant.DataDirPath, "ip2region.xdb"),
			Languages:    LanguagesZH,
			Types:        TypesIPv4,
			DownloadUrls: ip2region.DownloadUrls,
		},
		&DB{
			Name:      "ip2location",
			Format:    FormatIP2Location,
			File:      filepath.Join(constant.DataDirPath, "IP2LOCATION-LITE-DB3.IPV6.BIN"),
			Languages: LanguagesEN,
			Types:     TypesIP,
		},

		&DB{
			Name:         "cdn",
			Format:       FormatCDNYml,
			File:         filepath.Join(constant.DataDirPath, "cdn.yml"),
			Languages:    LanguagesZH,
			Types:        TypesCDN,
			DownloadUrls: cdn.DownloadUrls,
		},
	}
}

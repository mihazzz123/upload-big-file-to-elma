package company

import (
	"fmt"
	"net/http"
	"strings"

	"git.elewise.com/elma365/common/pkg/config"
	"git.elewise.com/elma365/common/pkg/errs"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

const (
	// CompanyHeaderName названия хидера, хранящего компанию
	CompanyHeaderName = "X-Company"
	// DefaultCompany компания, используемае в случае отсуствия хидера или части урла
	DefaultCompany = "head"
	// список алиасов для компаний
	companyAliasKey = "%s:companies:alias"
)

var (
	// ErrNotFound ошибка при отсутствии компании
	ErrNotFound = errors.New("company name not found")
)

// Company компания
type Company string

// FromRequest получает компанию/алиас из текущего запроса
func FromRequest(cfg config.Config, r *http.Request) (Company, error) {
	company := parseByPriority(cfg, r, "")
	if company == "" {
		sources := make(map[string]string)
		sources[CompanyHeaderName] = r.Header.Get(CompanyHeaderName)
		sources["Host"] = r.Header.Get("Host")
		sources["Origin"] = r.Header.Get("Origin")
		return "", errs.WithData(errors.WithStack(ErrNotFound), fmt.Sprintf("missing sources: %+v", sources))
	}

	if !testCompanyName(company) {
		return "", errors.WithStack(errors.Errorf("invalid company name %q", company))
	}

	return Company(company), nil
}

// FromURL получает компанию/алиас из заданного адреса
func FromURL(cfg config.Config, additionalURL string) (Company, error) {
	company := parseByPriority(cfg, nil, additionalURL)
	if company == "" {
		return "", errors.WithStack(ErrNotFound)
	}

	if !testCompanyName(company) {
		return "", errors.WithStack(errors.Errorf("invalid company name %q", company))
	}

	return Company(company), nil
}

// FromAlias получает значение заданного алиаса компании
func FromAlias(red RedisStorage, namespacePrefix, alias string) (string, error) {
	companyName, err := red.HGet(fmt.Sprintf(companyAliasKey, namespacePrefix), alias).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}

		return "", errors.Wrap(err, "redis get alias failed")
	}

	return companyName, nil
}

// RedisStorage интерфейс взаимодействия с хранилищем алиасов
type RedisStorage interface {
	HGet(key, field string) *redis.StringCmd
}

func parseByPriority(cfg config.Config, r *http.Request, additionalURL string) string {
	// Подумать над оптимизацией этого запроса, так же чтобы он работал на площадках разработчиков
	if r != nil {
		company := r.Header.Get(CompanyHeaderName)
		if company != "" {
			return company
		}

		company = extractFromHost(cfg, r.Host)
		if company != "" {
			return company
		}

		origin := r.Header.Get("Origin")
		originHost := getHostFromString(origin)
		company = extractFromHost(cfg, originHost)
		if company != "" {
			return company
		}
	}

	if additionalURL != "" {
		additionalHost := getHostFromString(additionalURL)
		company := extractFromHost(cfg, additionalHost)
		if company != "" {
			return company
		}
	}

	if cfg.Solution == config.Onpremise {
		// для Onpremise - в head
		return DefaultCompany
	}

	return ""
}

func extractFromHost(cfg config.Config, host string) string {
	if host == "" {
		return ""
	}

	baseHost := strings.ToLower(cfg.GetHost())
	host = strings.ToLower(host)

	if baseHost == "" {
		return ""
	}

	if !strings.HasSuffix(host, baseHost) {
		// если это какой-то левый домен прилетел
		return ""
	}

	company := strings.TrimSuffix(host, baseHost)
	company = strings.TrimSuffix(company, ".")

	return company
}

func getHostFromString(value string) string {
	if value == "" {
		return ""
	}

	// Отсекаем протокол
	parts := strings.Split(value, "://")
	host := parts[0]
	if len(parts) > 1 {
		host = parts[1]
	}

	// Отсекаем параметры
	parts = strings.Split(host, "/")
	host = parts[0]

	return host
}

const (
	code0         = rune('0')
	code9         = rune('9')
	codeALower    = rune('a')
	codeZLower    = rune('z')
	codeAUpper    = rune('A')
	codeZUpper    = rune('Z')
	codeMinus     = rune('-')
	codeUnderline = rune('_')
)

// testCompanyName выполняет проверку аналогичную регулярному выражению ^[0-9a-zA-Z][\w\-]+$.
//nolint:gocyclo // нечего выносить
func testCompanyName(name string) bool {
	if name == "" {
		return false
	}
	strRunes := []rune(name)
	result := true
	if strRunes[0] == codeMinus ||
		strRunes[len(strRunes)-1] == codeMinus ||
		strRunes[0] == codeUnderline ||
		strRunes[len(strRunes)-1] == codeUnderline {
		return false
	}
	for _, sym := range strRunes {
		result = result && (sym == codeMinus || sym == codeUnderline ||
			(sym >= code0 && sym <= code9) ||
			(sym >= codeALower && sym <= codeZLower) ||
			(sym >= codeAUpper && sym <= codeZUpper))
	}

	return result
}

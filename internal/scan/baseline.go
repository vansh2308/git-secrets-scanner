package scan

import (
	"github.com/rs/zerolog/log"
	"github.com/vansh2308/git-secrets-scanner.git/internal/report"
	"github.com/vansh2308/git-secrets-scanner.git/internal/report/secret"
)

var baseline map[string]secret.SecretSet

func SetBaseline(b map[string]secret.SecretSet) {
	baseline = b
}

func GetRepoBaseline(repository string) secret.SecretSet {
	return baseline[repository]
}

func LoadBaseline(path string) error {
	if path != "" {
		report, err := report.ReadCSVReport(path)
		if err != nil {
			return err
		}

		b := make(map[string]secret.SecretSet)
		num := 0

		for _, s := range report {
			_, ok := b[s.Repository]
			if !ok {
				b[s.Repository] = secret.NewSet()
			}
			b[s.Repository].Add(s)
			num += 1
		}

		SetBaseline(b)

		log.Debug().
			Int("num", num).
			Msg("baseline loaded")
	}
	return nil
}

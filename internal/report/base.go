package report

import "github.com/vansh2308/git-secrets-scanner.git/internal/report/secret"

type ReportWriter interface {
	WriteAll(s []*secret.Secret) error
	Close() error
}

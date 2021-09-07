package pkg

import (
	"github.com/aquasecurity/trivy/pkg/report"
	"github.com/xanzy/go-gitlab"
)

var git *gitlab.Client

type trivy struct {
	ProjName        string
	State           string
	Vulnerabilities int
	Ignore          []string
	ReportResult    report.Results
}

type TrivyResults []*trivy

func (t *TrivyResults) Check() {
	for _, result := range *t {
		var countVulnies int
		for _, pkgResult := range result.ReportResult {
			countVulnies += len(pkgResult.Vulnerabilities)
		}
		result.Vulnerabilities = countVulnies
	}
}

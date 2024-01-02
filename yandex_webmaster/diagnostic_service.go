package yandexwebmaster

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DiagnosticService - service for diagnostic
type DiagnosticService struct {
	client *Client
}

// newDiagnosticService - init DiagnosticService
func newDiagnosticService(cl *Client) *DiagnosticService {
	return &DiagnosticService{client: cl}
}

// GetDiagnositcs - get site diagnostics, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-diagnostics-get.html#response-format__ap-sites-problem-type
func (s *DiagnosticService) GetDiagnositcs(hostID string) (DiagnosticProblemsResponse, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s/diagnostics", s.client.userID, hostID)
	var result DiagnosticProblemsResponse
	_, err := s.client.sendAPIRequest(http.MethodGet, endpoint, nil, &result)
	return result, err
}

type DiagnosticProblem struct {
	Type            string
	Severity        string
	State           string
	LastStateUpdate string
}

func (d *DiagnosticProblem) UnmarshalJSON(bytes []byte) error {
	type raw struct {
		Severity        string `json:"severity"`
		State           string `json:"state"`
		LastStateUpdate string `json:"last_state_update"`
	}
	var r raw
	err := json.Unmarshal(bytes, &r)
	if err != nil {
		return err
	}
	d.Severity = r.Severity
	d.State = r.State
	d.LastStateUpdate = r.LastStateUpdate
	return nil
}

type DiagnosticProblems []DiagnosticProblem

type DiagnosticProblemsResponse struct {
	Problems DiagnosticProblems `json:"problems"`
}

func (d *DiagnosticProblems) UnmarshalJSON(bytes []byte) error {
	dpMap := make(map[string]DiagnosticProblem)
	err := json.Unmarshal(bytes, &dpMap)
	if err != nil {
		return err
	}
	problems := make(DiagnosticProblems, 0, len(dpMap))
	for t, p := range dpMap {
		p.Type = t
		problems = append(problems, p)
	}
	*d = problems
	return nil

}

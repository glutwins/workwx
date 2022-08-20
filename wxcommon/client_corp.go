package wxcommon

import "fmt"

type SuiteCorpClient struct {
	CorpId     string
	CorpSecret string
	AgentId    int
	SuiteClient
}

func (scc *SuiteCorpClient) TicketGet() (string, error) {
	ticket, err := scc.TokenStore.GetSuiteJsTicket(scc.SuiteId, scc.CorpId)
	if err != nil {
		return "", err
	}
	if ticket == "" {
		accessToken, err := scc.TokenHandler()
		if err != nil {
			return "", err
		}
		resp := &TicketGetResp{}
		if err := scc.GetJSON(fmt.Sprintf("/ticket/get?access_token=%s", accessToken), resp); err != nil {
			return "", err
		}

		scc.TokenStore.SetSuiteJsTicket(scc.SuiteId, scc.CorpId, resp.Ticket, resp.ExpiresIn)
		return resp.Ticket, nil
	}
	return ticket, nil
}

func (scc *SuiteCorpClient) TicketGetAgent() (string, error) {
	ticket, err := scc.TokenStore.GetSuiteAgentJsTicket(scc.SuiteId, scc.CorpId)
	if err != nil {
		return "", err
	}
	if ticket == "" {
		accessToken, err := scc.TokenHandler()
		if err != nil {
			return "", err
		}
		resp := &TicketGetResp{}
		if err := scc.GetJSON(fmt.Sprintf("/ticket/get?access_token=%s&type=agent_config", accessToken), resp); err != nil {
			return "", err
		}
		scc.TokenStore.SetSuiteAgentJsTicket(scc.SuiteId, scc.CorpId, resp.Ticket, resp.ExpiresIn)
		return resp.Ticket, nil
	}
	return ticket, nil
}

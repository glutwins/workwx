package wxcommon

import "context"

type SuiteCorpClient struct {
	CorpId     string
	CorpSecret string
	AgentId    int
	SuiteClient
}

func (scc *SuiteCorpClient) SuiteCorpClientWithContext(c context.Context) *SuiteCorpClient {
	var nsc = *scc
	nsc.Context = c
	return &nsc
}

func (scc *SuiteCorpClient) TicketGet() (string, error) {
	ticket, err := scc.TokenStore.GetSuiteJsTicket(scc.SuiteId, scc.CorpId)
	if err != nil {
		return "", err
	}
	if ticket == "" {
		resp := &TicketGetResp{}
		if err := scc.GetRespWithToken("/ticket/get?access_token=%s", resp); err != nil {
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
		resp := &TicketGetResp{}
		if err := scc.GetRespWithToken("/ticket/get?access_token=%s&type=agent_config", resp); err != nil {
			return "", err
		}
		scc.TokenStore.SetSuiteAgentJsTicket(scc.SuiteId, scc.CorpId, resp.Ticket, resp.ExpiresIn)
		return resp.Ticket, nil
	}
	return ticket, nil
}

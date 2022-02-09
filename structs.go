package main

type WO struct {
	WoUuid                  string `json:"wo_uuid"`
	Client                  string `json:"client"`
	WorkClient              string `json:"work_client"`
	WoNumber                string `json:"wo_number"`
	WoRecipient             string `json:"wo_recipient"`
	WoNeedsClientAcceptance bool   `json:"wo_needs_client_acceptance"`
	WoClientAccepted        bool   `json:"wo_client_accepted"`
}

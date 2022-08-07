package types

type Endpoint struct {
	Path      string
	ApiPath   string // Path from the OpenAPI spec
	Operation string
}

func (e *Endpoint) GetApiPath() string {
	if e.ApiPath == "" {
		return e.Path
	}
	return e.ApiPath
}

var Endpoints = map[string]Endpoint{
	"AuditLogsList": {
		Path:      "/v1/audit_logs",
		Operation: "GET",
	},
	"Health": {
		Path:      "/v1/health",
		Operation: "GET",
	},
	"ListSecretIncidents": {
		Path:      "/v1/incidents/secrets",
		Operation: "GET",
	},
	"GetSecretIncidents": {
		Path:      "/v1/incidents/secrets/%d",
		ApiPath:   "/v1/incidents/secrets/{incident_id}",
		Operation: "GET",
	},
	"UpdateSecretIncidents": {
		Path:      "/v1/incidents/secrets/%d",
		ApiPath:   "/v1/incidents/secrets/{incident_id}",
		Operation: "PATCH",
	},
	"AssignSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/assign",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/assign",
		Operation: "POST",
	},
	"GrantAccessSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/grant_access",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/grant_access",
		Operation: "POST",
	},
	"IgnoreSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/ignore",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/ignore",
		Operation: "POST",
	},
	"ListSecretMembers": {
		Path:      "/v1/incidents/secrets/%d/members",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/members",
		Operation: "GET",
	},
	"CreateNote": {
		Path:      "/v1/incidents/secrets/%d/notes",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/notes",
		Operation: "POST",
	},
	"UpdateNote": {
		Path:      "/v1/incidents/secrets/%d/notes/%d",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/notes/{note_id}",
		Operation: "PATCH",
	},
	"DeleteNote": {
		Path:      "/v1/incidents/secrets/%d/notes/%d",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/notes/{note_id}",
		Operation: "DELETE",
	},
	"ListNotes": {
		Path:      "/v1/incidents/secrets/%d/notes",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/notes",
		Operation: "GET",
	},
	"ReopenSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/reopen",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/reopen",
		Operation: "POST",
	},
	"ResolveSecretIncident": {
		Path:      "/v1/incidents/secrets/{incident_id}/resolve",
		Operation: "POST",
	},
	"RevokeAccessSecretIncident": {
		Path:      "/v1/incidents/secrets/{incident_id}/revoke_access",
		Operation: "POST",
	},
	"ShareSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/share",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/share",
		Operation: "POST",
	},
	"UnassignSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/unassign",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/unassign",
		Operation: "POST",
	},
	"UnshareSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/unshare",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/unshare",
		Operation: "POST",
	},
	"ListOccurrences": {
		Path:      "/v1/occurrences/secrets",
		Operation: "GET",
	},
	"InvitationsList": {
		Path:      "/v1/invitations",
		Operation: "GET",
	},
	"InvitationsCreate": {
		Path:      "/v1/invitations",
		Operation: "POST",
	},
	"InvitationsDelete": {
		Path:      "/v1/invitations/%d",
		ApiPath:   "/v1/invitations/{invitation_id}",
		Operation: "DELETE",
	},
	"InvitationsResend": {
		Path:      "/v1/invitations/%d/resend",
		ApiPath:   "/v1/invitations/{invitation_id}/resend",
		Operation: "POST",
	},
	"MembersList": {
		Path:      "/v1/members",
		Operation: "GET",
	},
	"Scan": {
		Path:      "/v1/scan",
		Operation: "POST",
	},
	"ScanMultiple": {
		Path:      "/v1/multiscan",
		Operation: "POST",
	},
	"ScanQuotas": {
		Path:      "/v1/quotas",
		Operation: "GET",
	},
	"SourcesList": {
		Path:      "/v1/sources",
		Operation: "GET",
	},
	"SourcesGet": {
		Path:      "/v1/sources/%d",
		ApiPath:   "/v1/sources/{source_id}",
		Operation: "GET",
	},
}

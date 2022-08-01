package types

type Endpoint struct {
	Path      string
	Operation string
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
		Operation: "GET",
	},
	"UpdateSecretIncidents": {
		Path:      "/v1/incidents/secrets/%d",
		Operation: "PATCH",
	},
	"AssignSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/assign",
		Operation: "POST",
	},
	"GrantAccessSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/grant_access",
		Operation: "POST",
	},
	"IgnoreSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/ignore",
		Operation: "POST",
	},
	"ListSecretMembers": {
		Path:      "/v1/incidents/secrets/%d/members",
		Operation: "GET",
	},
	"CreateNote": {
		Path:      "/v1/incidents/secrets/%d/notes",
		Operation: "POST",
	},
	"UpdateNote": {
		Path:      "/v1/incidents/secrets/%d/notes/%d",
		Operation: "PATCH",
	},
	"ListNotes": {
		Path:      "/v1/incidents/secrets/%d/notes",
		Operation: "GET",
	},
	"DeleteNote": {
		Path:      "/v1/incidents/secrets/%d/notes/%d",
		Operation: "DELETE",
	},
	"ReopenSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/reopen",
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
		Operation: "POST",
	},
	"UnassignSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/unassign",
		Operation: "POST",
	},
	"UnshareSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/unshare",
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
		Operation: "DELETE",
	},
	"InvitationsResend": {
		Path:      "/v1/invitations/%d/resend",
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
		Operation: "GET",
	},
}

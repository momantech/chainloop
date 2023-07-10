
{{- define "chainloop.postgresql.fullname" -}}
{{- include "common.names.dependency.fullname" (dict "chartName" "postgresql" "chartValues" .Values.postgresql "context" $) -}}
{{- end -}}

{{- define "chainloop.vault.fullname" -}}
{{- include "common.names.dependency.fullname" (dict "chartName" "vault" "chartValues" .Values.vault "context" $) -}}
{{- end -}}

{{/*
Returns a private key used for CAS <-> Controlplane communication
If we are running ind development mode we add a default one otherwise we require providing it
*/}}
{{- define "chainloop.casjwt.private_key" -}}
  {{- if .Values.development }}
    {{- coalesce .Values.casJWTPrivateKey (include "chainloop.casjwt.private_key.devel" .) }}
  {{- else }}
  {{- required "Authentication Private Key \"casJWTPrivateKey\" required" .Values.casJWTPrivateKey }}
  {{- end }}
{{- end }}

{{/*
Returns a public key used for CAS <-> Controlplane communication
If we are running ind development mode we add a default one otherwise we require providing it
*/}}
{{- define "chainloop.casjwt.public_key" -}}
  {{- if .Values.development }}
    {{- coalesce .Values.casJWTPublicKey (include "chainloop.casjwt.public_key.devel" .) }}
  {{- else }}
  {{- required "Authentication Public Key \"casJWTPublicKey\" required" .Values.casJWTPublicKey }}
  {{- end }}
{{- end }}

{{/*
DEVELOPMENT ONLY PRIVATE KEY
NOTE: It can not be generated by HELM because we also need a public key
*/}}
{{- define "chainloop.casjwt.private_key.devel" -}}
-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIA762MbJK9IBnaqG0sd9uFRM+Z7Y+Aq5UfmbWf0+acKMYpYoy/8kBE
tI6cpcA2KvmW5qurOjIMh5ISr+P2GmzSZX+gBwYFK4EEACOhgYkDgYYABAFzPMcM
NUnPoC7b+s+/OyxRC7V/+elthj6Cq85WCj0KZ2qDvmd4QsYnsTIQ7NM7E+9WztdP
rJBaMdfauMarLlc7/AAHqoa0lv7HNIa0PpupZD4VXmnIe/ZkhHvKOuw0Bdoq2D2B
3U25sylQQto3nZ4IqnsXmrtYGIFI9om3PoliT9/J7g==
-----END EC PRIVATE KEY-----
{{- end -}}

{{/*
DEVELOPMENT ONLY PUBLIC KEY
*/}}
{{- define "chainloop.casjwt.public_key.devel" -}}
-----BEGIN PUBLIC KEY-----
MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQBczzHDDVJz6Au2/rPvzssUQu1f/np
bYY+gqvOVgo9Cmdqg75neELGJ7EyEOzTOxPvVs7XT6yQWjHX2rjGqy5XO/wAB6qG
tJb+xzSGtD6bqWQ+FV5pyHv2ZIR7yjrsNAXaKtg9gd1NubMpUELaN52eCKp7F5q7
WBiBSPaJtz6JYk/fye4=
-----END PUBLIC KEY-----
{{- end -}}

{{- define "chainloop.credentials_service_settings" -}}
{{- with .Values.secretsBackend }}
{{- if eq .backend "vault" }}
vault:
  secretPrefix: {{ required "secret prefix required" .secretPrefix | quote }}
  {{- if and $.Values.development (or (not .vault) not .vault.address) }}
  address: {{ printf "http://%s:8200" (include "chainloop.vault.fullname" $) | quote }}
  token: {{ $.Values.vault.server.dev.devRootToken | quote }}
  {{- else if (required "vault backend selected but configuration not provided" .vault ) }}
  address: {{ required "vault address required" .vault.address | quote }}
  token: {{ required "vault token required" .vault.token | quote }}
  {{- end }}

{{- else if eq .backend "awsSecretManager" }}
awsSecretManager:
  secretPrefix: {{ required "secret prefix required" .secretPrefix | quote }}
  region: {{ required "region required" .awsSecretManager.region | quote }}
  creds:
    accessKey: {{ required "access key required" .awsSecretManager.accessKey | quote }}
    secretKey: {{ required "secret key required" .awsSecretManager.secretKey | quote }}
{{- end }}
{{- end }}
{{- end -}}

{{- define "chainloop.node_port" -}}
{{- if (and (or (eq .type "NodePort") (eq .type "LoadBalancer")) .nodePorts (not (empty .nodePorts.http))) }}
{{- .nodePorts.http }}
{{- else -}}
null
{{- end -}}
{{- end -}}

{{/*
Figure out the external URL the controlplane can be reached at
This endpoint is used for the CLI to know where to go for log in
NOTE: Load balancer service type is not supported
*/}}
{{- define "chainloop.controlplane.external_url" -}}
{{- $service := .Values.controlplane.service }}
{{- $ingress := .Values.controlplane.ingress }}

{{- if (and $ingress $ingress.enabled $ingress.hostname) }}
{{- printf "%s://%s" (ternary "https" "http" $ingress.tls ) $ingress.hostname }}
{{- else if (and (eq $service.type "NodePort") $service.nodePorts (not (empty $service.nodePorts.http))) }}
{{- printf "http://localhost:%s" $service.nodePorts.http }}
{{- else -}}
null
{{- end -}}
{{- end -}}

{{/*
##############################################################################
Controlplane helpers
##############################################################################
*/}}

{{/*
Chainloop Controlplane release name
*/}}
{{- define "chainloop.controlplane.fullname" -}}
{{- printf "%s-%s" (include "common.names.fullname" .) "controlplane" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Chainloop Controlplane Chart fullname
*/}}
{{- define "chainloop.controlplane.name" -}}
{{- printf "%s-%s" (include "common.names.name" .) "controlplane" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "chainloop.controlplane.labels" -}}
{{- include "common.labels.standard" . }}
app.kubernetes.io/part-of: chainloop
app.kubernetes.io/component: controlplane
{{- end }}

{{/*
Migration labels
*/}}
{{- define "chainloop.controlplane.migration.labels" -}}
{{- include "common.labels.standard" . }}
app.kubernetes.io/part-of: chainloop
app.kubernetes.io/component: controlplane-migration
{{- end }}

{{/*
Selector labels
*/}}
{{- define "chainloop.controlplane.selectorLabels" -}}
{{- include "common.labels.matchLabels" .}}
app.kubernetes.io/component: controlplane
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "controlplane.serviceAccountName" -}}
{{- if .Values.controlplane.serviceAccount.create }}
{{- default (include "chainloop.controlplane.fullname" .) .Values.controlplane.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.controlplane.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Return the Postgresql connection string
*/}}
{{- define "controlplane.database.connection_string" -}}
{{- printf "postgresql://%s:%s@%s:%s/%s" (include "controlplane.database.user" .) (include "controlplane.database.escapedPassword" .) (include "controlplane.database.host" .) (include "controlplane.database.port" .) (include "controlplane.database.name" .) }}
{{- end -}}

{{/*
Return the Postgresql connection string for Atlas migration
*/}}
{{- define "controlplane.database.atlas_connection_string" -}}
{{- printf "postgres://%s:%s@%s:%s/%s?sslmode=disable" (include "controlplane.database.user" .) (include "controlplane.database.escapedPassword" .) (include "controlplane.database.host" .) (include "controlplane.database.port" .) (include "controlplane.database.name" .) }}
{{- end -}}

{{/*
Return the Postgresql hostname
*/}}
{{- define "controlplane.database.host" -}}
{{- ternary (include "chainloop.postgresql.fullname" .) .Values.controlplane.externalDatabase.host .Values.postgresql.enabled -}}
{{- end -}}

{{/*
Return the Postgresql port
*/}}
{{- define "controlplane.database.port" -}}
{{- ternary 5432 .Values.controlplane.externalDatabase.port .Values.postgresql.enabled -}}
{{- end -}}

{{/*
Return the Postgresql password
*/}}
{{- define "controlplane.database.password" -}}
{{- if .Values.postgresql.enabled }}
    {{- if .Values.global.postgresql }}
        {{- if .Values.global.postgresql.auth }}
            {{- coalesce .Values.global.postgresql.auth.password .Values.postgresql.auth.password -}}
        {{- else -}}
            {{- .Values.postgresql.auth.password -}}
        {{- end -}}
    {{- else -}}
        {{- .Values.postgresql.auth.password -}}
    {{- end -}}
{{- else -}}
    {{- .Values.controlplane.externalDatabase.password -}}
{{- end -}}
{{- end -}}


{{/*
Return the URL-scaped Postgresql password 
*/}}
{{ define "controlplane.database.escapedPassword" -}}
  {{- include "controlplane.database.password" . | urlquery | replace "+" "%20" -}}
{{- end -}}

{{/*
Return the Postgresql database name
*/}}
{{- define "controlplane.database.name" -}}
{{- if .Values.postgresql.enabled }}
    {{- if .Values.global.postgresql }}
        {{- if .Values.global.postgresql.auth }}
            {{- coalesce .Values.global.postgresql.auth.database .Values.postgresql.auth.database -}}
        {{- else -}}
            {{- .Values.postgresql.auth.database -}}
        {{- end -}}
    {{- else -}}
        {{- .Values.postgresql.auth.database -}}
    {{- end -}}
{{- else -}}
    {{- .Values.controlplane.externalDatabase.database -}}
{{- end -}}
{{- end -}}

{{/*
Return the Postgresql user
*/}}
{{- define "controlplane.database.user" -}}
{{- if .Values.postgresql.enabled }}
    {{- if .Values.global.postgresql }}
        {{- if .Values.global.postgresql.auth }}
            {{- coalesce .Values.global.postgresql.auth.username .Values.postgresql.auth.username -}}
        {{- else -}}
            {{- .Values.postgresql.auth.username -}}
        {{- end -}}
    {{- else -}}
        {{- .Values.postgresql.auth.username -}}
    {{- end -}}
{{- else -}}
    {{- .Values.controlplane.externalDatabase.user -}}
{{- end -}}
{{- end -}}

{{/*
##############################################################################
CAS Helpers
##############################################################################
*/}}

{{/*
Chainloop CAS release name
*/}}

{{- define "chainloop.cas.fullname" -}}
{{- printf "%s-%s" (include "common.names.fullname" .) "cas" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Chainloop CAS Chart fullname
*/}}
{{- define "chainloop.cas.name" -}}
{{- printf "%s-%s" (include "common.names.name" .) "cas" | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{/*

Common labels
*/}}
{{- define "chainloop.cas.labels" -}}
{{- include "common.labels.standard" . }}
app.kubernetes.io/part-of: chainloop
app.kubernetes.io/component: cas
{{- end }}

{{/*
Selector labels
*/}}
{{- define "chainloop.cas.selectorLabels" -}}
{{- include "common.labels.matchLabels" .}}
app.kubernetes.io/component: cas
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "chainloop.cas.serviceAccountName" -}}
{{- if .Values.cas.serviceAccount.create }}
{{- default (include "chainloop.cas.fullname" .) .Values.cas.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.cas.serviceAccount.name }}
{{- end }}
{{- end }}

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package annotations

import (
	"github.com/rikatz/ingress-nginx-annotations/annotations/alias"
	"github.com/rikatz/ingress-nginx-annotations/annotations/auth"
	"github.com/rikatz/ingress-nginx-annotations/annotations/authreq"
	"github.com/rikatz/ingress-nginx-annotations/annotations/authreqglobal"
	"github.com/rikatz/ingress-nginx-annotations/annotations/authtls"
	"github.com/rikatz/ingress-nginx-annotations/annotations/backendprotocol"
	"github.com/rikatz/ingress-nginx-annotations/annotations/canary"
	"github.com/rikatz/ingress-nginx-annotations/annotations/clientbodybuffersize"
	"github.com/rikatz/ingress-nginx-annotations/annotations/connection"
	"github.com/rikatz/ingress-nginx-annotations/annotations/cors"
	"github.com/rikatz/ingress-nginx-annotations/annotations/customheaders"
	"github.com/rikatz/ingress-nginx-annotations/annotations/customhttperrors"
	"github.com/rikatz/ingress-nginx-annotations/annotations/defaultbackend"
	"github.com/rikatz/ingress-nginx-annotations/annotations/disableproxyintercepterrors"
	"github.com/rikatz/ingress-nginx-annotations/annotations/fastcgi"
	"github.com/rikatz/ingress-nginx-annotations/annotations/http2pushpreload"
	"github.com/rikatz/ingress-nginx-annotations/annotations/ipallowlist"
	"github.com/rikatz/ingress-nginx-annotations/annotations/ipdenylist"
	"github.com/rikatz/ingress-nginx-annotations/annotations/loadbalancing"
	"github.com/rikatz/ingress-nginx-annotations/annotations/log"
	"github.com/rikatz/ingress-nginx-annotations/annotations/mirror"
	"github.com/rikatz/ingress-nginx-annotations/annotations/modsecurity"
	"github.com/rikatz/ingress-nginx-annotations/annotations/opentelemetry"
	"github.com/rikatz/ingress-nginx-annotations/annotations/portinredirect"
	"github.com/rikatz/ingress-nginx-annotations/annotations/proxy"
	"github.com/rikatz/ingress-nginx-annotations/annotations/proxyssl"
	"github.com/rikatz/ingress-nginx-annotations/annotations/ratelimit"
	"github.com/rikatz/ingress-nginx-annotations/annotations/redirect"
	"github.com/rikatz/ingress-nginx-annotations/annotations/rewrite"
	"github.com/rikatz/ingress-nginx-annotations/annotations/satisfy"
	"github.com/rikatz/ingress-nginx-annotations/annotations/serversnippet"
	"github.com/rikatz/ingress-nginx-annotations/annotations/serviceupstream"
	"github.com/rikatz/ingress-nginx-annotations/annotations/sessionaffinity"
	"github.com/rikatz/ingress-nginx-annotations/annotations/snippet"
	"github.com/rikatz/ingress-nginx-annotations/annotations/sslcipher"
	"github.com/rikatz/ingress-nginx-annotations/annotations/sslpassthrough"
	"github.com/rikatz/ingress-nginx-annotations/annotations/streamsnippet"
	"github.com/rikatz/ingress-nginx-annotations/annotations/upstreamhashby"
	"github.com/rikatz/ingress-nginx-annotations/annotations/upstreamvhost"
	"github.com/rikatz/ingress-nginx-annotations/annotations/xforwardedprefix"
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

// Extractor defines the annotation parsers to be used in the extraction of annotations
type Extractor struct {
	annotations map[string]parser.IngressAnnotation
}

func NewAnnotationFactory() map[string]parser.IngressAnnotation {
	return map[string]parser.IngressAnnotation{
		"Aliases":                     alias.NewParser(cfg),
		"BasicDigestAuth":             auth.NewParser(auth.AuthDirectory, cfg),
		"Canary":                      canary.NewParser(cfg),
		"CertificateAuth":             authtls.NewParser(cfg),
		"ClientBodyBufferSize":        clientbodybuffersize.NewParser(cfg),
		"CustomHeaders":               customheaders.NewParser(cfg),
		"ConfigurationSnippet":        snippet.NewParser(cfg),
		"Connection":                  connection.NewParser(cfg),
		"CorsConfig":                  cors.NewParser(cfg),
		"CustomHTTPErrors":            customhttperrors.NewParser(cfg),
		"DisableProxyInterceptErrors": disableproxyintercepterrors.NewParser(cfg),
		"DefaultBackend":              defaultbackend.NewParser(cfg),
		"FastCGI":                     fastcgi.NewParser(cfg),
		"ExternalAuth":                authreq.NewParser(cfg),
		"EnableGlobalAuth":            authreqglobal.NewParser(cfg),
		"HTTP2PushPreload":            http2pushpreload.NewParser(cfg),
		"Opentelemetry":               opentelemetry.NewParser(cfg),
		"Proxy":                       proxy.NewParser(cfg),
		"ProxySSL":                    proxyssl.NewParser(cfg),
		"RateLimit":                   ratelimit.NewParser(cfg),
		"Redirect":                    redirect.NewParser(cfg),
		"Rewrite":                     rewrite.NewParser(cfg),
		"Satisfy":                     satisfy.NewParser(cfg),
		"ServerSnippet":               serversnippet.NewParser(cfg),
		"ServiceUpstream":             serviceupstream.NewParser(cfg),
		"SessionAffinity":             sessionaffinity.NewParser(cfg),
		"SSLPassthrough":              sslpassthrough.NewParser(cfg),
		"UsePortInRedirects":          portinredirect.NewParser(cfg),
		"UpstreamHashBy":              upstreamhashby.NewParser(cfg),
		"LoadBalancing":               loadbalancing.NewParser(cfg),
		"UpstreamVhost":               upstreamvhost.NewParser(cfg),
		"Allowlist":                   ipallowlist.NewParser(cfg),
		"Denylist":                    ipdenylist.NewParser(cfg),
		"XForwardedPrefix":            xforwardedprefix.NewParser(cfg),
		"SSLCipher":                   sslcipher.NewParser(cfg),
		"Logs":                        log.NewParser(cfg),
		"BackendProtocol":             backendprotocol.NewParser(cfg),
		"ModSecurity":                 modsecurity.NewParser(cfg),
		"Mirror":                      mirror.NewParser(cfg),
		"StreamSnippet":               streamsnippet.NewParser(cfg),
	}
}

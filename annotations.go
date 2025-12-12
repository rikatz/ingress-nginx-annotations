/*
Copyright 2025 The Kubernetes Authors.

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
	"log/slog"
	"maps"

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

func NewAnnotationFactory() parser.AnnotationFields {
	factory := make(parser.AnnotationFields)
	maps.Copy(factory, alias.AliasAnnotation.Annotations)
	maps.Copy(factory, auth.AuthSecretAnnotations.Annotations)
	maps.Copy(factory, authreq.AuthReqAnnotations.Annotations)
	maps.Copy(factory, authreqglobal.GlobalAuthAnnotations.Annotations)
	maps.Copy(factory, authtls.AuthTLSAnnotations.Annotations)
	maps.Copy(factory, backendprotocol.BackendProtocolConfig.Annotations)
	maps.Copy(factory, canary.CanaryAnnotations.Annotations)
	maps.Copy(factory, clientbodybuffersize.ClientBodyBufferSizeConfig.Annotations)
	maps.Copy(factory, connection.ConnectionHeadersAnnotations.Annotations)
	maps.Copy(factory, cors.CORSAnnotation.Annotations)
	maps.Copy(factory, customheaders.CustomHeadersAnnotation.Annotations)
	maps.Copy(factory, customhttperrors.CustomHTTPErrorsAnnotations.Annotations)
	maps.Copy(factory, defaultbackend.DefaultBackendAnnotations.Annotations)
	maps.Copy(factory, disableproxyintercepterrors.DisableProxyInterceptErrorsAnnotations.Annotations)
	maps.Copy(factory, fastcgi.FastCGIAnnotations.Annotations)
	maps.Copy(factory, http2pushpreload.HTTP2PushPreloadAnnotations.Annotations)
	maps.Copy(factory, ipallowlist.AllowlistAnnotations.Annotations)
	maps.Copy(factory, ipdenylist.DenylistAnnotations.Annotations)
	maps.Copy(factory, loadbalancing.LoadBalanceAnnotations.Annotations)
	maps.Copy(factory, log.LogAnnotations.Annotations)
	maps.Copy(factory, mirror.MirrorAnnotation.Annotations)
	maps.Copy(factory, modsecurity.ModsecurityAnnotation.Annotations)
	maps.Copy(factory, opentelemetry.OtelAnnotations.Annotations)
	maps.Copy(factory, portinredirect.PortsInRedirectAnnotations.Annotations)
	maps.Copy(factory, proxy.ProxyAnnotations.Annotations)
	maps.Copy(factory, proxyssl.ProxySSLAnnotation.Annotations)
	maps.Copy(factory, ratelimit.RateLimitAnnotations.Annotations)
	maps.Copy(factory, redirect.RedirectAnnotations.Annotations)
	maps.Copy(factory, rewrite.RewriteAnnotations.Annotations)
	maps.Copy(factory, satisfy.SatisfyAnnotations.Annotations)
	maps.Copy(factory, serversnippet.ServerSnippetAnnotations.Annotations)
	maps.Copy(factory, serviceupstream.ServiceUpstreamAnnotations.Annotations)
	maps.Copy(factory, sessionaffinity.SessionAffinityAnnotations.Annotations)
	maps.Copy(factory, snippet.ConfigurationSnippetAnnotations.Annotations)
	maps.Copy(factory, sslcipher.SSLCipherAnnotations.Annotations)
	maps.Copy(factory, sslpassthrough.SSLPassthroughAnnotations.Annotations)
	maps.Copy(factory, streamsnippet.StreamSnippetAnnotations.Annotations)
	maps.Copy(factory, upstreamhashby.UpstreamHashByAnnotations.Annotations)
	maps.Copy(factory, upstreamvhost.UpstreamVhostAnnotations.Annotations)
	maps.Copy(factory, xforwardedprefix.XForwardedForAnnotations.Annotations)

	for _, val := range factory {
		for _, alias := range val.AnnotationAliases {
			factory[alias] = val
		}
	}

	slog.Info("loaded annotations", "amount", len(factory))

	return factory
}

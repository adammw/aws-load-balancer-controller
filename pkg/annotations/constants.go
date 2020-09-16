package annotations

const (
	// IngressClass
	AnnotationIngressClass = "kubernetes.io/ingress.class"

	// IngressGroup
	AnnotationSuffixGroupName  = "group.name"
	AnnotationSuffixGroupOrder = "group.order"

	// Ingress annotation suffixes
	IngressSuffixTags                         = "tags"
	IngressSuffixIPAddressType                = "ip-address-type"
	IngressSuffixScheme                       = "scheme"
	IngressSuffixSubnets                      = "subnets"
	IngressSuffixLoadBalancerAttributes       = "load-balancer-attributes"
	IngressSuffixSecurityGroups               = "security-groups"
	IngressSuffixListenPorts                  = "listen-ports"
	IngressSuffixInboundCIDRs                 = "inbound-cidrs"
	IngressSuffixCertificateARN               = "certificate-arn"
	IngressSuffixSSLPolicy                    = "ssl-policy"
	IngressSuffixTargetType                   = "target-type"
	IngressSuffixBackendProtocol              = "backend-protocol"
	IngressSuffixTargetGroupAttributes        = "target-group-attributes"
	IngressSuffixHealthCheckPort              = "healthcheck-port"
	IngressSuffixHealthCheckProtocol          = "healthcheck-protocol"
	IngressSuffixHealthCheckPath              = "healthcheck-path"
	IngressSuffixHealthCheckIntervalSeconds   = "healthcheck-interval-seconds"
	IngressSuffixHealthCheckTimeoutSeconds    = "healthcheck-timeout-seconds"
	IngressSuffixHealthyThresholdCount        = "healthy-threshold-count"
	IngressSuffixUnhealthyThresholdCount      = "unhealthy-threshold-count"
	IngressSuffixSuccessCodes                 = "success-codes"
	IngressSuffixAuthType                     = "auth-type"
	IngressSuffixAuthIDPCognito               = "auth-idp-cognito"
	IngressSuffixAuthIDPOIDC                  = "auth-idp-oidc"
	IngressSuffixAuthOnUnauthenticatedRequest = "auth-on-unauthenticated-request"
	IngressSuffixAuthScope                    = "auth-scope"
	IngressSuffixAuthSessionCookie            = "auth-session-cookie"
	IngressSuffixAuthSessionTimeout           = "auth-session-timeout"

	// NLB annotation suffixes
	// prefixes service.beta.kubernetes.io, service.kubernetes.io
	SvcLBSuffixLoadBalancerType              = "aws-load-balancer-type"
	SvcLBSuffixInternal                      = "aws-load-balancer-internal"
	SvcLBSuffixProxyProtocol                 = "aws-load-balancer-proxy-protocol"
	SvcLBSuffixAccessLogEnabled              = "aws-load-balancer-access-log-enabled"
	SvcLBSuffixAccessLogS3BucketName         = "aws-load-balancer-access-log-s3-bucket-name"
	SvcLBSuffixAccessLogS3BucketPrefix       = "aws-load-balancer-access-log-s3-bucket-prefix"
	SvcLBSuffixCrossZoneLoadBalancingEnabled = "aws-load-balancer-cross-zone-load-balancing-enabled"
	SvcLBSuffixSSLCertificate                = "aws-load-balancer-ssl-cert"
	SvcLBSuffixSSLPorts                      = "aws-load-balancer-ssl-ports"
	SvcLBSuffixSSLNegotiationPolicy          = "aws-load-balancer-ssl-negotiation-policy"
	SvcLBSuffixBEProtocol                    = "aws-load-balancer-backend-protocol"
	SvcLBSuffixAdditionalTags                = "aws-load-balancer-additional-resource-tags"
	SvcLBSuffixHCHealthyThreshold            = "aws-load-balancer-healthcheck-healthy-threshold"
	SvcLBSuffixHCUnhealthyThreshold          = "aws-load-balancer-healthcheck-unhealthy-threshold"
	SvcLBSuffixHCTimeout                     = "aws-load-balancer-healthcheck-timeout"
	SvcLBSuffixHCInterval                    = "aws-load-balancer-healthcheck-interval"
	SvcLBSuffixHCProtocol                    = "aws-load-balancer-healthcheck-protocol"
	SvcLBSuffixHCPort                        = "aws-load-balancer-healthcheck-port"
	SvcLBSuffixHCPath                        = "aws-load-balancer-healthcheck-path"
)

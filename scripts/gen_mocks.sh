#!/usr/bin/env bash
# This shell script is used to generate & update mock objects for testing
mockery -name CloudAPI -dir ./internal/aws/
mockery -name Storer -dir ./internal/ingress/controller/store/ -inpkg

mockery -name Controller -dir ./internal/alb/tags/ -inpkg
mockery -name Controller -dir ./internal/alb/ls/ -inpkg
mockery -name RulesController -dir ./internal/alb/ls/ -inpkg

mockery -name ACMAPI -dir ./vendor/github.com/aws/aws-sdk-go/service/acm/acmiface
mockery -name EC2API -dir ./vendor/github.com/aws/aws-sdk-go/service/ec2/ec2iface
mockery -name ELBV2API -dir ./vendor/github.com/aws/aws-sdk-go/service/elbv2/elbv2iface
mockery -name IAMAPI -dir ./vendor/github.com/aws/aws-sdk-go/service/iam/iamiface
mockery -name ResourceGroupsTaggingAPIAPI -dir ./vendor/github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/resourcegroupstaggingapiiface
mockery -name WAFRegionalAPI -dir ./vendor/github.com/aws/aws-sdk-go/service/wafregional/wafregionaliface
mockery -name WAFV2API -dir ./vendor/github.com/aws/aws-sdk-go/service/wafv2/wafv2iface
mockery -name ShieldAPI -dir ./vendor/github.com/aws/aws-sdk-go/service/shield/shieldiface



mockgen -destination=./mocks/aws-alb-ingress-controller/ingress/auth/mock.go github.com/kubernetes-sigs/aws-alb-ingress-controller/internal/ingress/auth Module
mockgen -destination=./mocks/controller-runtime/cache/mock.go sigs.k8s.io/controller-runtime/pkg/cache Cache
mockgen -destination=./mocks/controller-runtime/controller/mock.go sigs.k8s.io/controller-runtime/pkg/controller Controller
mockgen -destination=./mocks/controller-runtime/client/mock_client.go sigs.k8s.io/controller-runtime/pkg/client Client
mockgen -destination=./mocks/aws/services/mock_elbv2.go sigs.k8s.io/aws-load-balancer-controller/pkg/aws/services ELBV2
mockgen -destination=./mocks/aws/services/mock_ec2.go sigs.k8s.io/aws-load-balancer-controller/pkg/aws/services EC2
mockgen -destination=./mocks/webhook/mock_mutator.go sigs.k8s.io/aws-load-balancer-controller/pkg/webhook Mutator
mockgen -destination=./mocks/webhook/mock_validator.go sigs.k8s.io/aws-load-balancer-controller/pkg/webhook Validator
mockgen -destination=./mocks/k8s/mock_finalizer.go sigs.k8s.io/aws-load-balancer-controller/pkg/k8s FinalizerManager
mockgen -destination=./mocks/networking/mock_security_group_manager.go sigs.k8s.io/aws-load-balancer-controller/pkg/networking SecurityGroupManager
mockgen -destination=./mocks/networking/mock_subnet_resolver.go sigs.k8s.io/aws-load-balancer-controller/pkg/networking SubnetsResolver
mockgen -destination=./mocks/ingress/mock_cert_discovery.go sigs.k8s.io/aws-load-balancer-controller/pkg/ingress CertDiscovery

// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/conf/controlplane/config/v1"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/dispatcher"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/server"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/service"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/authz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/ca"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data"
	"github.com/chainloop-dev/chainloop/app/controlplane/plugins/sdk/v1"
	"github.com/chainloop-dev/chainloop/pkg/blobmanager/loader"
	"github.com/chainloop-dev/chainloop/pkg/credentials"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func wireApp(bootstrap *conf.Bootstrap, readerWriter credentials.ReaderWriter, logger log.Logger, availablePlugins sdk.AvailablePlugins, certificateAuthority ca.CertificateAuthority) (*app, func(), error) {
	confData := bootstrap.Data
	data_Database := confData.Database
	newConfig := newDataConf(data_Database)
	dataData, cleanup, err := data.NewData(newConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	membershipRepo := data.NewMembershipRepo(dataData, logger)
	organizationRepo := data.NewOrganizationRepo(dataData, logger)
	casBackendRepo := data.NewCASBackendRepo(dataData, logger)
	providers := loader.LoadProviders(readerWriter)
	casBackendUseCase := biz.NewCASBackendUseCase(casBackendRepo, readerWriter, providers, logger)
	integrationRepo := data.NewIntegrationRepo(dataData, logger)
	integrationAttachmentRepo := data.NewIntegrationAttachmentRepo(dataData, logger)
	workflowRepo := data.NewWorkflowRepo(dataData, logger)
	newIntegrationUseCaseOpts := &biz.NewIntegrationUseCaseOpts{
		IRepo:   integrationRepo,
		IaRepo:  integrationAttachmentRepo,
		WfRepo:  workflowRepo,
		CredsRW: readerWriter,
		Logger:  logger,
	}
	integrationUseCase := biz.NewIntegrationUseCase(newIntegrationUseCaseOpts)
	v := bootstrap.Onboarding
	organizationUseCase := biz.NewOrganizationUseCase(organizationRepo, casBackendUseCase, integrationUseCase, membershipRepo, v, logger)
	membershipUseCase := biz.NewMembershipUseCase(membershipRepo, organizationUseCase, logger)
	newUserUseCaseParams := &biz.NewUserUseCaseParams{
		UserRepo:            userRepo,
		MembershipUseCase:   membershipUseCase,
		OrganizationUseCase: organizationUseCase,
		OnboardingConfig:    v,
		Logger:              logger,
	}
	userUseCase := biz.NewUserUseCase(newUserUseCaseParams)
	robotAccountRepo := data.NewRobotAccountRepo(dataData, logger)
	auth := bootstrap.Auth
	robotAccountUseCase := biz.NewRootAccountUseCase(robotAccountRepo, workflowRepo, auth, logger)
	casCredentialsUseCase, err := biz.NewCASCredentialsUseCase(auth)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	bootstrap_CASServer := bootstrap.CasServer
	v2 := _wireValue
	casClientUseCase := biz.NewCASClientUseCase(casCredentialsUseCase, bootstrap_CASServer, logger, v2...)
	referrerRepo := data.NewReferrerRepo(dataData, workflowRepo, logger)
	referrerSharedIndex := bootstrap.ReferrerSharedIndex
	referrerUseCase, err := biz.NewReferrerUseCase(referrerRepo, workflowRepo, membershipRepo, referrerSharedIndex, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	apiTokenRepo := data.NewAPITokenRepo(dataData, logger)
	enforcer, err := authz.NewDatabaseEnforcer(data_Database)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	apiTokenUseCase, err := biz.NewAPITokenUseCase(apiTokenRepo, auth, enforcer, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	workflowContractRepo := data.NewWorkflowContractRepo(dataData, logger)
	workflowContractUseCase := biz.NewWorkflowContractUseCase(workflowContractRepo, logger)
	workflowUseCase := biz.NewWorkflowUsecase(workflowRepo, workflowContractUseCase, logger)
	v3 := serviceOpts(logger)
	workflowService := service.NewWorkflowService(workflowUseCase, workflowContractUseCase, v3...)
	orgInvitationRepo := data.NewOrgInvitation(dataData, logger)
	orgInvitationUseCase, err := biz.NewOrgInvitationUseCase(orgInvitationRepo, membershipRepo, userRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	confServer := bootstrap.Server
	authService, err := service.NewAuthService(userUseCase, organizationUseCase, membershipUseCase, orgInvitationUseCase, auth, confServer, v3...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	robotAccountService := service.NewRobotAccountService(robotAccountUseCase, v3...)
	workflowRunRepo := data.NewWorkflowRunRepo(dataData, logger)
	workflowRunUseCase, err := biz.NewWorkflowRunUseCase(workflowRunRepo, workflowRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	newWorkflowRunServiceOpts := &service.NewWorkflowRunServiceOpts{
		WorkflowRunUC:      workflowRunUseCase,
		WorkflowUC:         workflowUseCase,
		WorkflowContractUC: workflowContractUseCase,
		CredsReader:        readerWriter,
		Opts:               v3,
	}
	workflowRunService := service.NewWorkflowRunService(newWorkflowRunServiceOpts)
	attestationUseCase := biz.NewAttestationUseCase(casClientUseCase, logger)
	fanOutDispatcher := dispatcher.New(integrationUseCase, workflowUseCase, workflowRunUseCase, readerWriter, casClientUseCase, availablePlugins, logger)
	casMappingRepo := data.NewCASMappingRepo(dataData, casBackendRepo, logger)
	casMappingUseCase := biz.NewCASMappingUseCase(casMappingRepo, membershipRepo, logger)
	newAttestationServiceOpts := &service.NewAttestationServiceOpts{
		WorkflowRunUC:      workflowRunUseCase,
		WorkflowUC:         workflowUseCase,
		WorkflowContractUC: workflowContractUseCase,
		OCIUC:              casBackendUseCase,
		CredsReader:        readerWriter,
		IntegrationUseCase: integrationUseCase,
		CasCredsUseCase:    casCredentialsUseCase,
		AttestationUC:      attestationUseCase,
		FanoutDispatcher:   fanOutDispatcher,
		CASMappingUseCase:  casMappingUseCase,
		ReferrerUC:         referrerUseCase,
		OrgUC:              organizationUseCase,
		Opts:               v3,
	}
	attestationService := service.NewAttestationService(newAttestationServiceOpts)
	workflowContractService := service.NewWorkflowSchemaService(workflowContractUseCase, v3...)
	contextService := service.NewContextService(casBackendUseCase, userUseCase, v3...)
	casCredentialsService := service.NewCASCredentialsService(casCredentialsUseCase, casMappingUseCase, casBackendUseCase, enforcer, v3...)
	orgMetricsRepo := data.NewOrgMetricsRepo(dataData, logger)
	orgMetricsUseCase, err := biz.NewOrgMetricsUseCase(orgMetricsRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	orgMetricsService := service.NewOrgMetricsService(orgMetricsUseCase, v3...)
	integrationsService := service.NewIntegrationsService(integrationUseCase, workflowUseCase, availablePlugins, v3...)
	organizationService := service.NewOrganizationService(membershipUseCase, organizationUseCase, v3...)
	casBackendService := service.NewCASBackendService(casBackendUseCase, providers, v3...)
	casRedirectService, err := service.NewCASRedirectService(casMappingUseCase, casCredentialsUseCase, bootstrap_CASServer, v3...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	orgInvitationService := service.NewOrgInvitationService(orgInvitationUseCase, v3...)
	referrerService := service.NewReferrerService(referrerUseCase, v3...)
	apiTokenService := service.NewAPITokenService(apiTokenUseCase, v3...)
	attestationStateRepo := data.NewAttestationStateRepo(dataData, logger)
	attestationStateUseCase, err := biz.NewAttestationStateUseCase(attestationStateRepo, workflowRunRepo)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	newAttestationStateServiceOpt := &service.NewAttestationStateServiceOpt{
		AttestationStateUseCase: attestationStateUseCase,
		WorkflowUseCase:         workflowUseCase,
		WorkflowRunUseCase:      workflowRunUseCase,
		Opts:                    v3,
	}
	attestationStateService := service.NewAttestationStateService(newAttestationStateServiceOpt)
	userService := service.NewUserService(membershipUseCase, organizationUseCase, v3...)
	signingUseCase := biz.NewChainloopSigningUseCase(certificateAuthority)
	signingService := service.NewSigningService(signingUseCase, v3...)
	validator, err := newProtoValidator()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	opts := &server.Opts{
		UserUseCase:         userUseCase,
		RobotAccountUseCase: robotAccountUseCase,
		CASBackendUseCase:   casBackendUseCase,
		CASClientUseCase:    casClientUseCase,
		IntegrationUseCase:  integrationUseCase,
		ReferrerUseCase:     referrerUseCase,
		APITokenUseCase:     apiTokenUseCase,
		OrganizationUseCase: organizationUseCase,
		WorkflowUseCase:     workflowUseCase,
		WorkflowSvc:         workflowService,
		AuthSvc:             authService,
		RobotAccountSvc:     robotAccountService,
		WorkflowRunSvc:      workflowRunService,
		AttestationSvc:      attestationService,
		WorkflowContractSvc: workflowContractService,
		ContextSvc:          contextService,
		CASCredsSvc:         casCredentialsService,
		OrgMetricsSvc:       orgMetricsService,
		IntegrationsSvc:     integrationsService,
		OrganizationSvc:     organizationService,
		CASBackendSvc:       casBackendService,
		CASRedirectSvc:      casRedirectService,
		OrgInvitationSvc:    orgInvitationService,
		ReferrerSvc:         referrerService,
		APITokenSvc:         apiTokenService,
		AttestationStateSvc: attestationStateService,
		UserSvc:             userService,
		SigningSvc:          signingService,
		Logger:              logger,
		ServerConfig:        confServer,
		AuthConfig:          auth,
		Credentials:         readerWriter,
		Enforcer:            enforcer,
		Validator:           validator,
	}
	grpcServer, err := server.NewGRPCServer(opts)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	httpServer, err := server.NewHTTPServer(opts, grpcServer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	httpMetricsServer, err := server.NewHTTPMetricsServer(opts)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	workflowRunExpirerUseCase := biz.NewWorkflowRunExpirerUseCase(workflowRunRepo, logger)
	apiTokenSyncerUseCase := biz.NewAPITokenSyncerUseCase(apiTokenUseCase)
	mainApp := newApp(logger, grpcServer, httpServer, httpMetricsServer, workflowRunExpirerUseCase, availablePlugins, apiTokenSyncerUseCase)
	return mainApp, func() {
		cleanup()
	}, nil
}

var (
	_wireValue = []biz.CASClientOpts{}
)

// wire.go:

func newDataConf(in *conf.Data_Database) *data.NewConfig {
	return &data.NewConfig{Driver: in.Driver, Source: in.Source}
}

func serviceOpts(l log.Logger) []service.NewOpt {
	return []service.NewOpt{service.WithLogger(l)}
}

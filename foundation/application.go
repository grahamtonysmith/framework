package application

import (
	"github.com/grahamtonysmith/framework/config"
	"github.com/grahamtonysmith/framework/contracts"
	"github.com/grahamtonysmith/framework/facades"
)

func init() {
	app := Application{}
	app.registerBaseServiceProviders()
	app.bootBaseServiceProviders()
}

type Application struct {
}

func (app *Application) registerServiceProviders(serviceProviders []contracts.ServiceProvider) {
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Register()
	}
}

func (app *Application) bootServiceProviders(serviceProviders []contracts.ServiceProvider) {
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Boot()
	}
}

func (app *Application) getBaseServiceProviders() []contracts.ServiceProvider {
	return []contracts.ServiceProvider{
		&config.ServiceProvider{},
	}
}

func (app *Application) getConfiguredServiceProviders() []contracts.ServiceProvider {
	return facades.Config.Get("app.providers").([]contracts.ServiceProvider)
}

func (app *Application) registerBaseServiceProviders() {
	app.registerServiceProviders(app.getBaseServiceProviders())
}

func (app *Application) bootBaseServiceProviders() {
	app.bootServiceProviders(app.getBaseServiceProviders())
}

func (app *Application) registerConfiguredServiceProviders() {
	app.registerServiceProviders(app.getConfiguredServiceProviders())
}

func (app *Application) bootConfiguredServiceProviders() {
	app.bootServiceProviders(app.getConfiguredServiceProviders())
}

func (app *Application) Boot() {
	app.registerConfiguredServiceProviders()
	app.bootConfiguredServiceProviders()
}

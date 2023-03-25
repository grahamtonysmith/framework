package contracts

type ServiceProvider interface {
	// Register any application services
	Register()

	// Boot any application services after register
	Boot()
}

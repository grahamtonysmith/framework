package contracts

type ServiceProvider interface {
	// Registers an application service
	Register()

	// Boots an application service after register
	Boot()
}

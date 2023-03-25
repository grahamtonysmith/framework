### notes

- service provider Register method calls NewApplication, which is a factory
- if an implementation of a contact interface only has one implementation, it is called Application by default (see config)
- foundation.init will be used to register the default service providers, which is just the config service provider
### notes

- service provider Register method calls NewApplication, which is a factory
- if an implementation of a contact interface only has one implementation, it is called Application by default (see config)
- foundation.init will be used to register the default service providers, which is just the config service provider

I'm not in love with the structure. 

Having a contracts dir is nice because I don't have to worry about circular imports if I define my interfaces next to 
where they are needed

Having the everything occur by side effect is a bit off. I think I would prefer to have the foundation application be
a mediator between service providers and have it serve as a facade to the user
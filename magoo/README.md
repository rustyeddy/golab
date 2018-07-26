# What is Magoo?

## A Easy Going Form Processor

Magoo can be used to process forms submitted from just about anywhere,
you decide you want to recieve them from.

Magoo starts up a very lightweight server, opens a port (1199 by
default) and starts waiting for POSTs to /v01/entry/.

The values (entries) submitted to magoo will written to a repository of your
choosing (/srv/magfs/entries/) by default.

Zero configuration, just start magoo and start submitting forms.

### Zero Configuration

Magoo requires zero configuration, though it is highly configurable.
It just starts up with some really basic but useful defaults.

#### Storage - Filesystem as Default

Unless told otherwise, magoo uses a standard filesystem layout.
Entries are stored as JSON files and hence can easily be resurected
and used by any application capable of reading JSON.

All _Entries_ and _Forms_ are isolated between accounts.  Forms can be
created how YOU want, use the tools and format you love the most.  By 
default forms can be defined with JSON or YML will be combined with
professional templates and displayed to clients.

### Extensibility

You can create and add forms to your hearts content. You can supply
custom templates to define how your forms are going to look.

### Notifications

Magoo can send notifications via email, webhook or computer (MQTT) or
people (text) messaging is available. Just setup some simple
configurations and it is off to the races.

## MagooFS

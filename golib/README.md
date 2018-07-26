# golib

Is a _Go library_ (did you see that coming?) that acts as a library of utilities that help automation tasks that I get caught up in.  The utilities are:

## DNS

The DNS package provides

### **DNSHostRecords(*** *hostname string* ***)** ***DNSHostRecord***

This function will gather all DNS records from the given DNS server. The various DNS record for the given will be neatly stored in the *DNSHostRecord* for later consumption.

### ***NC*** - *Namecheap* domain name registrar functions

This module allows us to watch all of our domain names, check availability, buy more, renew old ones and set contact information.  It also lets you setup the nameserver, and manage DNS records.

### ***RLog*** Cherry pick log messages

This module accepts flags, which are runtime variables (may be changed at anytime) to select when to print certain debug messages or not.

### ***work*** Upwork atom reader

Read form various Upwork job feeds and take action on the data that has been returned.  This may be subordinated to getting access to upworks API.

### ***webster*** Generate Simple HTML Pages

Create simple HTML pages from Markdown and optional Front Matter to produce simple HTML pages.  This can be a powerful little tool for generating and managing a large number of small sites (or SPAs).

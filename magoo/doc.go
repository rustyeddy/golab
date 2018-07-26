/*

# Magoo Form Processor

Magoo serves up forms and accepts entries in response to form
submissions.

Forms are forwarded via static pages, the submissions are accepted
as form submissions and stored as JSON objects by the storage
module.

## Magoo Directory Layout

magfs:   (per user)
  public:
    - static files part of the public site
  forms:
    - Set of forms or form groups with an associated Ids.
      Defined in yml/md requiring translation.
  tmpls:
    - Set of templates used to generate forms and pages and
      submission responses.
  pages:
    - static pages, some generated from templates to serve up

  entries:
    - entries submitted by forms
      - entry-name-1-ts.json
      - entry-name-2-ts.json
      - entry-name-3-ts.json

*/
package magoo

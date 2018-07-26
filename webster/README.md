# Webster
Webster generates _static websites_ from _markdown_ files in a
directory. It does so by recursively walking a directory structure and
perfoms the following translations:

> **NOTE** After I had started writing webster, I came across Hugo that
does everything I was hoping to do with webster and so much more (and 
**super** fast).   I am now using Hugo.

1. **Markdown** files are converted into chunks of HTML content used
   to construct complete HTML pages.
2. **Front Matter (FM)** is YAML or JSON embedded in the markdown
   file.  _FM_ is used as Page meta data.
2. **Go Templates** are used to produce complete HTML pages with HTML
   content hunks and templates
3. **Tranplant** resulting tree will placed anywhere desired.

In a nutshell _webster_ recursively traverses a _source directory_
looking for _markdown_ files to be converted fast lean static html
pages.  Markdown files may have optional *front matter*, which is
embedded YAML or JSON to provide configuration (meta) information for
the page.

Example of page meta is:

```YAML
---
Title: Example using Page meta
Description: This article takes a little bit of time showing you how to use page meta in your markdown files and what it can be used for.
Permalink: /notes/example-using-websters-page-meta.html
Author: Rusty Eddy
Publish Date: 1/20/2021
Last Update: 2/22/2012
---
# This is standard markdown
The remainder of this example block of text would be interpreted as markdown.
```

The Front Matter is contained in the bracketing '---', it's YAML and
you can do whatever you want with it.  I'm using it to capture
information as in the above example.

## What Gets Copied

Webster will process markdown files, scss/sass and less to produce a
single .css file to be copied to destination.  Otherwise, webster will
ignore _.hidden_ directories.  Webster will scan directories that
begin with a '_' for go templates (.tmpl) to be used when generating
final HTML pages.  However, these '_template' directories will not be
copied to the destination.

Pretty much every thing else will get copied, or filtered out.  That is, all images files (.png, jpg, gif), PDF files, JavaScript and whatever else you want to serve up.

The source tree will be roughly replicated, after tranlations and eliminating certian files required for translation, but unecessary with final delivery.

## Example

See the example in the _example_ directory _simple.site_ which contains the sources for this utilites website.

```yaml
---

# This source structure will create the following destination
src: ./examples/simple.site

    # Markdown files get converted to html, html files are run through
    # go templates (or optionally not) to generate final html.  webster
    # will also process optionally processes YAML _front-matter_.
    - home.md
    - about.md
    - contact.md
    - notes:
        - note1.md
        - note2.md

    # Copy over all images (assets ...)
    - img:
        - logo.png

    # all scss files compile into a single style.css
    - style.scss
    - scss:
        - reset.scss
        - variables.scss
        - style.scss

    # Copy over all JavaScript (optionally minimize)
    - js/webster.js

# This directory is what our generated sites filesystem looks like.  Pretty sexy (and simple!) huh?

# The destination directory
dst: /srv/simple.site
    # HTML files generated from Markdown and optional FrontMatter merged with a Go template produced these html files.
    - home.html
    - about.html
    - contact.html

    # subdirs by default generate an index.html providing for a "menu" of subdir items.
    - notes:
        - index.html
        - note1.html
        - note2.html

    # images and other files (pdf, etc.). Bigger stuff should use a CDN*!
    - images:
        - logo.png

    # compiled style.css from .scss files
    - style.css

    # JavaScript
    - js
        - webster.js

This website can be moved to any location to be served up by a standard webserver like Apache or NGinx or any webserver capable of serving static HTML/JS/Image files.   Coming soon a built in, industrial strenght web static page web server...

---
```

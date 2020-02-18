# GoMVClean ![Go Report Card](https://goreportcard.com/badge/github.com/samiam2013/GoMVClean)
A Go HTTP2 library with it's own database. It's VERY configurable. Oh and it's an MVC.

(homonym) - A Go language MVC template written for keeping with only native dependencies

(exception) - it will utilize JavaScript and JSON.

I'm building out an API that will scale across a local network with very basic configuration. My library should be burstable to system limits with just the operating system as a middle layer and it should be distributable to an extent that seems like magic.

Rule 1 of GoMVClean is safety.

Rule 2 of GoMVClean is speed.

Rule 3 of GoMVClean is no middleware, just Go and Javascript.

Rule 4 of GoMVClean is `const separateDB = false`

# How to use it

0) hit my local host repository, make sure the code is working by clicking all
the links, try to crash my raspberry pi with traffic, lol:
https://www.dronetriage.com
it's going to warn you that the certificate is self-signed.

1) Download Go at [GoLang.Org](https://golang.org) if you don't have it

2) clone this project into a directory.

3) open a powershell or bash(? not tested) console

4) change to the package directory

5) type `go build`

6) `go run`

7) find a way to browse https://localhost

8) change the keys by generating your own for the `TLS/` directory before you connect it to the internet or you'll have no security

# Check Out it's Features (ultra basic mode for now)
Look at the index, it's a static file, and so is the form.

Look at the view, it's a go file, it just outputs what you send it

Realize that this view "/view/" tells you what it sees whether you visit it as a GET request or a POST form.

Look at modelQuery.go , modelFileTest.go , and modelBreakStuff.go

These are the important database checks and they also include the most important, most functional model functionality to be used.

load a page: renderStatic() in static.go

load an error page: errorShortCircuit()

get data from the model: loadStaticBody()

check and see if something's in the model: query()

write to a space in the model: uQuery()


# Basics
In Go, dynamic form-data handling is really easy because of access to low level data and high accessibility through libraries of functions. It's also insanely easy to reach into the folder structure to read and write, especially if you limit yourself to static HTML and leave page templating problems to Javascript and an API.

So with a Go language MVC website template, the entire model can be fractured into private and public instance folders. To handle this, permissions can be dynamically set for access to each user as measured by keeping a persistent-state variable for each user after a login. I'll be limit logins to session cookies and just avoid tracking otherwise since that stops me from having to implement persistent cookie storage and eliminates a man-in-the middle attack on "remember me" login cookies.

# Model Structure
The MySQL structure `database -> table -> column` is really well known and is being implemented by folders in descending order. This allows for a private folder for development of model scenarios to be implemented alongside the actual live site. This is implemented in a `query("folderPath")` with main folder paths (databases) `public/` and `private/` .  For now there's a Boolean switch disabling use of the private database until I implement a session cookies and database access permission classes.

This private/public structure means data from the public side is available for developing with the server while it is still running the site and running your new "private" model development (e.g. for private development of a smartphone app against the private model API)

The Model is made of folders with individual `schema.json` files scattered into folders by schema-self-defined structures. In the future, this will mean self-modifying file structure and documentation, so entire table templates and self-spawning tables are possible.

queries are just a URI attached to your domain like `site.domain/public/table/column/hashed(userId)`

Here, `hashed(userId)` is an assurance that even if permissions on the database are failing, brute-force search of the database will still be near impossible or impossible.

Based on if the library is set up to load balance, there could be a tag on the table name or api endpoint so that you get to the server with that particular table. Moving entire tables will require a copy paste, moving them to an entirely separate server could be implemented in a self-crawler (since only the server will have access to dangerous-if-stolen schema files and indexes). Once a table is moved, re-association will be a simple one line change to a go struct array with table names and local or fully qualified domain  associations, depending on whether you want to load balance with internal or external access. you can either have your user access the model or you could use requests inside Go itself.

I have public model endpoints for arbitrary data upload and download and I'm working on Cross-Site request forgery tokening so that forms are automatically validated. Since JavaScript comes with every browser, I'm going to try to write very vanilla JS or use JQuery for AJAX so that every page pull generates a new token and every page request deletes that token and generates a new one. CSRF everywhere except the home page + HTTPS = more safety. 

After all of the basic functionality is done, I'm going to attempt to write a Go database engine that understands the schema and can find a file arbitrarily and index arbitrarily after being queried with JSON or just a URL. I'm basically looking to write my own NoSQL MVC in Go in the form of an api that you could write a client for. So that my application, as long as it's running for me, allows anyone to just write an interface above it with web calls in effectively any user space connected to the internet. 

# Credits
This is not easy, and I'm only doing it because I personally want to see the results.

[me, if this is a moonshot]: "Houston, the flight trajectory gimbal has locked in lunar command!"

Dedicated to my lost Parents
Karla Dec 26 1963 - Nov 2 2002
Glenn Nov 2 1965 - Dec 17 2008

I want to thank: 

the team at atom editor, github.com/atom , 

the team at github desktop, https://desktop.github.com/ , 

github's github https://github.com/github , 

the team working on Go https://github.com/golang/go , 

and very ironically the team at microsoft Windows 10 

All of you for making 100 + updates to a github repository possible in like 40 hours or less of writing Go.


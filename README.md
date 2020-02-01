# GoMVClean [![Go Report Card](https://goreportcard.com/badge/github.com/samiam2013/GoMVClean)]
A Go language MVC template written for keeping with only native dependencies

A quote from internet land : "Go really doesn't need a framework. There have been many attempts, but have all failed because frameworks just get in the way. The core language was designed for building rest APIs. If you want MVC [in] Go, just start structuring your code that way."

What, so I can't do a good job of it and make the result free?

# How to use it

Download Go at Golang.org, clone this project into a directory.

go build;

go run;


Look at the index, it's a static file, and so is the form.

Look at the view, it's a go file, it just outputs what you send it

Realize that the view file tells you what it sees whether you visit it as a GET request or a POST form.

look at modelQuery.go and modelBreakStuff.go, these are the important database checks and they also include the most important, most functional stuff to be used.

Ignore the schema files, they're just stand-ins for whatever you want to use.


load a page: renderStatic() in static.go

load an error page: errorShortCircuit()

get data from the model: loadStaticBody()

check and see if something's in the model: query()

write to a space in the model: uQuery()


as far as I can tell this is all you *need* to make a website.


# Basics
In Go, dynamic form-data handling is really easy because of access to low level data and high accessibility through libraries of functions. It's also insanely easy to reach into the folder structure to read and write.

So with a Go language MVC website template, the entire model can be fractured into private and public instance folders.

Permissions can be dynamically set for access to each user as measured by keeping a persistent-state variable for each user after a login. I don't want to use cookies, I don't like cookies. I'll probably just limit logins to session cookies since that stops me from having to implement persistent cookie storage.

# Model Structure
The MySQL structure `database -> table -> column` is really well known and is being implemented by folders in descending order. This allows for a private folder for development of model scenarios to be implemented alongside the actual live site. This is implemented in a `query("folderPath")` with main folder paths `public/` and `private/` .  For now there's a Boolean switch disabling it until I can implement a solution for deciding if a user should be able to access the private database.

However, this private/public structure means data from the public side is available for developing with the server while it is still running the site and running your new "private" model development (e.g. for private development of a smartphone app against the private model API)

The Model is made of folders with individual `schema.json` files scattered into folders by schema-defined structures. In the future, this will mean self-modifying file structure.

like `site.domain/public/table/column/hashed(userId)`

Here, `hashed(userId)` is an assurance that even if permissions on the database are failing, brute-force search of the database will still be near impossible or impossible.

I have public model endpoints for arbitrary data upload and download and I'm working on Cross-Site request forgery tokening so that forms are automatically validated. Since JavaScript comes with every browser, I'm going to try to write very vanilla JS for AJAX so that every page pull generates a new token and every page request deletes that token and generates a new one. CSRF everywhere + a reverse HTTPS proxy like NGINX == hopefully safe

After that, I'm going to attempt to write a Go database engine that understands the schema and can find a file arbitrarily and index arbitrarily after being queried with JSON. I'm basically looking to write my own NoSQL MVC in Go.

# Epitaph
If I can pull this off, I'm pretty sure I just lasso'ed the moon into low earth orbit and tied it down. If it's finished, you get to write a website from a template that's written in JS and Go. As long as those two languages are supported it will have enforcement of how the code runs aside from browser ajax calls.

I will have EXACTING precision ability to enforce rules over what you send me, this is a second layer for the model just so that I can run my model with JS and users will never see any of the information. Users won't be able to see API endpoints, nor a schema, nor an algorithm, nor a giant library of JS waiting to load in their browser. Maybe this will be around with Go and JS for like hundreds, thousands of years based on the fact that signal processing work is largely done in assembly or C, 60 year old languages.

That's one small Go Lang Library for a man on this big blue marble, one giant

func leap(*manKind){manKind.stepSize(" YUGE - ish. ")}

Look Mom I got a good grade!
Karla Dec 26 1963 - Nov 2 2002

Look Dad I still don't have a job!
Glenn Myres Nov 2 1965 - Dec 17 2008

I was put on the back of your tombstone, now you're right above my epitaph! lol

seriously though, the money you left is all gone because I got a business idea but I didn't inherit much motivation, so here I am begging, https://paypal.me/GoMVClean Please don't give me any money if you don't have money, and please don't shame me for not getting a job programming; I'm not and never will be a sell-out.

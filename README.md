# GoMVClean [![Go Report Card](https://goreportcard.com/badge/github.com/samiam2013/GoMVClean)]
A Go language MVC template written for keeping with only native dependencies

Look Mom I got a good grade!
Karla Dec 26 1963 - Nov 2 2002

Look Dad I still don't have a job!
Glenn Myres Nov 2 1965 - Dec 17 2008

I was put on the back of your tombstone, now you're right above my epitaph! lol

# Basics
In Go, dynamic form-data handling is really easy because of access to low level data and high accessibility through libraries of functions. It's also insanely easy to reach into the folder structure to read and write.

So with a Go language MVC website template, the entire model can be fractured into private and public instance folders.

Permissions can be dynamically set for access to each user as measured by keeping a persistent-state variable for each user after a login. I don't want to use cookies, I don't like cookies. I'll probably just limit logins to session cookies since that stops me from having to implement persistent cookie storage.

# Model Structure
The MySQL structure `database -> table -> column` is really well known and is being implemented by folders in descending order. This allows for a private folder for development of model scenarios to be implemented alongside the actual live site. This is implemented in a `query("folderPath")` with main folder paths `public/` and `private/` .  For now there's a Boolean switch disabling it until I can implement a solution for deciding if a user should be able to access the private database.

However, this private/public structure means data from the public side is available for developing with the server while it is still running the site and running your new "private" model development (e.g. for private development of a smartphone app against the private model API)

The Model is basic folders with individual schemas, JSON files scattered into folders by schema-define structures.

like `site.domain/public/table/column/hashed(userId)`

Here, `hashed(userId)` is an assurance that even if permissions on the database are failing, brute-force search of the database will still be near impossible or impossible.

I have public model endpoints for arbitrary data upload and download and I'm working on Cross-Site request forgery tokening so that forms are automatically validated. Since JavaScript comes with every browser, I'm going to try to write very vanilla JS for AJAX so that every page pull generates a new token and every page request deletes that token and generates a new one. CSRF everywhere + a reverse HTTPS proxy like NGINX == hopefully safe

After that, I'm going to attempt to write a Go database engine that understands the schema and can find a file arbitrarily and index arbitrarily after being queried with JSON. I'm basically looking to write my own NoSQL MVC in Go.

# Epitaph
If I can pull this off, I'm pretty sure I just lasso'ed the moon into low earth orbit and tied it down. If it's finished, you get to write a website from a template that's written in JS and Go. As long as those two languages are supported it will have enforcement of how the code runs aside from browser ajax calls.

I will have EXACTING precision ability to enforce rules over what you send me, this is a second layer for the model just so that I can run my model with JS and users will never see any of the information. Users won't be able to see API endpoints, nor a schema, nor an algorithm, nor a giant library of JS waiting to load in their browser. Maybe this will be around with Go and JS for like hundreds, thousands of years based on the fact that signal processing work is largely done in assembly or C, 60 year old languages.

That's one small Go Lang Library for a man on this big blue marble, one giant

func leap(*manKind){manKind.stepSize(" YUGE - ish. ")}

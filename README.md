# GoMVClean
A Go language MVC template written for keeping with only native dependencies

# Basics
In Go, dynamic form-data handling is really easy because of access to low level data and high accessibility through libraries of functions. it's also insanely easy to reach into the folder structure to read and write.

So with a Go language MVC website template, the entire model can be fractured into private and public instance folders.

Permissions can be dynamically set for access to each user as measured by keeping a persistent-state variable for each user after a login. I don't want to use cookies, I'll probably just keep a persistent web socket if I can do that instead.

# Model Structure
The MySQL structure 'database' -> 'table' -> 'column' is really well known and is being implemented by folders in descending order. This allows for a private folder for development of model scenarios to be implemented alongside the actual live site. For now there's a Boolean switch disabling it until I can implement a solution for deciding if a user should be able to access the private database.

However, this private/public structure means data from the public side is available for developing with the server while it is still running the site and running your new "private" model development (e.g. for private development of a smartphone app against the private model API)

The Model is basic folders with individual schemas, JSON files scattered into folders by schema-define structures.

like `site.domain/public/table/column/hashed(userId)`

Here, `hashed(userId)` is an assurance that even if permissions on the database are failing, brute-force search of the database will still be near impossible or impossible.


# Let's make a JS emulator
The above is mostly done. I'm going to be working on a havaAv8.go script to boot-strap v8.

I'm going to have to force you to haveAv8.go yourself and build v8 if you're not using
1) Microsoft windows, because that's the development environment (thank GitHub and Atom)
2) Ubuntu/Debian Linux, because that's what I'm familiar with. My apologies go to the "OSX" and RedHat communities.

I cloned a build of Google's v8 engine for JS written in Go. This way I can create a website that enforces the version and runs almost all of the JS, LITERALLY on the backend.

How to emulate JS at the speed of C: https://github.com/augustoroman/v8

There will be a master V8 boot script. The console will be up-to speed on what it can query right after it's summoned and then a web socket can hold the connection open and return data.

This can all happen before the user asks for any data, for functionality like search recommendations. All without telling users how the model is being queried, running a ton of JS on their end, or having a centralized schema for the data.

This also has the potential to free front-end web designers from having to know anything about MySQL or even having to know there's a database because you can write one-off queries into random-data test folders for them to use to speed up development.

# Epitaph
If I can pull this off, I'm pretty sure I just lasso'ed the moon into low earth orbit and tied it down. If it's finished, you get to write a website that's written in JS and Go. As long as those two languages are supported it will have enforcement of how the code runs aside from browser ajax calls, because if the user doesn't run the PGP library to get into v8 they don't get in, and that's a lot of work, and it makes sending messages a lot of work.

However, I will have EXACTING precision ability to enforce what you send me, and it will always be secure, whether you use HTTPS:// or not, this is a second layer for the model just so that . Maybe this will be around with Go and Javascript for like hundreds, thousands of years based on the fact that signal processing work is largely done in assembly or C, 60 year old languages.

That's one small Go Lang Library for a man on this big blue marble, one giant

func leap(*manKind){manKind.stepSize(" YUGE - ish. ")}

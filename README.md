# GoMVClean
A Golang MVC template written for keeping with only native dependencies

In Go, dynamic form-data handling is really easy because of access to low level data and high accesibility through libraries of functions. it's also insanely easy to reach into the folder structure to read and write. 

So with a Go language MVC website template, the entire model (or database (or folder)) can be fractured into private and public instances. Permissions can be dynamically set for access to each user as measured by keeping a persistent-state variable for each user after a login.

The MySQL structure 'database' -> 'table' -> 'row' is really well known and can be implemented by folders in descending order. This allows for a private database for development of model scenarios to be implemented alongside the actual live site which means data for developing with the server while it is still running the site (e.g. for private development of the model API)

I'm going to write the entire database schema into basic folders with individual JSON files scattered into folders by date and then by id or potentially a hash value fractured into descending folders, 

like site.domain/public/users/bio/{hash(hash(user_id)}

Here, `hash(hash(user_id))` is an ensurement that even if permissions on the database are failing, brute-force search of the database will still be near impossible or impossible.


# I'm Go-ing Mad
Here's a fun joke, I might just clone a build of google's v8 engine for javascript in Go lang into the project, that way I can create a website that runs all of that javascript, LITERALLY on the backend.

I'll take a what is a js emulator for 2000 Alex. Hey, Node.JS you've got competition, Go lang and JQuery are stiff competiti-oin if you're Javascript + a JS library + a backend JS language + a type of compiled JS server or whatever. Magic to these ears. Lambdas.

let's emulate that at the speed of C. Let's get rid of C. https://github.com/augustoroman/v8

if I write the Javascript engine into this and implement model-endpoint calls with a tiny amount of javascript and jQuery I'll have the world's most powerful website-in-a-box because you can enforce the version of javascript EVERYONE is running.

The biggest problem lies ahead, to run the hashing algorithm against a jQuery call so that they can't be forged, so that the user can send me my own javascript but not someone else's Javascript. I have to implement https://github.com/openpgpjs/openpgpjs

But if I can pull that off, I'm pretty sure I just put a rope around a room sized moon in low earth orbit and tied it down, and if it's finished, you get to write a website that's written in Javascript and Golang, and as long as those two languages are supported it will have enforcement of how the code runs, because if the user doesn't run the PGP library to get into v8 they don't get in, and that's a lot of work, and it makes sending messages a lot of work. However, I will have EXACTING precision ability to enforce what you send me, and it will always be secure, whether you use HTTPS:// or not, this is a second layer for the model layer just for JS. Maybe this will be around with Go and Javascript for like hundreds, thousands of years based on the fact that signal processing work is largely done in assembly or C, 60 year old languages.

That's one small Go Lang Library for a man on this big blue marble, one giant 

func leap(*manKind){ *manKind.Realize( " this is Yuge-ish. " )  }

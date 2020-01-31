# GoMVClean
A Golang MVC template written for keeping with only native dependencies

In Go, dynamic form-data handling is really easy because of access to low level data and high accesibility through libraries of functions. it's also insanely easy to reach into the folder structure to read and write. 

So with a Go language MVC website template, the entire model (or database (or folder)) can be fractured into private and public instances. Permissions can be dynamically set for access to each user as measured by keeping a persistent-state variable for each user after a login.

The MySQL structure 'database' -> 'table' -> 'row' is really well known and can be implemented by folders in descending order. This allows for a private database for development of model scenarios to be implemented alongside the actual live site which means data for developing with the server while it is still running the site (e.g. for private development of the model API)

I'm going to write the entire database schema into basic folders with individual JSON files scattered into folders by date and then by id or potentially a hash value fractured into descending folders, 

like site.domain/public/users/bio/{hash(hash(user_id)}

Here, `hash(hash(user_id))` is an ensurement that even if permissions on the database are failing, brute-force search of the database will still be near impossible or impossible.

Here's a fun joke, I might just clone a build of google's v8 engine for javascript in Go lang into the project, that way I can create a website that runs all of that javascript, LITERALLY on the backend, what's up Node.JS you've got competition, Go lang and JQuery are stiff competitoin if you're Javascript + a library + a backend language + a type of compiled javascript server or whatever. Magic to these ears. 

if I write the Javascript engine into this and implement model-endpoint calls with a tiny amount of javascript and jQuery I'll have the world's most powerful website-in-a-box because you can enforce the version of javascript EVERYONE is running.

The biggest problem lies ahead, to run the hashing algorithm against a jQuery call so that they can't be forged, so that the user can sent me my own javascript but not someone else's, I have to implement PGP into JavaScript?, I have to integrate a Javascript PGP Library somehow.

But if I can pull that off, I'm pretty sure I just put a rope around a room sized moon in low earth orbit and tied it down, and if it's finished, you get to write a website that's written in Javascript and Golang, and as long as those two languages are supported, (maybe like hundreds, thousands of years based on the fact that signal processing work is largely done in assembly or C, 60 year old languages.

That's one small Go Lang Library for a man, one giant func leap(*manKind){ *manKind.Realize( " boom. " )  }

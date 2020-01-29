# GoMVClean
A Golang MVC template written for keeping with only native dependencies

In Go, dynamic form-data handling is really easy because of access to low level data and high accesibility through libraries of functions. it's also insanely easy to reach into the folder structure to read and write. 

So with a Go language MVC website template, the entire model (or database (or folder)) can be fractured into private and public instances. Permissions can be dynamically set for access to each user as measured by keeping a persistent-state variable for each user after a login.

The MySQL structure 'database' -> 'table' -> 'row' is really well known and can be implemented by folders in descending order. This allows for a private database for development of model scenarios to be implemented alongside the actual live site which means data for developing with the server while it is still running the site (e.g. for private development of the model API)

I'm going to write the entire database schema into basic folders with individual JSON files scattered into folders by date and then by id or potentially a hash value fractured into descending folders, 

like site.domain/public/users/bio/{hash(hash(user_id)}

Here, `hash(hash(user_id))` is an ensurement that even if permissions on the database are failing, brute-force search of the database will still be near impossible or impossible.

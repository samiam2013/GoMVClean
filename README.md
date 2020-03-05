# GoMVClean [![Go Report Card](https://goreportcard.com/badge/github.com/samiam2013/goMVClean)](https://goreportcard.com/report/github.com/samiam2013/goMVClean)
Go Website Template [not a framework] Written for Native-Only Dependencies

The goal is to have a singleton-pattern binary-compiled MVC framework.

You're welcome to play with this code, but it is not yet finished.

# What is this thing?
Have you ever wished that you didn't have to have several separarte pieces of software running to make your website answer a request for a single HTTPS page?

Does anyone really want a non-language dependency to be able to update and break the functionality of their application?

Have you ever wondered why websites and APIs are separate entities? Why memory caching systems are independent of persisten storage database engines? 

I hate duplicate code, I've seen a lot of it. It seems like the same wheel has been optimized to be rectangular, trapezoidal, triagnular and marginally round, sometimes, 4,5,6 times over for the same website.

I'm a backend-first programmer. I hate when an interface gets designed, the backend is implemented to make it work, and the job of making code scale to a larger, more robust, faster system is ignored as long as the company/entity has enough money to run a larger or faster individual server. 

I hate PHP. I hate SQL. I hate dependencies. I want to understand my own website, top to bottom, end to end, and I want it to run slow until it has to run fast, but I want it to run without ever duplicating a single line of code. 

I don't think websites should be capable of doing a single thing that they weren't created to do. Time after time, separate HTTP server, Database servers, cryptopgraphy libraries and frameworks like wordpress create security holes and don't get patched when they should. 

Google made Go to run their otherwise C-based APIs with more readable code. Hash maps are the core of every high-speed persistent storage, and the two can be combined to turn operating systems into single-entity binary-compiled high-speed websites.

Is this stupid, an uphill battle that I'll probably use? yes. Do I care? No. 

I see this as an effort to make a website kernel the way that Linux made an operating system kernel. I don't want it to be any bigger than it has to be, and I don't ever expect it to gain the attention of Linux or help people like Linux, but if it helps anyone, or encourages anyone to understand the entire set of algorithms that goes into making a website happen, I will have accomplished more than I wanted.



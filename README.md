## Usage

Clone the repository

`git clone https://github.com/nimbl3/go-buffalo-twitt.git`

Require dependency and Go plugins:
* [dep](https://github.com/golang/dep)
* [Goth] (https://github.com/markbates/goth)
* Postgres
* Node/Yarn
* [Buffalo](https://github.com/gobuffalo/buffalo) - v0.9.5 (or later or `development` branch)
* Go plugins for generator: layout, resource, bootstrap.

## Database Setup

It looks like you chose to set up your application using a postgres database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start postgres for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started postgres, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a 
	
### Buffalo setup

	$ buffalo setup
	
## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see the "Welcome" page.

[Powered by Buffalo](http://gobuffalo.io)

## License

This project is Copyright (c) 2014-2017 Nimbl3 Ltd. It is free software,
and may be redistributed under the terms specified in the [LICENSE] file.

[LICENSE]: /LICENSE

## About

![Nimbl3](https://dtvm7z6brak4y.cloudfront.net/logo/logo-repo-readme.jpg)

This project is maintained and funded by Nimbl3 Ltd.

We love open source and do our part in sharing our work with the community!
See [our other projects][community] or [hire our team][hire] to help build your product.

[community]: https://github.com/nimbl3
[hire]: https://nimbl3.com/

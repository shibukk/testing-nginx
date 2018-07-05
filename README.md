# testing-nginx
A configuration file testing of nginx

## Description

It's simple proxy for nginx configuration file testing.

browser <----> proxy server (nginx) <----> app server (golang/ruby)

## Example

See on [this commit](https://github.com/shibukk/testing-nginx/commit/50fb759b6a19eeeba27592bc9752ab978cad860a), if you track application user id at nginx.

## Requirement

`testing-nginx` depend on docker-compose.  
So you need installing [docker](https://docs.docker.com/install/).

## Install

```bash
$ docker-compose up
```

And try [http://localhost](http://localhost).

## Usage

### Change golang to ruby of application server

You can change language of application server.  
Please replace dockerfile path in `docker-compose.yml`.

```diff
   app:
     build:
       context: .
-      dockerfile: ./go/Dockerfile
+      dockerfile: ./ruby/Dockerfile
```

## Contribution

1. Fork it ( http://github.com/shibukk/testing-nginx/fork )
2. Create your feature branch ( git checkout -b my-new-feature )
3. Commit your changes ( git commit -am 'Add some feature' )
4. Push to the branch ( git push origin my-new-feature )
5. Create new Pull Request

## Licence

[MIT](https://github.com/shibukk/testing-nginx/blob/master/LICENSE)

## Author

[shibukk](https://github.com/shibukk)

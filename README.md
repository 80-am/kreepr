# kreepr
[![License: MIT](https://img.shields.io/github/license/80-am/kreepr)](https://opensource.org/licenses/MIT)


Creeping social media for changes in trends and behavior.

## Getting Started
Thsese instructions will get you up and running on your local machine.

Create empty db.
```mysql
CREATE DATABASE kreepr;
```

Then import db.sql using mysql or a db visualization program.
```bash
mysql -u username -p password kreepr < ./db.sql
```

Copy [secrets.yml.sample](secrets.yml.sample) into secrets.yml and fill in your secrets.


```yml
# db
user: "your db user"
password: "your db password"
schema: "/kreepr"

# twitter
key: "your twitter key"
secret: "your twitter secret" 
token: "your twitter token"
access_token: "your twitter access token"
access_secret: "your twitter access secret
```
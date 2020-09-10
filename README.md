# kreepr
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/80-am/kreepr)
[![License: GPL-3.0](https://img.shields.io/github/license/80-am/kreepr)](https://opensource.org/licenses/GPL-3.0)

Creeping social media for changes in trends and behavior.

## Getting Started
Thsese instructions will get you up and running on your local machine.

```sql
CREATE DATABASE kreepr;
```

Import db.sql using mysql or a db visualization program.
```bash
mysql -u username -p password kreepr < ./db.sql
```

Set up your subject to kreep by inserting their Twitter username into the subjects table. 

```sql
INSERT into subjects(userName) VALUES('LisaSu');
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
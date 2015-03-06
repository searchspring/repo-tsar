RepoTsar
========

Manage multiple git repos

Installing
==========

## Requirements
This project is written in go.

https://golang.org/doc/install

Don't forget to set your GOPATH and add the locations of the go bin to your PATH.  Something like this in your .bashrc :

```bash
export GOPATH=$HOME/code/go
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin/
```

# Installing git2go

This project requires git2go.v22, which in turn requires libgit2.  You will need to install a go and libgit2.  If you want SSH support with this application, ensure you have libssh2 installed as well.  A recent version of cmake is required to build libgit2 http://www.cmake.org/download/ . 


```bash
go get -d gopkg.in/libgit2/git2go.v22
cd $GOPATH/src/gopkg.in/libgit2/git2go.v22
git submodule update --init 
make install
```

# Installing yaml

This project also requires yaml.v2

```bash
go get gopkg.in/yaml.v2
```

# Installing RepoTsar

```bash
go get github.com/SearchSpring/RepoTsar
``` 


Usage
=====

Edit the repotsar.yml to configure you signature and repos.

```YAML
signature:
    name: "Your Name"
    email: email@address.com

repos:
    git2go:
        url: ssh://git@github.com/libgit2/git2go.git
        path: /tmp/git2go
        branch: master

```
(Reminder: YAML format is space indented, not tab)

Running repotsar without arguments will concurrently and idempotently create paths if needed, clone repos if needed, and git pull.

### Options
```RepoTsar --repos repo1,repo2,repo3 ```

Supply a comma seperated list of defined repos from your repotsar.yml to act on.

``` --branch BranchName ```

Will in addition create local branches in all repos being acted on.


License and Author
==================

* Author:: Greg Hellings (<greg@thesub.net>)


Copyright 2014, B7 Interactive, LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

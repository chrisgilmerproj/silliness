# Silliness

This is a repo to track silly code projects.  These may or may not end up in
their own repos at some later date.

## How To Break Out Projects

Speaking of which, having now broken out two projects from this repo I want to
keep the commands for doing it again.

Clone this repo into a new directory `NEWREPO` and get just the directory you
want as master:

```sh
$ git clone git@github.com:chrisgilmerproj/silliness.git NEWREPO
$ cd NEWREPO
$ git remote rm origin
$ git filter-branch --prune-empty --subdirectory-filter DIRNAME master
$ git remote add origin https://github.com/chrisgilmerproj/NEWREPO.git
$ git push origin .
```

Verify that everything is now correctly up in GitHub and all the history is how
you want it.  Now go back to this repo and run this command:

```sh
$ git filter-branch --index-filter "git rm -r --cached --ignore-unmatch DIRNAME" --prune-empty
$ git push -f origin .
```

For more information see:
- https://help.github.com/articles/splitting-a-subfolder-out-into-a-new-repository/
- http://blogs.atlassian.com/2014/04/tear-apart-repository-git-way/

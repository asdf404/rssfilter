# RSS/Atom Proxy

For filtering annoying posts.

## Usage

```yml
settings:
  port: 8080
feeds:
- url: http://user:password@example.tld/rss.xml
  path: filtered.xml
  keep:
  - author: my-favorite-author
  - author: another-cool-guy
  skip:
  - category: some boring category
    author: hate-this-guy
  - title: some boring post

```

```
docker run -p 8080:8080 -v config.yml:/config.yml asdf404/rssfilter
```

# blog-proto

this repository is protocol buffer list for pos's Blog posts.

following is core file tree.

```bash
.
├── proto
│   └── post
│       └── post.proto
└── rpc
    └── post
        └── post.pb.go
```

## update post package

```bash
$ make gen
```

## Other Repository

post repository manages blog posts

<a href="https://github.com/po3rin/post"><img src="img/post.png" width="450px" /></a>

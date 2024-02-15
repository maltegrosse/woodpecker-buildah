# Woodpecker-Buildah Plugin
A basic wrapper for buildah commands to run as a woodpecker-ci pipeline.

Inspired by https://codeberg.org/Taywee/woodpecker-buildah/ , rewritten in golang.

## Links
A collection of useful buildah articles
- https://codeberg.org/Taywee/woodpecker-buildah/
- https://www.redhat.com/sysadmin/7-transports-features
- https://github.com/containers/buildah/issues/2554
- https://www.redhat.com/sysadmin/podman-inside-kubernetes
- https://opensource.com/article/19/3/tips-tricks-rootless-buildah
- https://github.com/containers/buildah/blob/main/docs/buildah.1.md
- https://insujang.github.io/2020-11-09/building-container-image-inside-container-using-buildah/
- https://danmanners.com/posts/2022-01-buildah-multi-arch/


flags 
- storage driver + other flags
- build containe + buildah (fuse + qemu other package)


## License

This wrapper is under MIT, [buildah image](https://github.com/containers/buildah/blob/04c61a7b7277e44ea69ea93ebbded92fdecac036/contrib/buildahimage/Containerfile) is under the Apache license.
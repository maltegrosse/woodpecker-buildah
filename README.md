# Woodpecker-Buildah Plugin
A basic wrapper for buildah commands to run as a woodpecker-ci pipeline.

Inspired by https://codeberg.org/Taywee/woodpecker-buildah/ , rewritten in golang.

## Usage
```
steps:
  build_and_release_and_push:
    image: maltegrosse/woodpecker-buildah:0.0.8
    pull: true
    settings:
      registry: somehub.com
      repository: theuser/mytarget_repo
      tag: 4.0.12c
      architectures: amd64 aarch64
      context: Dockerfile
      username:
        from_secret: docker_username
      password:
        from_secret: docker_password
```
## Limitation
The plugin runs with vfs - and is quite slow... see links for further information.

Plugin in early stage. Only tested with kubernetes backend. To run multi-arch builds, a second qemu container needs to be deployed (in privileged mode) --> see example-qemu.yaml

Fuse package is preinstalled, and fuse storage could be added as a flag. (untested)

If buildah runs in privileged mode, woodpecker needs to trust the container repo. See https://woodpecker-ci.org/docs/administration/server-config#all-server-configuration-options --> WOODPECKER_ESCALATE

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
- https://www.itix.fr/blog/build-multi-architecture-container-images-with-kubernetes-buildah-tekton-and-qemu/
- https://developers.redhat.com/blog/2019/08/14/best-practices-for-running-buildah-in-a-container#running_buildah_inside_a_container


## License

This wrapper is under MIT, [buildah image](https://github.com/containers/buildah/blob/04c61a7b7277e44ea69ea93ebbded92fdecac036/contrib/buildahimage/Containerfile) is under the Apache license.
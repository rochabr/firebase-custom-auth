Firebase currently supports many forms of authentication outside of google, but not all of them.

This project provides an API written in GO that exposes an endpoint which receives a externally provided token and creates a firebase custom token based on that.

For it to work, a firebase project needs to be created and a service account file named "service-acc.json" needs to be generated from the project options.

- Dockerfile generates a docker image.
- app.yaml pushes the project to appengine
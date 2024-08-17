# gmh (Gitops Map-Hack)

gmh is a tool to generate architecture drawings using architecture as code principles while parsing exiting gitops infrastructure.


## Roadmap

- [] Allow annotations on specific blocks
- [] Allow defining nesting and grouping of similar entities
- [] Allow defining architectural views
- [] Check a PR change for any architectural view break (meaning breaking change pending a review)
* Support providers:
  - []  kubernetes yaml
  - []  helm templating
  - []  kustomize
  - []  terraform state
  - []  terraform plan
* Support output diagrams:
  - [] mermaid
  - [] excalidraw
  - [] d2
* Support drawing:
  - [] allow resources to be rendered as arrows (e.g. netpols for kubernetes)
* Support styles:
  - [] associate image with element
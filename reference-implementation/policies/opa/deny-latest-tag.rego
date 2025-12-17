package kubernetes.admission

# Deny images using :latest tag (training example)
deny[msg] {
  input.request.kind.kind == "Pod"
  container := input.request.object.spec.containers[_]
  endswith(container.image, ":latest")
  msg := sprintf("image %q uses forbidden tag :latest", [container.image])
}

components {
  id: "body"
  component: "/examples/create_body/body.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"box\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/images/images.atlas\"\n"
  "}\n"
  ""
}

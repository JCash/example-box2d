components {
  id: "controller"
  component: "/main_menu.script"
}
components {
  id: "gui"
  component: "/main_menu.gui"
}
embedded_components {
  id: "many_proxy"
  type: "collectionproxy"
  data: "collection: \"/examples/many/many.collection\"\n"
}
embedded_components {
  id: "joints_proxy"
  type: "collectionproxy"
  data: "collection: \"/examples/joints/joints.collection\"\n"
}
embedded_components {
  id: "create_body_proxy"
  type: "collectionproxy"
  data: "collection: \"/examples/create_body/create_body.collection\"\n"
}

job "ipfs" {
  datacenters = ["dc1"]
  type = "service"

  group "ipfs-group" {
    count = 5

    task "ipfs" {
      driver = "docker"

      config {
        image = "ipfs/go-ipfs:latest"
        port_map {
          http = 8080
        }
      }

      resources {
        cpu    = 500
        memory = 256
      }
    }
  }
}

variable "envfile" {
  type    = string
  default = "../.env"
}

locals {
  envfile = {
    for line in split("\n", file(var.envfile)) : split("=", line)[0] => regex("=(.*)", line)[0]
    if !startswith(line, "#") && length(split("=", line)) > 1
  }
}



// Define an environment named "local"
env "local" {
  src = "file://schema.hcl"
  url = "mysql://root:password@localhost:33066/gva"

  // Define the URL of the Dev Database for this environment
  dev = "mysql://root:password@localhost:33066/gva_dev"
  migration {
    dir = "file://sql"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "docker" {
  src = "file://schema.hcl"
  url = local.envfile["DB__URL"]
  dev = local.envfile["DB__URL__DEV"]

  migration {
    dir = "file://sql"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

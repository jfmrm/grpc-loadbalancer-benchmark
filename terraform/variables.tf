variable "cluster_name" {
    type = string
    default = "joao-moura-test-cluster"
}

variable "cluster_version" {
    type = string
    default = "1.22"
}

variable "region" {
    type = string
    default = "us-east-1"
}

variable "tags" {
    type = map(string)
    default = {
        Terraform = "true"
        Owner = "João Moura"
        DeleteMe = "true"
    }
}

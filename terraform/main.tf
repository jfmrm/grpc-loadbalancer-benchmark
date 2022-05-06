provider "aws" {
    region = var.region
}

module "vpc" {
    source = "terraform-aws-modules/vpc/aws"
    version = "3.14"

    name = "${var.cluster_name}-vpc"
    cidr = "10.0.0.0/16"

    azs             = ["${var.region}a", "${var.region}b", "${var.region}c"]
    private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
    public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

    enable_nat_gateway = true
    single_nat_gateway = true
    enable_vpn_gateway = false

    tags = var.tags
}

module "eks" {
    source  = "terraform-aws-modules/eks/aws"
    version = "~> 18.20"

    cluster_name    = var.cluster_name
    cluster_version = "1.21"

    cluster_endpoint_private_access = true
    cluster_endpoint_public_access  = true

    vpc_id     = module.vpc.vpc_id
    subnet_ids = module.vpc.private_subnets

    # EKS Managed Node Group(s)
    eks_managed_node_group_defaults = {
        disk_size      = 50
        instance_types = ["t3.medium"]
    }

    eks_managed_node_groups = {
        default = {
            min_size     = 3
            max_size     = 3
            desired_size = 3

            instance_types = ["t3.large"]
            capacity_type  = "ON_DEMAND"
        }
    }

    aws_auth_users = [
        {
            userarn  = "arn:aws:iam::169950403031:user/joao.moura"
            username = "joao.moura"
            groups   = ["system:masters"]
        },
    ]

    tags = var.tags
}

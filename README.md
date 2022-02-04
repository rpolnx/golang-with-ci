# Golang CI/CD with aws deploy

## Purpose

Deploy basic golang application on aws using ansible and building infrastructure from terraform

## Technologies

- Golang
- Terraform
- AWS stack
- Ansible
- Docker + docker-compose

# How to use

## Setup

You need to provide those secrets:

- ANSIBLE_HOSTS as example infra/ansible/lightsail.hosts
- DOCKERHUB_USERNAME and DOCKERHUB_PASSWORD (dockerhub)
- PERSONAL_TOKEN github personal token to trigger CD from CI
- SSH_LIGHTSAIL_DEV_KEY (ssh private key to connect on instance)

First, connect on instance on /usr/app and populates a new file .env.production with current values.
I choose to do manually this part because is not good to use "secrets" and env values com CI. Unfortunately, we aren't using kubernetes to has a secret structure.

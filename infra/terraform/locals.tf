locals {
    user_data = <<-EOF
apt update -y
apt install apt-transport-https -y
curl -fsSL https://get.docker.com | bash | tee /var/log/docker-instalation.log
apt install docker-compose -y
apt install python3-docker -y
EOF
}